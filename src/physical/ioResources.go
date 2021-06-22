package physical

import (
	"fmt"
	"steel-lang/datastructure"

	"gobot.io/x/gobot/drivers/gpio"
)

const (
	devLed = iota
	devButton
)

type IOResources struct {
	datastructure.Resources
	adaptor IOAdaptor
	inputs  chan datastructure.Action
	devices map[string]int
	leds    map[string]*gpio.LedDriver
	buttons map[string]*gpio.ButtonDriver
}

func MakeIOResources(a IOAdaptor) IOResources {
	return IOResources{
		Resources: datastructure.MakeResources(),
		adaptor:   a,
		inputs:    make(chan datastructure.Action),
		devices:   make(map[string]int),
		leds:      make(map[string]*gpio.LedDriver),
		buttons:   make(map[string]*gpio.ButtonDriver),
	}
}

func (i IOResources) Start() error {
	err := i.adaptor.Connect()
	if err != nil {
		return err
	}
	for _, l := range i.leds {
		err = l.Start()
		if err != nil {
			return err
		}
	}
	for k, b := range i.buttons {
		err = b.Start()
		if err != nil {
			return err
		}
		go getButtonInput(k, b, i.inputs)
	}
	return nil
}

func (i IOResources) Inputs() <-chan datastructure.Action {
	return i.inputs
}

func (i IOResources) Modified(r string) error {
	if !i.Has(r) {
		return fmt.Errorf("no resource is named: %s", r)
	}
	t, present := i.devices[r]
	if !present {
		return nil
	}
	switch t {
	case devLed:
		if i.Resources.Bool[r] {
			i.leds[r].On()
		} else {
			i.leds[r].Off()
		}
	}
	return nil
}

func (i IOResources) Clone() datastructure.ResourceController {
	return IOResources{
		Resources: i.Resources.Clone().GetResources(),
		adaptor:   i.adaptor,
		inputs:    i.inputs,
		devices:   i.devices,
		leds:      i.leds,
		buttons:   i.buttons,
	}
}

func (i IOResources) AddLed(r string, pin string) error {
	if i.Has(r) {
		return fmt.Errorf("there is already a resource named: %s", r)
	}
	i.Bool[r] = false
	i.devices[r] = devLed
	i.leds[r] = gpio.NewLedDriver(i.adaptor, pin)
	return nil
}

func (i IOResources) AddButton(r string, pin string) error {
	if i.Has(r) {
		return fmt.Errorf("there is already a resource named: %s", r)
	}
	i.Bool[r] = false
	i.devices[r] = devButton
	i.buttons[r] = gpio.NewButtonDriver(i.adaptor, pin)
	return nil
}

func getButtonInput(name string, button *gpio.ButtonDriver, in chan<- datastructure.Action) {
	events := button.Subscribe()
	status := false
	event := <-events
	for {
		var inputs chan<- datastructure.Action = nil
		var action datastructure.Action
		switch event.Name {
		case gpio.ButtonPush:
			action = datastructure.Action{Resource: name, Expression: "true"}
			if !status {
				inputs = in
			}
		case gpio.ButtonRelease:
			action = datastructure.Action{Resource: name, Expression: "false"}
			if status {
				inputs = in
			}
		case gpio.Error:
			fmt.Println(event.Data)
		}
		select {
		case inputs <- action:
			status = !status
		case event = <-events:
		}
	}
}
