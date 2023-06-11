// Copyright 2021 Massimo Comuzzo, Michele Pasqua and Marino Miculan
// SPDX-License-Identifier: Apache-2.0

package physical

import (
	"gobot.io/x/gobot/v2"
	"gobot.io/x/gobot/v2/drivers/gpio"
)

// TODO check if sufficient
type IOadaptor interface {
	gobot.Adaptor
	gpio.DigitalReader
	gpio.DigitalWriter
	gpio.PwmWriter
}
