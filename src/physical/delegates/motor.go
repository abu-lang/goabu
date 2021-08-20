package delegates

import (
	"errors"
	"steel-lang/datastructure"
	"steel-lang/physical"

	"gobot.io/x/gobot/drivers/gpio"
)

type motor struct {
	forwardPin  string
	backwardPin string
}

func MakeMotor(adaptor physical.IOAdaptor, name string, args ...interface{}) (physical.IOdelegate, datastructure.Resources, error) {
	if len(args) != 2 {
		return physical.MakeLazyDelegate(), datastructure.MakeResources(), errors.New("motor constructor invocation should have 4 arguments")
	}
	forward, ok := args[0].(string)
	if !ok {
		return physical.MakeLazyDelegate(), datastructure.MakeResources(), errors.New("third argument of motor constructor should be a string specifying a pin")
	}
	backward, ok := args[1].(string)
	if !ok {
		return physical.MakeLazyDelegate(), datastructure.MakeResources(), errors.New("fourth argument of motor constructor should be a string specifying a pin")
	}
	resources := datastructure.MakeResources()
	resources.Integer[name] = 0
	return motor{forwardPin: forward, backwardPin: backward}, resources, nil
}

func (m motor) Start(adaptor physical.IOAdaptor, inputs chan<- string, errors chan<- error) error {
	return nil
}

func (m motor) Modified(adaptor physical.IOAdaptor, name string, resources datastructure.Resources, errors chan<- error) *datastructure.Resources {
	speed := resources.Integer[name]
	forward := speed >= 0
	if !forward {
		speed = speed * -1
	}
	if speed > 255 {
		speed = 255
	}
	err := m.set(adaptor, byte(speed), forward)
	if err != nil {
		panic(err)
	}
	return nil
}

func (m motor) set(writer gpio.PwmWriter, speed byte, forward bool) error {
	err := writer.PwmWrite(m.forwardPin, 0)
	if err != nil {
		return err
	}
	err = writer.PwmWrite(m.backwardPin, 0)
	if err != nil || speed == 0 {
		return err
	}
	if forward {
		return writer.PwmWrite(m.forwardPin, speed)
	} else {
		return writer.PwmWrite(m.backwardPin, speed)
	}
}
