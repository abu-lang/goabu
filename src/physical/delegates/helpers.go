package delegates

import "steel-lang/physical"

func MakeIOResources(a physical.IOAdaptor) *physical.IOResources {
	res := physical.MakeEmptyIOResources(a)
	res.AddOutputFrame("DigitalPin", MakeDigitalPin)
	res.AddOutputFrame("Motor", MakeMotor)
	res.AddInputFrame("Button", MakeButton)
	return res
}
