package semantics

import (
	"errors"
	"steel-lang/datastructure"
)

type MockAgent struct {
	running           bool
	operations        chan chan []datastructure.ExternalAction
	operationCommands chan chan string
}

func MakeMockAgent() ISteelAgent {
	return &MockAgent{
		running:           false,
		operations:        make(chan chan []datastructure.ExternalAction),
		operationCommands: make(chan chan string),
	}
}

func (a *MockAgent) IsRunning() bool {
	return a.running
}

func (a *MockAgent) Start() error {
	if a.running {
		return errors.New("agent is already running")
	}
	a.running = true
	return nil
}

func (a *MockAgent) Join() error {
	if !a.running {
		return errors.New("agent is not running")
	}
	return nil
}

func (a *MockAgent) ForAll(actions []datastructure.ExternalAction) error {
	if !a.running {
		return errors.New("agent is not running")
	}
	if len(actions) == 0 {
		return nil
	}
	actionsCh := make(chan []datastructure.ExternalAction)
	commandsCh := make(chan string)
	a.operations <- actionsCh
	a.operationCommands <- commandsCh
	actionsCh <- actions
	if <-commandsCh == "interested" {
		commandsCh <- "do_commit"
		<-commandsCh
	}
	return nil
}

func (a *MockAgent) ReceivedActions() (<-chan chan []datastructure.ExternalAction, <-chan chan string) {
	return a.operations, a.operationCommands
}

func (a *MockAgent) Stop() error {
	if !a.running {
		return errors.New("agent is not running")
	}
	a.running = false
	return nil
}
