// Copyright 2021 Massimo Comuzzo, Michele Pasqua and Marino Miculan
// SPDX-License-Identifier: Apache-2.0

package iodelegates

import (
	"errors"
	"fmt"

	"github.com/abu-lang/goabu/memory"
	"github.com/abu-lang/goabu/physical"

	"gobot.io/x/gobot/drivers/gpio"
)

type Button struct {
	name   string
	driver *gpio.ButtonDriver
}

func MakeButton(adaptor physical.IOadaptor, name string, args ...interface{}) (physical.IOdelegate, memory.Resources, error) {
	if len(args) != 1 {
		return nil, memory.MakeResources(), errors.New("button constructor invocation should have 3 arguments")
	}
	pin, ok := args[0].(string)
	if !ok {
		return nil, memory.MakeResources(), errors.New("third argument of button constructor should be a string specifying a pin")
	}
	resources := memory.MakeResources()
	resources.Bool[name] = false
	return Button{name: name, driver: gpio.NewButtonDriver(adaptor, pin)}, resources, nil
}

func (b Button) Start(adaptor physical.IOadaptor, inputs chan<- string, errors chan<- error) error {
	err := b.driver.Start()
	if err != nil {
		return err
	}
	go b.getButtonInput(inputs, errors)
	return nil
}

func (b Button) Modified(adaptor physical.IOadaptor, name string, resources memory.Resources, errors chan<- error) *memory.Resources {
	return nil
}

func (b Button) getButtonInput(in chan<- string, errs chan<- error) {
	events := b.driver.Subscribe()
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
