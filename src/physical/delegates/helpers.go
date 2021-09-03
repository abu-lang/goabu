package delegates

import "steel-lang/physical"

func MakeIOresources(a physical.IOadaptor) *physical.IOresources {
	res := physical.MakeEmptyIOresources(a)
	res.AddOutputFrame("DigitalPin", MakeDigitalPin)
	res.AddOutputFrame("Motor", MakeMotor)
	res.AddInputFrame("Button", MakeButton)
	return res
}
