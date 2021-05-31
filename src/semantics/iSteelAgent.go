package semantics

import "steel-lang/datastructure"

type ISteelAgent interface {
	Start() error
	Join() error
	ForAll([]datastructure.ExternalAction) error
	ReceivedActions() (<-chan chan []datastructure.ExternalAction, <-chan chan string)
	Stop() error
	IsRunning() bool
}
