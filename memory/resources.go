// Copyright 2021 Massimo Comuzzo, Michele Pasqua and Marino Miculan
// SPDX-License-Identifier: Apache-2.0

package memory

import (
	"fmt"
	"time"

	"github.com/abu-lang/goabu/stringset"
)

// Resources is a struct implementing the [ResourceController] interface modeling the state of a node
// that has no sensors nor actuators.
type Resources struct {
	Bool    map[string]bool
	Integer map[string]int64
	Float   map[string]float64
	Text    map[string]string
	Time    map[string]time.Time
	Other   map[string]interface{}
}

// MakeResources returns a new empty [Resources] struct.
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

// Start is a no-op.
func (r Resources) Start() error {
	return nil
}

// Inputs returns nil.
func (r Resources) Inputs() <-chan string {
	return nil
}

// Errors returns nil.
func (r Resources) Errors() <-chan error {
	return nil
}

// Modified is a no-op.
func (r Resources) Modified(resource string) {}

// HasDuplicates verifies if there are multiple resources sharing the same identifier.
func (r Resources) HasDuplicates() bool {
	atts := stringset.Make()
	for a := range r.Bool {
		if atts.Has(a) {
			return true
		}
		atts.Insert(a)
	}
	for a := range r.Integer {
		if atts.Has(a) {
			return true
		}
		atts.Insert(a)
	}
	for a := range r.Float {
		if atts.Has(a) {
			return true
		}
		atts.Insert(a)
	}
	for a := range r.Text {
		if atts.Has(a) {
			return true
		}
		atts.Insert(a)
	}
	for a := range r.Time {
		if atts.Has(a) {
			return true
		}
		atts.Insert(a)
	}
	for a := range r.Other {
		if atts.Has(a) {
			return true
		}
		atts.Insert(a)
	}
	return false
}

// Has checks if there is a resource identified by the provided string.
func (r Resources) Has(resource string) bool {
	_, present := r.Bool[resource]
	if present {
		return true
	}
	_, present = r.Integer[resource]
	if present {
		return true
	}
	_, present = r.Float[resource]
	if present {
		return true
	}
	_, present = r.Text[resource]
	if present {
		return true
	}
	_, present = r.Time[resource]
	if present {
		return true
	}
	_, present = r.Other[resource]
	return present
}

// Types returns a map with an entry for each resource specifying its type
// (one of the following: "Bool", "Integer", "Float", "Text", "Time", "Other").
//
// Prerequisite: !HasDuplicates()
func (r Resources) Types() map[string]string {
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

// GetResources returns the [Resources] struct itself.
func (r Resources) GetResources() Resources {
	return r
}

// ResourceNames returns the list of all the contained resources' identifiers (without repeated elements).
func (r Resources) ResourceNames() []string {
	atts := stringset.Make()
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
	return atts.Slice()
}

// Extract returns a shallow copy of only the resources specified by the provided identifiers.
func (r Resources) Extract(resources []string) Resources {
	s := stringset.Make(resources...)
	res := MakeResources()
	for k, v := range r.Bool {
		if s.Has(k) {
			res.Bool[k] = v
		}
	}
	for k, v := range r.Integer {
		if s.Has(k) {
			res.Integer[k] = v
		}
	}
	for k, v := range r.Float {
		if s.Has(k) {
			res.Float[k] = v
		}
	}
	for k, v := range r.Text {
		if s.Has(k) {
			res.Text[k] = v
		}
	}
	for k, v := range r.Time {
		if s.Has(k) {
			res.Time[k] = v
		}
	}
	for k, v := range r.Other {
		if s.Has(k) {
			res.Other[k] = v
		}
	}
	return res
}

// Enclose adds the provided resources to the ResourceController, overwriting previous values if present.
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

// String returns a string representation of the struct for debugging purposes.
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

// Copy returns a shallow copy of the struct.
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
