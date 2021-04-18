package datastructure

import (
	"strings"
)

type StringSet map[string]bool

func NewStringSet(csv string) *StringSet {
	events := strings.Split(csv, ",")
	var res StringSet = make(map[string]bool)
	for _, e := range events {
		res[e] = true
	}
	return &res
}

func (s *StringSet) Contains(e string) bool {
	_, present := (*s)[e]
	return present
}

func (s *StringSet) Insert(e string) {
	(*s)[e] = true
}

func (s *StringSet) Empty() bool {
	return len(*s) == 0
}

func (fst *StringSet) Intersects(snd *StringSet) bool {
	res := false
	if snd != nil {
		for e := range *fst {
			if snd.Contains(e) {
				res = true
				break
			}
		}
	}
	return res
}

func (fst *StringSet) Add(snd *StringSet) {
	if snd != nil {
		for e := range *snd {
			(*fst)[e] = true
		}
	}
}
