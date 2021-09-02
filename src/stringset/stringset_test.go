package stringset_test

import (
	"steel-lang/stringset"
	"testing"
)

func TestMake(t *testing.T) {
	s1 := stringset.Make()
	s2 := stringset.Make("Lorem")
	s3 := stringset.Make("ipsum", "dolor", "42")
	s4 := stringset.Make("", "", "", "")
	s5 := stringset.Make("", "amet")
	s6 := stringset.Make("consectetur", "consectetur")
	tests := []struct {
		index    int
		set      stringset.Set
		empty    bool
		elements []string
	}{
		//  {_, set, empty, elements},
		{1, s1, true, []string{}},
		{2, s2, false, []string{"Lorem"}},
		{3, s3, false, []string{"ipsum", "dolor", "42"}},
		{4, s4, false, []string{""}},
		{5, s5, false, []string{"", "amet"}},
		{6, s6, false, []string{"consectetur"}},
	}
	for _, test := range tests {
		if test.set.Empty() != test.empty {
			t.Errorf("TestMake#%d failed: emptiness", test.index)
		}
		if test.set.Size() != len(test.elements) {
			t.Errorf("TestMake#%d failed: size", test.index)
		}
		for _, el := range test.elements {
			if !test.set.Has(el) {
				t.Errorf("TestMake#%d failed: set does not contains \"%s\"", test.index, el)
			}
		}
	}
}
