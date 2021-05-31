package semantics

type ISteelAgent interface {
	Start() error
	Join() error
	ForAll([]ExternalAction) error
	ReceivedActions() (<-chan chan []ExternalAction, <-chan chan string)
	Stop() error
	IsRunning() bool
}
