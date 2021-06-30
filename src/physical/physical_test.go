package physical

import (
	"fmt"
	"testing"

	"gobot.io/x/gobot/platforms/raspi"
)

func TestAddLed(t *testing.T) {
	resources := MakeIOResources(raspi.NewAdaptor())
	tests := []struct {
		index int
		led   string
		pin   string
		good  bool
	}{
		//  {_, led, pin, good},
		{1, "lorem", "3", true},
		{2, "Ipsum_123", "10", true},
		{3, "dolor__", "8", true},
		{4, "Ipsum_123", "5", false},
		{5, "lorem", "12", false},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("TestAddLed#%d", test.index), func(t *testing.T) {
			err := resources.AddLed(test.led, test.pin)
			if err != nil {
				if test.good {
					t.Error(err.Error())
				}
				return
			}
			if !test.good {
				t.Error("AddLed should return an error")
			}
			checkAdded(t, resources, test.led, devLed)
			p, present := resources.ledPins[test.led]
			if !present {
				t.Errorf("%s: missing pin", test.led)
			}
			if p != test.pin {
				t.Errorf("%s: pin should be %s", test.led, test.pin)
			}
		})
	}
}

func TestAddButton(t *testing.T) {
	resources := MakeIOResources(raspi.NewAdaptor())
	tests := []struct {
		index  int
		button string
		pin    string
		good   bool
	}{
		//  {_, button, pin, good},
		{1, "consectetur", "22", true},
		{2, "elit__456", "24", true},
		{3, "consectetur", "32", false},
		{4, "elit__456", "28", false},
		{5, "Adipiscing_7", "26", true},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("TestAddButton#%d", test.index), func(t *testing.T) {
			err := resources.AddButton(test.button, test.pin)
			if err != nil {
				if test.good {
					t.Error(err.Error())
				}
				return
			}
			if !test.good {
				t.Error("AddButton should return an error")
			}
			checkAdded(t, resources, test.button, devButton)
			d, present := resources.buttons[test.button]
			if !present {
				t.Errorf("%s: missing ButtonDriver", test.button)
			}
			if d.Pin() != test.pin {
				t.Errorf("%s: pin should be %s", test.button, test.pin)
			}
		})
	}
}

func checkAdded(t *testing.T, resources IOResources, r string, tp int) {
	t.Helper()
	devStr := ""
	switch tp {
	case devLed:
		devStr = "devLed"
	case devButton:
		devStr = "devButton"
	}
	if !resources.Has(r) {
		t.Errorf("%s: missing resource", r)
	}
	d, present := resources.devices[r]
	if !present {
		t.Errorf("%s: missing type", r)
	}
	if d != tp {
		t.Errorf("%s: type should be %s", r, devStr)
	}
}
