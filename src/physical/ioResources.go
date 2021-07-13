package physical

import (
	"fmt"
	"steel-lang/datastructure"

	"gobot.io/x/gobot/drivers/gpio"
)

const (
	devLed = iota
	devMotor
	devButton
)

type motor struct {
	writer      gpio.PwmWriter
	forwardPin  string
	backwardPin string
}

type IOResources struct {
	datastructure.Resources
	adaptor IOAdaptor
	inputs  chan string
	devices map[string]int
	ledPins map[string]string
	buttons map[string]*gpio.ButtonDriver
	motors  map[string]motor
}

func MakeIOResources(a IOAdaptor) IOResources {
	return IOResources{
		Resources: datastructure.MakeResources(),
		adaptor:   a,
		inputs:    make(chan string),
		devices:   make(map[string]int),
		ledPins:   make(map[string]string),
		buttons:   make(map[string]*gpio.ButtonDriver),
		motors:    make(map[string]motor),
	}
}

func (i IOResources) Start() error {
	err := i.adaptor.Connect()
	if err != nil {
		return err
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

func (i IOResources) Inputs() <-chan string {
	return i.inputs
}

func (i IOResources) InputsNumber() int {
	return len(i.buttons)
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
		pin := i.ledPins[r]
		if i.Resources.Bool[r] {
			err := i.adaptor.DigitalWrite(pin, 1)
			if err != nil {
				i.Resources.Bool[r] = false
				return err
			}
		} else {
			err := i.adaptor.DigitalWrite(pin, 0)
			if err != nil {
				i.Resources.Bool[r] = true
				return err
			}
		}
	case devMotor:
		speed := i.Integer[r]
		forward := speed >= 0
		if !forward {
			speed = speed * -1
		}
		if speed > 255 {
			speed = 255
		}
		return i.motors[r].set(byte(speed), forward)
	}
	return nil
}

func (i IOResources) Clone() datastructure.ResourceController {
	return IOResources{
		Resources: i.Resources.Clone().GetResources(),
		adaptor:   i.adaptor,
		inputs:    i.inputs,
		devices:   i.devices,
		ledPins:   i.ledPins,
		buttons:   i.buttons,
		motors:    i.motors,
	}
}

func (i IOResources) AddLed(r string, pin string) error {
	if i.Has(r) {
		return fmt.Errorf("there is already a resource named: %s", r)
	}
	i.Bool[r] = false
	i.devices[r] = devLed
	i.ledPins[r] = pin
	return nil
}

func (i IOResources) AddMotor(r string, forwardPin string, backwardPin string) error {
	if i.Has(r) {
		return fmt.Errorf("there is already a resource named: %s", r)
	}
	i.Integer[r] = 0
	i.devices[r] = devMotor
	i.motors[r] = motor{
		writer:      i.adaptor,
		forwardPin:  forwardPin,
		backwardPin: backwardPin,
	}
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

func getButtonInput(name string, button *gpio.ButtonDriver, in chan<- string) {
	events := button.Subscribe()
	status := false
	push := name + " = true;"
	release := name + " = false;"
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
			fmt.Println(event.Data)
		}
		select {
		case inputs <- action:
			status = !status
		case event = <-events:
		}
	}
}

func (m motor) set(speed byte, forward bool) error {
	err := m.writer.PwmWrite(m.forwardPin, 0)
	if err != nil {
		return err
	}
	err = m.writer.PwmWrite(m.backwardPin, 0)
	if err != nil || speed == 0 {
		return err
	}
	if forward {
		return m.writer.PwmWrite(m.forwardPin, speed)
	} else {
		return m.writer.PwmWrite(m.backwardPin, speed)
	}
}
