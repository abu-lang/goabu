package goabu

import (
	"container/list"
	"sync"

	"github.com/abu-lang/goabu/stringset"
)

type reader struct {
	workingSet stringset.Set
	status     string
	blocking   bool
}

type writer struct {
	workingSet stringset.Set
	optimistic bool
	awaiting   int

	// wake if not nil will receive a signal each time a closeRead
	// returns if optimistic == false otherwise it will receive a
	// signal when awaiting becomes 0.
	wake chan struct{}
}

type coordinator struct {
	readers    map[string]uint
	requesting *list.List
	awake      bool
	reading    map[key]reader
	writing    *writer
	mutex      sync.Mutex
}

func newCoordinator() *coordinator {
	return &coordinator{
		readers:    make(map[string]uint),
		requesting: list.New(),
		reading:    make(map[key]reader),
	}
}

func (c *coordinator) requestRead(ws stringset.Set) key {
	var e *list.Element = nil
	for {
		c.mutex.Lock()
		if c.writing == nil || !ws.IntersectsWith(c.writing.workingSet) {
			for r := range ws {
				c.readers[r]++
			}
			if e != nil {
				n := e.Next()
				c.requesting.Remove(e)
				if n != nil {
					w := n.Value.(chan struct{})
					w <- struct{}{}
				} else {
					c.awake = false
				}
			}
			var i key = 1
			for {
				_, present := c.reading[i]
				if !present {
					break
				}
				i++
			}
			c.reading[i] = reader{workingSet: ws, status: "interested"}
			c.mutex.Unlock()
			return i
		}
		if e == nil {
			e = c.requesting.PushBack(make(chan struct{}))
		} else {
			c.awake = false
		}
		wake := e.Value.(chan struct{})
		c.mutex.Unlock()
		<-wake
	}
}

func (c *coordinator) requestWrite(optimistic bool) {
	var e *list.Element = nil
	for {
		c.mutex.Lock()
		if e != nil {
			n := e.Next()
			if n != nil {
				w := n.Value.(chan struct{})
				w <- struct{}{}
			} else {
				c.awake = false
			}
		}
		if c.writing == nil {
			c.writing = &writer{optimistic: optimistic}
			if e != nil {
				c.requesting.Remove(e)
			}
			c.mutex.Unlock()
			return
		}
		if e == nil {
			e = c.requesting.PushBack(make(chan struct{}))
		}
		wake := e.Value.(chan struct{})
		c.mutex.Unlock()
		<-wake
	}
}

func (c *coordinator) fixWorkingSetWrite(ws stringset.Set) {
	if c.writing.optimistic {
		c.startOptimistic(ws)
	} else {
		c.startWrite(ws)
	}
}

func (c *coordinator) startWrite(ws stringset.Set) {
	for {
		c.mutex.Lock()
		if c.writing.wake != nil {
			c.awake = false
		}
		c.wakeRequesting()
		ok := true
		for r := range ws {
			if c.readers[r] > 0 {
				ok = false
				break
			}
		}
		if ok {
			c.writing.workingSet = ws
			c.writing.wake = nil
			c.mutex.Unlock()
			return
		}
		if c.writing.wake == nil {
			c.writing.wake = make(chan struct{})
		}
		c.mutex.Unlock()
		<-c.writing.wake
	}
}

func (c *coordinator) startOptimistic(ws stringset.Set) {
	c.mutex.Lock()
	cs := stringset.Make()
	for r := range ws {
		if c.readers[r] > 0 {
			cs.Insert(r)
		}
	}
	c.writing.workingSet = ws
	for k, reader := range c.reading {
		if reader.workingSet.IntersectsWith(cs) {
			if reader.status != "prepared" {
				reader.status = "aborted"
			} else {
				reader.blocking = true
				c.writing.awaiting++
			}
			c.reading[k] = reader
		}
	}
	if c.writing.awaiting == 0 {
		c.mutex.Unlock()
		return
	}
	c.writing.wake = make(chan struct{})
	c.mutex.Unlock()
	<-c.writing.wake
}

func (c *coordinator) confirmRead(k key) bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	reader, present := c.reading[k]
	if !present || reader.status == "aborted" {
		return false
	}
	reader.status = "prepared"
	c.reading[k] = reader
	return true
}

func (c *coordinator) confirmWrite() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.writing.workingSet = stringset.Make()
	c.wakeRequesting()
}

func (c *coordinator) closeRead(k key) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	reader, present := c.reading[k]
	if !present {
		return
	}
	delete(c.reading, k)
	for r := range reader.workingSet {
		c.readers[r]--
	}
	if reader.blocking {
		c.writing.awaiting--
	}
	c.wakeNext()
}

func (c *coordinator) closeWrite() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.writing = nil
	c.wakeRequesting()
}

// wakeRequesting awakes the next pending call to requestRead or requestWrite if such a call exists
// and c.awake == false.
func (c *coordinator) wakeRequesting() {
	if !c.awake && c.requesting.Len() > 0 {
		w := c.requesting.Front().Value.(chan struct{})
		w <- struct{}{}
		c.awake = true
	}
}

func (c *coordinator) wakeNext() {
	if c.awake {
		return
	}
	if c.writing != nil && c.writing.wake != nil && (!c.writing.optimistic || c.writing.awaiting == 0) {
		c.wakeWriter()
	} else {
		c.wakeRequesting()
	}
}

// wakeWriter awakes a pending fixWorkingSetWrite call by sending an empty struct on c.writing.wake.
// It should be called only when c.awake == false.
func (c *coordinator) wakeWriter() {
	c.writing.wake <- struct{}{}
	if !c.writing.optimistic {
		c.awake = true
	} else {
		c.writing.wake = nil
	}
}
