package datastructure

import (
	"fmt"
	"steel-lang/misc"
	"time"
)

type Resources struct {
	Bool    map[string]bool
	Integer map[string]int64
	Float   map[string]float64
	Text    map[string]string
	Time    map[string]time.Time
	Other   map[string]interface{}
}

func MakeResources() Resources {
	return Resources{
		Bool:    make(map[string]bool),
		Integer: make(map[string]int64),
		Float:   make(map[string]float64),
		Text:    make(map[string]string),
		Time:    make(map[string]time.Time),
		Other:   make(map[string]interface{}),
	}
}

func (r Resources) IsValid() bool {
	atts := NewStringSet("")
	for a := range r.Bool {
		if atts.Contains(a) {
			return false
		}
		atts.Insert(a)
	}
	for a := range r.Integer {
		if atts.Contains(a) {
			return false
		}
		atts.Insert(a)
	}
	for a := range r.Float {
		if atts.Contains(a) {
			return false
		}
		atts.Insert(a)
	}
	for a := range r.Text {
		if atts.Contains(a) {
			return false
		}
		atts.Insert(a)
	}
	for a := range r.Time {
		if atts.Contains(a) {
			return false
		}
		atts.Insert(a)
	}
	for a := range r.Other {
		if atts.Contains(a) {
			return false
		}
		atts.Insert(a)
	}
	return atts.AllMatch(`[a-zA-Z]+[a-zA-Z0-9_]*`)
}

func (r Resources) GetTypes() map[string]string {
	res := make(map[string]string)
	for a := range r.Bool {
		res[a] = "Bool"
	}
	for a := range r.Integer {
		res[a] = "Integer"
	}
	for a := range r.Float {
		res[a] = "Float"
	}
	for a := range r.Text {
		res[a] = "Text"
	}
	for a := range r.Time {
		res[a] = "Time"
	}
	for a := range r.Other {
		res[a] = "Other"
	}
	return res
}

func (r Resources) String() string {
	var str string = "[ "
	for key, value := range r.Bool {
		str = str + fmt.Sprintf("(%T)%s->%v ", value, key, value)
	}
	for key, value := range r.Integer {
		str = str + fmt.Sprintf("(%T)%s->%v ", value, key, value)
	}
	for key, value := range r.Float {
		str = str + fmt.Sprintf("(%T)%s->%v ", value, key, value)
	}
	for key, value := range r.Text {
		str = str + fmt.Sprintf("(%T)%s->%v ", value, key, value)
	}
	for key, value := range r.Time {
		str = str + fmt.Sprintf("(%T)%s->%v ", value, key, value)
	}
	for key, value := range r.Other {
		str = str + fmt.Sprintf("(%T)%s->%v ", value, key, value)
	}
	return str + "]"
}

func (r Resources) Clone() Resources {
	res := MakeResources()
	for k, v := range r.Bool {
		res.Bool[k] = v
	}
	for k, v := range r.Integer {
		res.Integer[k] = v
	}
	for k, v := range r.Float {
		res.Float[k] = v
	}
	for k, v := range r.Text {
		res.Text[k] = v
	}
	for k, v := range r.Time {
		res.Time[k] = v
	}
	res.Other = misc.CopyMap(r.Other)
	return res
}
