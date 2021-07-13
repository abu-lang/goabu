package physical

import (
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
)

type IOAdaptor interface {
	gobot.Adaptor
	gpio.DigitalReader
	gpio.DigitalWriter
	gpio.PwmWriter
}
