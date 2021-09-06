package physical

import (
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
)

// TODO check if sufficient
type IOadaptor interface {
	gobot.Adaptor
	gpio.DigitalReader
	gpio.DigitalWriter
	gpio.PwmWriter
}
