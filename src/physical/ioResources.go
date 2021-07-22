package physical

import (
	"errors"
	"fmt"
	"steel-lang/datastructure"
	"steel-lang/misc"
)

type ioResourceMeta struct {
	resourceType string
	isInput      bool
	isOutput     bool
}

type frame struct {
	constructor func(IOAdaptor, string, ...interface{}) (IOdelegate, datastructure.Resources, error)
	ioResourceMeta
}

type resource struct {
	meta ioResourceMeta
	IOdelegate
	managed misc.StringSet
}

type IOResources struct {
	datastructure.Resources
	adaptor   IOAdaptor
	inputs    chan string
	delegates []*resource
	managers  map[string]*resource
	frames    map[string]frame
}

func MakeEmptyIOResources(a IOAdaptor) *IOResources {
	return &IOResources{
		Resources: datastructure.MakeResources(),
		adaptor:   a,
		inputs:    make(chan string),
		managers:  make(map[string]*resource),
		frames:    make(map[string]frame),
	}
}

func (i *IOResources) Start() error {
	err := i.adaptor.Connect()
	if err != nil {
		return err
	}
	for _, r := range i.delegates {
		err = r.Start(i.adaptor, i.inputs)
		if err != nil {
			return err
		}
	}
	return nil
}

func (i *IOResources) Inputs() <-chan string {
	return i.inputs
}

func (i *IOResources) InputsNumber() int {
	res := 0
	for _, r := range i.delegates {
		if r.meta.isInput {
			res++
		}
	}
	return res
}

func (i *IOResources) Modified(r string) {
	if !i.Has(r) {
		fmt.Printf("no resource is named: %s\n", r)
		return
	}
	resource, present := i.managers[r]
	if !present {
		return
	}
	mods := resource.Modified(i.adaptor, r, i.Resources.Extract(resource.managed))
	if mods != nil {
		i.Resources.Enclose(mods.Extract(resource.managed))
	}
}

func (i *IOResources) Clone() datastructure.ResourceController {
	return &IOResources{
		Resources: i.Resources.Clone().GetResources(),
		adaptor:   i.adaptor,
		inputs:    i.inputs,
		delegates: i.delegates,
		managers:  i.managers,
		frames:    i.frames,
	}
}

func (i *IOResources) AddInputFrame(t string, c func(IOAdaptor, string, ...interface{}) (IOdelegate, datastructure.Resources, error)) error {
	return i.addFrame(true, false, t, c)
}

func (i *IOResources) AddOutputFrame(t string, c func(IOAdaptor, string, ...interface{}) (IOdelegate, datastructure.Resources, error)) error {
	return i.addFrame(false, true, t, c)
}

func (i *IOResources) AddInputOutputFrame(t string, c func(IOAdaptor, string, ...interface{}) (IOdelegate, datastructure.Resources, error)) error {
	return i.addFrame(true, true, t, c)
}

func (i *IOResources) Add(t string, name string, args ...interface{}) error {
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
	newNames := newResources.ResourceNames()
	if i.ResourceNames().IntersectsWith(newNames) {
		return errors.New("conflict in resource names")
	}
	i.Enclose(newResources)
	resource.IOdelegate = delegate
	resource.managed = newNames
	i.delegates = append(i.delegates, resource)
	for k := range newNames {
		i.managers[k] = resource
	}
	return nil
}

func (i *IOResources) addFrame(input, output bool, t string, c func(IOAdaptor, string, ...interface{}) (IOdelegate, datastructure.Resources, error)) error {
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
