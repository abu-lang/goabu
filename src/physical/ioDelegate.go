package physical

import (
	"steel-lang/datastructure"
)

type IOdelegate interface {
	Start(IOAdaptor, chan<- string, chan<- error) error
	Modified(IOAdaptor, string, datastructure.Resources, chan<- error) *datastructure.Resources
}

type LazyDelegate struct{}

func MakeLazyDelegate() IOdelegate {
	var res LazyDelegate
	return res
}

func (d LazyDelegate) Start(a IOAdaptor, i chan<- string, e chan<- error) error {
	return nil
}

func (d LazyDelegate) Modified(a IOAdaptor, n string, r datastructure.Resources, e chan<- error) *datastructure.Resources {
	return nil
}
