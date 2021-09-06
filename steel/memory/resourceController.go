package memory

type ResourceController interface {
	Start() error
	Inputs() <-chan string
	Errors() <-chan error
	Modified(string)
	HasDuplicates() bool
	Has(string) bool
	Types() map[string]string
	GetResources() Resources
	ResourceNames() []string
	InputsNumber() int
	String() string
	Copy() ResourceController
}
