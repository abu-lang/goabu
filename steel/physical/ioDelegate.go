package physical

import (
	"steel/memory"
)

type IOdelegate interface {
	Start(IOadaptor, chan<- string, chan<- error) error
	Modified(IOadaptor, string, memory.Resources, chan<- error) *memory.Resources
}
