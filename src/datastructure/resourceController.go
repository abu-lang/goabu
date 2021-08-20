package datastructure

import "steel-lang/misc"

type ResourceController interface {
	Start() error
	Inputs() <-chan string
	Errors() <-chan error
	Modified(string)
	IsValid() bool
	Has(string) bool
	GetTypes() map[string]string
	GetResources() Resources
	ResourceNames() misc.StringSet
	InputsNumber() int
	String() string
	Clone() ResourceController
}
