package datastructure

import (
	"fmt"
	"steel-lang/stringset"
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

func (r Resources) Start() error {
	return nil
}

func (r Resources) Inputs() <-chan string {
	return nil
}

func (r Resources) Errors() <-chan error {
	return nil
}

func (r Resources) Modified(resource string) {}

func (r Resources) InputsNumber() int {
	return 0
}

func (r Resources) IsValid() bool {
	atts := stringset.Make("")
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
	return atts.AllMatch(`\A[a-zA-Z]+[a-zA-Z0-9_]*\z`)
}

func (r Resources) Has(resource string) bool {
	return r.ResourceNames().Contains(resource)
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

func (r Resources) GetResources() Resources {
	return r
}

func (r Resources) ResourceNames() stringset.StringSet {
	atts := stringset.Make("")
	for a := range r.Bool {
		atts.Insert(a)
	}
	for a := range r.Integer {
		atts.Insert(a)
	}
	for a := range r.Float {
		atts.Insert(a)
	}
	for a := range r.Text {
		atts.Insert(a)
	}
	for a := range r.Time {
		atts.Insert(a)
	}
	for a := range r.Other {
		atts.Insert(a)
	}
	return atts
}

func (r Resources) Extract(s stringset.StringSet) Resources {
	res := MakeResources()
	for k, v := range r.Bool {
		if s.Contains(k) {
			res.Bool[k] = v
		}
	}
	for k, v := range r.Integer {
		if s.Contains(k) {
			res.Integer[k] = v
		}
	}
	for k, v := range r.Float {
		if s.Contains(k) {
			res.Float[k] = v
		}
	}
	for k, v := range r.Text {
		if s.Contains(k) {
			res.Text[k] = v
		}
	}
	for k, v := range r.Time {
		if s.Contains(k) {
			res.Time[k] = v
		}
	}
	for k, v := range r.Other {
		if s.Contains(k) {
			res.Other[k] = v
		}
	}
	return res
}

func (r Resources) Enclose(i Resources) {
	for k, v := range i.Bool {
		r.Bool[k] = v
	}
	for k, v := range i.Integer {
		r.Integer[k] = v
	}
	for k, v := range i.Float {
		r.Float[k] = v
	}
	for k, v := range i.Text {
		r.Text[k] = v
	}
	for k, v := range i.Time {
		r.Time[k] = v
	}
	for k, v := range i.Other {
		r.Other[k] = v
	}
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

func (r Resources) Copy() ResourceController {
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
	for k, v := range r.Other {
		res.Other[k] = v
	}
	return res
}
