package memory

import "steel-lang/stringset"

type ResourceController interface {
	Start() error
	Inputs() <-chan string
	Errors() <-chan error
	Modified(string)
	HasDuplicates() bool
	Has(string) bool
	GetTypes() map[string]string
	GetResources() Resources
	ResourceNames() stringset.Set
	InputsNumber() int
	String() string
	Copy() ResourceController
}
