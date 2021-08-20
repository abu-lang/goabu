package datastructure

import "steel-lang/stringset"

type ResourceController interface {
	Start() error
	Inputs() <-chan string
	Errors() <-chan error
	Modified(string)
	IsValid() bool
	Has(string) bool
	GetTypes() map[string]string
	GetResources() Resources
	ResourceNames() stringset.StringSet
	InputsNumber() int
	String() string
	Copy() ResourceController
}
