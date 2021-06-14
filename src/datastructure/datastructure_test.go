package datastructure_test

import (
	"steel-lang/datastructure"
	"testing"
	"time"
)

func TestIsValid(t *testing.T) {
	r1 := datastructure.MakeResources()

	r2 := datastructure.MakeResources()
	r2.Integer["Lorem"] = 42
	r2.Text["ipsum"] = "dolor"

	r3 := datastructure.MakeResources()
	r3.Bool["x"] = false
	r3.Float["y"] = 3.14
	r3.Other["x"] = datastructure.MakeResources()

	r4 := datastructure.MakeResources()
	r4.Integer[""] = 123

	r5 := datastructure.MakeResources()
	r5.Time["10sit"] = time.Now()

	r6 := datastructure.MakeResources()
	r6.Text["a,met"] = ""

	tests := []struct {
		index     int
		resources datastructure.Resources
		isValid   bool
	}{
		//  {_, resources, isValid},
		{1, r1, true},
		{2, r2, true},
		{3, r3, false},
		{4, r4, false},
		{5, r5, false},
		{6, r6, false},
	}
	for _, test := range tests {
		if test.resources.IsValid() != test.isValid {
			t.Errorf("TestIsValid #%d failed", test.index)
		}
	}
}
