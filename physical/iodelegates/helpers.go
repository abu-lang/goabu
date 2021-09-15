// Package iodelegates implements the behaviour of some GoAbU I/O related resources.
package iodelegates

import "github.com/abu-lang/goabu/physical"

func MakeIOresources(a physical.IOadaptor) *physical.IOresources {
	res := physical.MakeEmptyIOresources(a)
	res.AddOutputFrame("DigitalPin", MakeDigitalPin)
	res.AddOutputFrame("Motor", MakeMotor)
	res.AddInputFrame("Button", MakeButton)
	return res
}
