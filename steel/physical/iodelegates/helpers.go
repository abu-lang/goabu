package iodelegates

import "steel/physical"

func MakeIOresources(a physical.IOadaptor) *physical.IOresources {
	res := physical.MakeEmptyIOresources(a)
	res.AddOutputFrame("DigitalPin", MakeDigitalPin)
	res.AddOutputFrame("Motor", MakeMotor)
	res.AddInputFrame("Button", MakeButton)
	return res
}
