package semantics

type ISteelAgent interface {
	Start() error
	Join() error
	ForAll([]ExternalAction) error
	ReceivedActions() <-chan []ExternalAction
	Stop() error
	IsRunning() bool
}
