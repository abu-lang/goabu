// Package memory implements GoAbU logical resources.
package memory

type ResourceController interface {
	Start() error
	Inputs() <-chan string
	Errors() <-chan error
	Modified(string)
	Extract([]string) Resources
	Enclose(Resources)
	HasDuplicates() bool
	Has(string) bool
	Types() map[string]string
	GetResources() Resources
	ResourceNames() []string
	InputsNumber() int
	String() string
	Copy() ResourceController
}
