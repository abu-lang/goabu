package stringset

type Set map[string]bool

func Make(els ...string) Set {
	var res Set = make(map[string]bool)
	for _, el := range els {
		res[el] = true
	}
	return res
}

func (set Set) Has(el string) bool {
	_, present := set[el]
	return present
}

// Precondition: set != nil
func (set Set) Insert(el string) {
	set[el] = true
}

func (set Set) Remove(el string) {
	delete(set, el)
}

func (set Set) Empty() bool {
	return len(set) == 0
}

func (set Set) Size() int {
	return len(set)
}

func (fst Set) IntersectsWith(snd Set) bool {
	res := false
	if snd != nil {
		for el := range fst {
			if snd.Has(el) {
				res = true
				break
			}
		}
	}
	return res
}

func (fst Set) Contains(snd Set) bool {
	for el := range snd {
		_, present := fst[el]
		if !present {
			return false
		}
	}
	return true
}

// Precondition: dst != nil
func (dst Set) Add(src Set) {
	for el := range src {
		dst[el] = true
	}
}

func (dst Set) Intersect(src Set) {
	for el := range dst {
		if !src.Has(el) {
			delete(dst, el)
		}
	}
}

func (set Set) Slice() []string {
	res := make([]string, 0, len(set))
	for el := range set {
		res = append(res, el)
	}
	return res
}

func (set Set) Clone() Set {
	res := Make()
	for el := range set {
		res[el] = true
	}
	return res
}
