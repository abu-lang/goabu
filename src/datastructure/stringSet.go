package datastructure

import (
	"regexp"
	"strings"
)

type StringSet map[string]bool

func NewStringSet(csv string) *StringSet {
	var res StringSet = make(map[string]bool)
	if csv == "" {
		return &res
	}
	events := strings.Split(csv, ",")
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

func (s *StringSet) AllMatch(reg string) bool {
	var idMatcher = regexp.MustCompile(reg)
	for e := range *s {
		if !idMatcher.MatchString(e) {
			return false
		}
	}
	return true
}
