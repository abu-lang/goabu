package communication

import (
	"encoding/json"
	"errors"
	"fmt"
	"steel-lang/config"
	"steel-lang/misc"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/hashicorp/memberlist"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	// milliseconds
	timeoutRegister = 1000
	msgBuffLen      = 10
	registrySize    = 0
)

const (
	TestsNothing = iota
	TestsAbort
	TestsUnreliableSend
	// Halting
	TestsMidInterested
	TestsAfterInterested
	TestsMidFirst
	TestsAfterFirst
	TestsMidSecond
)

var TestsMidSends = 2

type registryInventory struct {
	Sender    *memberlist.Node
	Inventory misc.StringSet
}

type messageUnion struct {
	Type   string
	Sender *memberlist.Node

	Registry    resourceRegistry
	Transaction transactionInfo
}

type memberlistAgent struct {
	// real time access from delegate
	localResources misc.StringSet
	initialNodes   []string
	registry       resourceRegistry
	lockRegistry   *sync.RWMutex
	listPtr        **memberlist.Memberlist
	terminated     map[string]string
	logLevel       zap.AtomicLevel
	logger         *zap.Logger

	// delegate sees value set in Start()
	running bool
	config  *memberlist.Config
	list    *memberlist.Memberlist
	// chan Node.Name
	waitingForRegistry    chan string
	haltUpdates           chan bool
	quitTransactions      chan chan bool
	quitGossip            chan chan bool
	quitDemux             chan chan bool
	pendingUpdates        chan resourceRegistry
	transactionMessages   chan messageUnion
	transactionResponses  chan messageUnion
	coordinatedChannels   chan chan transactionChannels
	trackGossip           chan chan *sync.WaitGroup
	initiatedTransactions int
	transaction           transactionInfo
	// not modified after constructor
	listeningPort     int
	operations        chan chan []byte
	operationCommands chan chan string
	// testing
	test       int
	halted     bool
	lockHalted *sync.Mutex
}

func MakeMemberlistAgent(names misc.StringSet, port int, nodes []string, lc config.LogConfig) *memberlistAgent {
	res := &memberlistAgent{
		running:               false,
		listeningPort:         port,
		localResources:        names,
		initialNodes:          nodes,
		lockRegistry:          &sync.RWMutex{},
		initiatedTransactions: 0,
		operations:            make(chan chan []byte),
		operationCommands:     make(chan chan string),
		transaction: transactionInfo{
			Initiator: "",
		},
	}
	res.listPtr = &res.list
	if lc.Encoding == "" {
		lc.Encoding = "console"
	}
	zapCfg, ok := config.LogPreset(lc.Encoding).(zap.Config)
	if !ok {
		zapCfg = zap.NewProductionConfig()
	}
	res.logLevel = zapCfg.Level
	var err error
	res.logger, err = zapCfg.Build()
	if err != nil {
		if ok { // fallback to zap default
			ok = false
			zapCfg = zap.NewProductionConfig()
			res.logLevel = zapCfg.Level
			res.logger, err = zapCfg.Build()
		}
		if err != nil {
			panic("could not create memberlistAgent logger")
		}
	}
	res.SetLogLevel(lc.Level)
	if !ok {
		res.logger.Warn("could not load memberlistAgent logger config")
	}
	return res
}

func (a *memberlistAgent) IsRunning() bool {
	return a.running
}

func (a *memberlistAgent) Start() error {
	if a.running {
		return errors.New("agent is already running")
	}
	uuid, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	a.registry = makeResourceRegistry(a.localResources, uuid.String(), registrySize)
	a.terminated = make(map[string]string)
	a.waitingForRegistry = make(chan string, msgBuffLen)
	a.haltUpdates = make(chan bool)
	a.quitTransactions = make(chan chan bool)
	a.quitGossip = make(chan chan bool)
	a.quitDemux = make(chan chan bool)
	a.pendingUpdates = make(chan resourceRegistry, msgBuffLen)
	a.transactionMessages = make(chan messageUnion, msgBuffLen)
	a.transactionResponses = make(chan messageUnion, msgBuffLen)
	a.coordinatedChannels = make(chan chan transactionChannels)
	a.trackGossip = make(chan chan *sync.WaitGroup)

	a.config = memberlist.DefaultLocalConfig()
	stdLog, err := zap.NewStdLogAt(a.logger, zapcore.DebugLevel)
	if err != nil {
		return err
	}
	a.config.Logger = stdLog
	a.config.BindPort = a.listeningPort
	a.config.Name = uuid.String()
	a.config.Delegate = *a
	a.config.Events = *a

	// start listening
	a.list, err = memberlist.Create(a.config)
	if err != nil {
		return err
	}

	a.running = true
	go joiner(a.trackGossip, a.quitGossip)
	go a.handleUpdates()
	go demuxResponses(a.coordinatedChannels, a.transactionResponses, a.quitDemux)
	go a.handleTransactions()
	return nil
}

func (a *memberlistAgent) Join() error {
	if !a.running {
		return errors.New("agent is not running")
	}
	if len(a.initialNodes) > 0 {
		_, err := a.list.Join(a.initialNodes)
		return err
	}
	return nil
}

func (a *memberlistAgent) ForAll(payload []byte) error {
	if !a.running {
		return errors.New("agent is not running")
	}
	info := transactionInfo{
		Initiator: a.list.LocalNode().Name,
		Number:    a.initiatedTransactions,
		Payload:   payload,
	}
	a.initiatedTransactions++
	info.Partecipants = a.interested(info)
	if len(info.Partecipants) == 0 {
		fmt.Println("terminated transaction: none interested")
		return nil
	}
	return a.coordinateTransaction(info)
}

func (a *memberlistAgent) ReceivedActions() (<-chan chan []byte, <-chan chan string) {
	return a.operations, a.operationCommands
}

func (a *memberlistAgent) Stop() error {
	if !a.running {
		return errors.New("agent is not running")
	}

	fmt.Println("Stopping transaction handling...")
	replyCh := make(chan bool)
	a.quitTransactions <- replyCh
	<-replyCh
	a.operations <- nil
	fmt.Println("Stopped transaction handling")
	fmt.Println("Stopping response demultiplexing...")
	replyCh = make(chan bool)
	a.quitDemux <- replyCh
	<-replyCh
	fmt.Println("Stopped response demultiplexing")
	fmt.Println("Stopping update handling...")
	a.haltUpdates <- true
	fmt.Println("Stopped update handling")
	fmt.Println("Gossiping leave...")
	err := a.list.Leave(a.config.PushPullInterval)
	if err != nil {
		fmt.Println("error in gossiping leave:", err.Error())
	} else {
		fmt.Println("Gossiped leave")
	}
	fmt.Println("Detaching from group...")
	err = a.list.Shutdown() // v0.2.4 always returns nil
	if err != nil {
		fmt.Println("error in leaving group:", err.Error())
	} else {
		fmt.Println("Left group")
	}
	fmt.Println("Stopping gossip handling...")
	replyCh = make(chan bool)
	a.quitGossip <- replyCh
	<-replyCh
	fmt.Println("Stopped gossip handling")

	// clean up
	a.registry = nil
	a.terminated = nil
	a.list = nil
	a.config = nil
	a.waitingForRegistry = nil
	a.pendingUpdates = nil
	a.transactionMessages = nil
	a.quitTransactions = nil
	a.haltUpdates = nil
	a.quitGossip = nil
	a.quitDemux = nil
	a.transactionResponses = nil
	a.coordinatedChannels = nil
	a.trackGossip = nil
	a.running = false
	return nil
}

func (a *memberlistAgent) SetLogLevel(l int) {
	if l < config.LogDebug {
		l = config.LogDebug
	} else if l > config.LogFatal {
		l = config.LogFatal
	}
	zapLevel := zapcore.InfoLevel
	switch l {
	case config.LogDebug:
		zapLevel = zapcore.DebugLevel
	case config.LogWarning:
		zapLevel = zapcore.WarnLevel
	case config.LogError:
		zapLevel = zapcore.ErrorLevel
	case config.LogFatal:
		zapLevel = zapcore.DPanicLevel
	}
	a.logLevel.SetLevel(zapLevel)
}

func (a *memberlistAgent) handleUpdates() {
	for {
		select {
		case <-a.haltUpdates:
			return
		case remoteRegistry := <-a.pendingUpdates:
			for nodeName, resources := range remoteRegistry {
				a.lockRegistry.RLock()
				entry, present := a.registry[nodeName]
				a.lockRegistry.RUnlock()
				if !present || (entry != nil && resources == nil) {
					a.lockRegistry.Lock()
					if len(a.registry) == registrySize {
						a.lockRegistry.Unlock()
						break
					}
					a.registry[nodeName] = resources
					a.lockRegistry.Unlock()
				}
			}
		case destName := <-a.waitingForRegistry:
			for _, node := range a.list.Members() {
				if node.Name == destName {
					message := messageUnion{
						Type:     "registry_response",
						Sender:   a.list.LocalNode(),
						Registry: a.registry,
					}
					a.lockRegistry.RLock()
					localRegistry, err := json.Marshal(message)
					a.lockRegistry.RUnlock()
					if err != nil {
						fmt.Println("error in message marshalling:", err.Error())
						return
					}
					err = a.list.SendReliable(node, localRegistry)
					if err != nil {
						fmt.Println("error in sending registry response to", destName, err.Error())
						return
					}
					fmt.Println("sent registry response to", node.Name)
					break
				}
			}
		}
	}
}

//---------------------------------DELEGATES----------------------------------

func joiner(track <-chan chan *sync.WaitGroup, quit <-chan chan bool) {
	var waitGroup sync.WaitGroup
	for {
		select {
		case t := <-track:
			waitGroup.Add(1)
			t <- &waitGroup
		case q := <-quit:
			defer func() { q <- true }()
			waitGroup.Wait()
			return
		}
	}
}

func (d memberlistAgent) register() (*sync.WaitGroup, error) {
	replyCh := make(chan *sync.WaitGroup)
	select {
	case d.trackGossip <- replyCh:
		return <-replyCh, nil
	case <-time.After(time.Millisecond * timeoutRegister):
		return nil, errors.New("timeout in waiting from joiner")
	}
}

func (d memberlistAgent) NodeMeta(limit int) []byte {
	return []byte{}
}

func (d memberlistAgent) NotifyMsg(m []byte) {
	group, err := d.register()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer group.Done()

	var message messageUnion
	err = json.Unmarshal(m, &message)
	if err != nil {
		fmt.Println("error in message unmarshalling:", err.Error())
		return
	}
	switch message.Type {
	case "registry_request":
		select {
		case d.waitingForRegistry <- message.Sender.Name:
			fmt.Println("received registry request from", message.Sender.Name)
		default:
			fmt.Println("discarded incoming registry request from", message.Sender.Name)
		}
	case "registry_response":
		select {
		case d.pendingUpdates <- message.Registry:
			fmt.Println("received registry response from", message.Sender.Name)
		default:
			fmt.Println("discarded incoming registry response from", message.Sender.Name)
		}
	case "interested":
		fallthrough
	case "not_interested":
		fallthrough
	case "prepared":
		fallthrough
	case "aborted":
		fallthrough
	case "committed":
		select {
		case d.transactionResponses <- message:
		default:
			fmt.Println("discarded response from", message.Sender.Name)
		}
	default:
		select {
		case d.transactionMessages <- message:
		default:
			fmt.Println("discarded incoming transaction message from", message.Sender.Name)
		}
	}
}

func (d memberlistAgent) GetBroadcasts(overhead, limit int) [][]byte {
	return [][]byte{}
}

func (d memberlistAgent) LocalState(join bool) []byte {
	group, err := d.register()
	if err != nil {
		fmt.Println(err.Error())
		return []byte{}
	}
	defer group.Done()

	if join {
		d.lockRegistry.RLock()
		localRegistry, err := json.Marshal(d.registry)
		d.lockRegistry.RUnlock()
		if err != nil {
			fmt.Println("error in registry marshalling:", err.Error())
			return []byte{}
		}
		fmt.Println("join: sending registry")
		return localRegistry
	}
	d.lockRegistry.RLock()
	inventory := registryInventory{
		Sender:    (*d.listPtr).LocalNode(),
		Inventory: d.registry.inventory(),
	}
	d.lockRegistry.RUnlock()
	inventoryMsg, err := json.Marshal(inventory)
	if err != nil {
		fmt.Println("error during inventory marshalling:", err.Error())
		return []byte{}
	}
	fmt.Println("sending inventory")
	return inventoryMsg
}

func (d memberlistAgent) MergeRemoteState(buf []byte, join bool) {
	group, err := d.register()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer group.Done()

	if join {
		var remoteRegistry resourceRegistry
		err := json.Unmarshal(buf, &remoteRegistry)
		if err != nil {
			fmt.Println("join: error in registry unmarshalling:", err.Error())
			return
		}

		d.lockRegistry.RLock()
		size := len(d.registry)
		d.lockRegistry.RUnlock()
		if size == registrySize {
			fmt.Println("join: discarded received registry as already reached maximum registry size")
			return
		}
		select {
		case d.pendingUpdates <- remoteRegistry:
			fmt.Println("join: received registry")
		default:
			fmt.Println("join: discarded received registry")
		}
		return
	}
	var remoteInventory registryInventory
	err = json.Unmarshal(buf, &remoteInventory)
	if err != nil {
		fmt.Println("error in registry unmarshalling:", err.Error())
		return
	}
	d.lockRegistry.RLock()
	size := len(d.registry)
	d.lockRegistry.RUnlock()
	if size == registrySize {
		return
	}
	d.lockRegistry.RLock()
	inventory := d.registry.inventory()
	d.lockRegistry.RUnlock()
	if !inventory.ContainsSet(remoteInventory.Inventory) {
		message := messageUnion{
			Type:   "registry_request",
			Sender: (*d.listPtr).LocalNode(),
		}
		marshalled, err := json.Marshal(message)
		if err != nil {
			fmt.Println("error in registry request marshalling:", err.Error())
			return
		}
		err = (*d.listPtr).SendBestEffort(remoteInventory.Sender, marshalled)
		if err != nil {
			fmt.Println("error in sending registry request to", remoteInventory.Sender.Name, err.Error())
		}
		fmt.Println("sent registry request to", remoteInventory.Sender.Name)
	}
}

func (d memberlistAgent) NotifyJoin(node *memberlist.Node) {
	// do nothing
}

func (d memberlistAgent) NotifyLeave(node *memberlist.Node) {
	group, err := d.register()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer group.Done()

	d.lockRegistry.Lock()
	// partially free space
	d.registry[node.Name] = nil
	d.lockRegistry.Unlock()
	fmt.Println(node.Name, "has left")
}

func (d memberlistAgent) NotifyUpdate(node *memberlist.Node) {
	// do nothing
}

//----------------------------------TESTING-----------------------------------

func TestsMakeMemberlistAgent(names misc.StringSet, port int, nodes []string, test int) *memberlistAgent {
	res := MakeMemberlistAgent(names, port, nodes, config.TestsLogConfig)
	res.test = test
	res.lockHalted = &sync.Mutex{}
	return res
}

func (a *memberlistAgent) TestsHaltReturn() {
	a.list.Shutdown()
	a.lockHalted.Lock()
	a.halted = true
	a.lockHalted.Unlock()
}

func (a *memberlistAgent) testsHalt() {
	a.TestsHaltReturn()
	var block chan bool = nil
	<-block
}

func (a *memberlistAgent) testsHaltIf(t int) {
	if a.test == t {
		a.testsHalt()
	}
}
