package datastructure_test

import (
	"steel-lang/datastructure"
	"testing"
	"time"
)

func TestMakeStringSet(t *testing.T) {
	s1 := datastructure.MakeStringSet("")
	s2 := datastructure.MakeStringSet("Lorem")
	s3 := datastructure.MakeStringSet("ipsum,dolor,42")
	s4 := datastructure.MakeStringSet(",,,")
	s5 := datastructure.MakeStringSet(",amet")
	s6 := datastructure.MakeStringSet("consectetur,consectetur")
	tests := []struct {
		index    int
		set      datastructure.StringSet
		empty    bool
		size     int
		elements []string
	}{
		//  {_, set, empty, size, elements},
		{1, s1, true, 0, []string{}},
		{2, s2, false, 1, []string{"Lorem"}},
		{3, s3, false, 3, []string{"ipsum", "dolor", "42"}},
		{4, s4, false, 1, []string{""}},
		{5, s5, false, 2, []string{"", "amet"}},
		{6, s6, false, 1, []string{"consectetur"}},
	}
	for _, test := range tests {
		if test.set.Empty() != test.empty {
			t.Errorf("TestMakeStringSet #%d failed: emptiness", test.index)
		}
		if test.set.Size() != test.size {
			t.Errorf("TestMakeStringSet #%d failed: size", test.index)
		}
		for _, el := range test.elements {
			if !test.set.Contains(el) {
				t.Errorf("TestMakeStringSet #%d failed: set does not contains \"%s\"", test.index, el)
			}
		}
	}
}

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
