package delegates

import (
	"errors"
	"fmt"
	"steel-lang/datastructure"
	"steel-lang/physical"

	"gobot.io/x/gobot/drivers/gpio"
)

type Button struct {
	name string
	*gpio.ButtonDriver
}

func MakeButton(adaptor physical.IOAdaptor, name string, args ...interface{}) (physical.IOdelegate, datastructure.Resources, error) {
	if len(args) != 1 {
		return physical.MakeLazyDelegate(), datastructure.MakeResources(), errors.New("button constructor invocation should have 3 arguments")
	}
	pin, ok := args[0].(string)
	if !ok {
		return physical.MakeLazyDelegate(), datastructure.MakeResources(), errors.New("third argument of button constructor should be a string specifying a pin")
	}
	resources := datastructure.MakeResources()
	resources.Bool[name] = false
	return Button{name: name, ButtonDriver: gpio.NewButtonDriver(adaptor, pin)}, resources, nil
}

func (b Button) Start(adaptor physical.IOAdaptor, inputs chan<- string, errors chan<- error) error {
	err := b.ButtonDriver.Start()
	if err != nil {
		return err
	}
	go b.getButtonInput(inputs, errors)
	return nil
}

func (b Button) Modified(adaptor physical.IOAdaptor, name string, resources datastructure.Resources, errors chan<- error) *datastructure.Resources {
	return nil
}

func (b Button) getButtonInput(in chan<- string, errs chan<- error) {
	events := b.Subscribe()
	status := false
	push := b.name + " = true;"
	release := b.name + " = false;"
	event := <-events
	for {
		var inputs chan<- string = nil
		var action string
		switch event.Name {
		case gpio.ButtonPush:
			action = push
			if !status {
				inputs = in
			}
		case gpio.ButtonRelease:
			action = release
			if status {
				inputs = in
			}
		case gpio.Error:
			errs <- fmt.Errorf("input error on button %s, received: %v", b.name, event.Data)
		}
		select {
		case inputs <- action:
			status = !status
		case event = <-events:
		}
	}
}
