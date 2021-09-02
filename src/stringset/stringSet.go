package stringset

import "strings"

type StringSet map[string]bool

func Make(csv string) StringSet {
	var res StringSet = make(map[string]bool)
	if csv == "" {
		return res
	}
	elements := strings.Split(csv, ",")
	for _, el := range elements {
		res[el] = true
	}
	return res
}

func (set StringSet) Contains(el string) bool {
	_, present := set[el]
	return present
}

// Precondition: set != nil
func (set StringSet) Insert(el string) {
	set[el] = true
}

func (set StringSet) Remove(el string) {
	delete(set, el)
}

func (set StringSet) Empty() bool {
	return len(set) == 0
}

func (set StringSet) Size() int {
	return len(set)
}

func (fst StringSet) IntersectsWith(snd StringSet) bool {
	res := false
	if snd != nil {
		for el := range fst {
			if snd.Contains(el) {
				res = true
				break
			}
		}
	}
	return res
}

func (fst StringSet) ContainsSet(snd StringSet) bool {
	for el := range snd {
		_, present := fst[el]
		if !present {
			return false
		}
	}
	return true
}

// Precondition: dst != nil
func (dst StringSet) Add(src StringSet) {
	for el := range src {
		dst[el] = true
	}
}

func (dst StringSet) Intersect(src StringSet) {
	for el := range dst {
		if !src.Contains(el) {
			delete(dst, el)
		}
	}
}

func (set StringSet) Clone() StringSet {
	res := Make("")
	for el := range set {
		res[el] = true
	}
	return res
}
