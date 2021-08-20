package physical

import (
	"steel-lang/memory"
)

type IOdelegate interface {
	Start(IOAdaptor, chan<- string, chan<- error) error
	Modified(IOAdaptor, string, memory.Resources, chan<- error) *memory.Resources
}

type LazyDelegate struct{}

func MakeLazyDelegate() IOdelegate {
	var res LazyDelegate
	return res
}

func (d LazyDelegate) Start(a IOAdaptor, i chan<- string, e chan<- error) error {
	return nil
}

func (d LazyDelegate) Modified(a IOAdaptor, n string, r memory.Resources, e chan<- error) *memory.Resources {
	return nil
}
