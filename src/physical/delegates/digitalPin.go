package delegates

import (
	"errors"
	"steel-lang/datastructure"
	"steel-lang/physical"
)

type DigitalPin struct {
	pin string
}

func MakeDigitalPin(adaptor physical.IOAdaptor, name string, args ...interface{}) (physical.IOdelegate, datastructure.Resources, error) {
	if len(args) != 1 {
		return physical.MakeLazyDelegate(), datastructure.MakeResources(), errors.New("digitalPin constructor invocation should have 3 arguments")
	}
	pin, ok := args[0].(string)
	if !ok {
		return physical.MakeLazyDelegate(), datastructure.MakeResources(), errors.New("third argument of digitalPin constructor should be a string specifying a pin")
	}
	resources := datastructure.MakeResources()
	resources.Bool[name] = false
	return DigitalPin{pin: pin}, resources, nil
}

func (p DigitalPin) Start(adaptor physical.IOAdaptor, inputs chan<- string, errors chan<- error) error {
	return nil
}

func (p DigitalPin) Modified(adaptor physical.IOAdaptor, name string, resources datastructure.Resources, errors chan<- error) *datastructure.Resources {
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
