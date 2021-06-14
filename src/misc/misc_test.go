package misc_test

import (
	"steel-lang/misc"
	"testing"
)

func TestMakeStringSet(t *testing.T) {
	s1 := misc.MakeStringSet("")
	s2 := misc.MakeStringSet("Lorem")
	s3 := misc.MakeStringSet("ipsum,dolor,42")
	s4 := misc.MakeStringSet(",,,")
	s5 := misc.MakeStringSet(",amet")
	s6 := misc.MakeStringSet("consectetur,consectetur")
	tests := []struct {
		index    int
		set      misc.StringSet
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
