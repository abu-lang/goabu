package datastructure

import "steel-lang/misc"

type ResourceController interface {
	Start() error
	Inputs() <-chan string
	Modified(string) error
	IsValid() bool
	Has(string) bool
	GetTypes() map[string]string
	GetResources() Resources
	ResourceNames() misc.StringSet
	InputsNumber() int
	String() string
	Clone() ResourceController
}
