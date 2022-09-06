// Package physical implements GoAbU resources associated with sensors and actuators.
package physical

import (
	"errors"
	"fmt"

	"github.com/abu-lang/goabu/memory"
	"github.com/abu-lang/goabu/stringset"
)

type ioResourceMeta struct {
	resourceType string
	isInput      bool
	isOutput     bool
}

type frame struct {
	constructor func(IOadaptor, string, ...interface{}) (IOdelegate, memory.Resources, error)
	ioResourceMeta
}

type resource struct {
	meta ioResourceMeta
	IOdelegate
	managed []string
}

type IOresources struct {
	memory.Resources
	adaptor   IOadaptor
	inputs    chan string
	errors    chan error
	delegates []*resource
	managers  map[string]*resource
	frames    map[string]frame
}

func MakeEmptyIOresources(a IOadaptor) *IOresources {
	return &IOresources{
		Resources: memory.MakeResources(),
		adaptor:   a,
		inputs:    make(chan string),
		errors:    make(chan error),
		managers:  make(map[string]*resource),
		frames:    make(map[string]frame),
	}
}

func (i *IOresources) Start() error {
	err := i.adaptor.Connect()
	if err != nil {
		return err
	}
	for _, r := range i.delegates {
		err = r.Start(i.adaptor, i.inputs, i.errors)
		if err != nil {
			return err
		}
	}
	return nil
}

func (i *IOresources) Inputs() <-chan string {
	return i.inputs
}

func (i *IOresources) Errors() <-chan error {
	return i.errors
}

func (i *IOresources) InputsNumber() int {
	res := 0
	for _, r := range i.delegates {
		if r.meta.isInput {
			res++
		}
	}
	return res
}

func (i *IOresources) Modified(r string) {
	if !i.Has(r) {
		return
	}
	resource, present := i.managers[r]
	if !present {
		return
	}
	mods := resource.Modified(i.adaptor, r, i.Resources.Extract(resource.managed), i.errors)
	if mods != nil {
		i.Resources.Enclose(mods.Extract(resource.managed))
	}
}

func (i *IOresources) Copy() memory.ResourceController {
	return &IOresources{
		Resources: i.Resources.Copy().GetResources(),
		adaptor:   i.adaptor,
		inputs:    i.inputs,
		delegates: i.delegates,
		managers:  i.managers,
		frames:    i.frames,
	}
}

func (i *IOresources) AddInputFrame(t string, c func(IOadaptor, string, ...interface{}) (IOdelegate, memory.Resources, error)) error {
	return i.addFrame(true, false, t, c)
}

func (i *IOresources) AddOutputFrame(t string, c func(IOadaptor, string, ...interface{}) (IOdelegate, memory.Resources, error)) error {
	return i.addFrame(false, true, t, c)
}

func (i *IOresources) AddInputOutputFrame(t string, c func(IOadaptor, string, ...interface{}) (IOdelegate, memory.Resources, error)) error {
	return i.addFrame(true, true, t, c)
}

func (i *IOresources) Add(t string, name string, args ...interface{}) error {
	frame, present := i.frames[t]
	if !present {
		return fmt.Errorf("no frame for type %s", t)
	}
	resource := &resource{
		meta: frame.ioResourceMeta,
	}
	delegate, newResources, err := frame.constructor(i.adaptor, name, args...)
	if err != nil {
		return err
	}
	if delegate == nil {
		return errors.New("created IOdelegate is nil")
	}
	newResources = nestResources(name, newResources)
	managed := newResources.ResourceNames()
	newNames := stringset.Make(managed...)
	for _, r := range i.ResourceNames() {
		if newNames.Has(r) {
			return errors.New("conflict in resource names")
		}
	}
	i.Enclose(newResources)
	resource.IOdelegate = delegate
	resource.managed = managed
	i.delegates = append(i.delegates, resource)
	for _, k := range managed {
		i.managers[k] = resource
	}
	return nil
}

func (i *IOresources) addFrame(input, output bool, t string, c func(IOadaptor, string, ...interface{}) (IOdelegate, memory.Resources, error)) error {
	_, present := i.frames[t]
	if present {
		return fmt.Errorf("there is already a frame for type %s", t)
	}
	i.frames[t] = frame{
		constructor: c,
		ioResourceMeta: ioResourceMeta{
			resourceType: t,
			isInput:      input,
			isOutput:     output,
		},
	}
	return nil
}

// nestResources returns the Resources argument if it contains at most one resource
// otherwise it returns a Resources struct where the names of the resources are prefixed
// with the string argument and an '_'.
func nestResources(name string, r memory.Resources) memory.Resources {
	if len(r.ResourceNames()) < 2 {
		return r
	}
	res := memory.MakeResources()
	for k, v := range r.Bool {
		res.Bool[name+"_"+k] = v
	}
	for k, v := range r.Integer {
		res.Integer[name+"_"+k] = v
	}
	for k, v := range r.Float {
		res.Float[name+"_"+k] = v
	}
	for k, v := range r.Text {
		res.Text[name+"_"+k] = v
	}
	for k, v := range r.Time {
		res.Time[name+"_"+k] = v
	}
	for k, v := range r.Other {
		res.Other[name+"_"+k] = v
	}
	return res
}
