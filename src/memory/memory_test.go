package memory_test

import (
	"steel-lang/memory"
	"testing"
	"time"
)

func TestHasDuplicates(t *testing.T) {
	r1 := memory.MakeResources()

	r2 := memory.MakeResources()
	r2.Integer["Lorem"] = 42
	r2.Text["ipsum"] = "dolor"

	r3 := memory.MakeResources()
	r3.Bool["x"] = false
	r3.Float["y"] = 3.14
	r3.Other["x"] = memory.MakeResources()

	r4 := memory.MakeResources()
	r4.Integer[""] = 123
	r4.Text[""] = "456"

	r5 := memory.MakeResources()
	r5.Time["10sit"] = time.Now()
	r5.Float["10sit"] = 3.14
	r5.Other["10sit"] = struct{}{}

	r6 := memory.MakeResources()
	r6.Text["a,met"] = ""
	r6.Bool["z"] = true

	tests := []struct {
		index         int
		resources     memory.Resources
		hasDuplicates bool
	}{
		//  {_, resources, HasDuplicates},
		{1, r1, false},
		{2, r2, false},
		{3, r3, true},
		{4, r4, true},
		{5, r5, true},
		{6, r6, false},
	}
	for _, test := range tests {
		if test.resources.HasDuplicates() != test.hasDuplicates {
			t.Errorf("TestHasDuplicates#%d failed", test.index)
		}
	}
}
