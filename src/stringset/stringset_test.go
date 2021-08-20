package stringset_test

import (
	"steel-lang/stringset"
	"testing"
)

func TestMake(t *testing.T) {
	s1 := stringset.Make("")
	s2 := stringset.Make("Lorem")
	s3 := stringset.Make("ipsum,dolor,42")
	s4 := stringset.Make(",,,")
	s5 := stringset.Make(",amet")
	s6 := stringset.Make("consectetur,consectetur")
	tests := []struct {
		index    int
		set      stringset.StringSet
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
			t.Errorf("TestMake #%d failed: emptiness", test.index)
		}
		if test.set.Size() != test.size {
			t.Errorf("TestMake #%d failed: size", test.index)
		}
		for _, el := range test.elements {
			if !test.set.Contains(el) {
				t.Errorf("TestMake #%d failed: set does not contains \"%s\"", test.index, el)
			}
		}
	}
}
