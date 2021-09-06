package iodelegates

import (
	"errors"
	"steel/memory"
	"steel/physical"
)

type DigitalPin struct {
	pin string
}

func MakeDigitalPin(adaptor physical.IOadaptor, name string, args ...interface{}) (physical.IOdelegate, memory.Resources, error) {
	if len(args) != 1 {
		return nil, memory.MakeResources(), errors.New("digitalPin constructor invocation should have 3 arguments")
	}
	pin, ok := args[0].(string)
	if !ok {
		return nil, memory.MakeResources(), errors.New("third argument of digitalPin constructor should be a string specifying a pin")
	}
	resources := memory.MakeResources()
	resources.Bool[name] = false
	return DigitalPin{pin: pin}, resources, nil
}

func (p DigitalPin) Start(adaptor physical.IOadaptor, inputs chan<- string, errors chan<- error) error {
	return nil
}

func (p DigitalPin) Modified(adaptor physical.IOadaptor, name string, resources memory.Resources, errors chan<- error) *memory.Resources {
	if resources.Bool[name] {
		err := adaptor.DigitalWrite(p.pin, 1)
		if err != nil {
			panic(err)
		}
	} else {
		err := adaptor.DigitalWrite(p.pin, 0)
		if err != nil {
			panic(err)
		}
	}
	return nil
}
