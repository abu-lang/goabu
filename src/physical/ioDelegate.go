package physical

import (
	"steel-lang/datastructure"
)

type IOdelegate interface {
	Start(IOAdaptor, chan<- string) error
	Modified(IOAdaptor, string, datastructure.Resources) *datastructure.Resources
}

type LazyDelegate struct{}

func MakeLazyDelegate() IOdelegate {
	var res LazyDelegate
	return res
}

func (d LazyDelegate) Start(a IOAdaptor, i chan<- string) error {
	return nil
}

func (d LazyDelegate) Modified(a IOAdaptor, n string, r datastructure.Resources) *datastructure.Resources {
	return nil
}
