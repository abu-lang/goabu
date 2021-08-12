package semantics

import (
	"errors"
)

type MockAgent struct {
	running           bool
	operations        chan chan []byte
	operationCommands chan chan string
}

func MakeMockAgent() ISteelAgent {
	return &MockAgent{
		running:           false,
		operations:        make(chan chan []byte),
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

func (a *MockAgent) ForAll(actions []byte) error {
	if !a.running {
		return errors.New("agent is not running")
	}
	if len(actions) == 0 {
		return nil
	}
	actionsCh := make(chan []byte)
	commandsCh := make(chan string)
	a.operations <- actionsCh
	a.operationCommands <- commandsCh
	actionsCh <- actions
	if <-commandsCh == "interested" {
		commandsCh <- "can_commit?"
		if <-commandsCh == "prepared" {
			commandsCh <- "do_commit"
			<-commandsCh
		}
	}
	return nil
}

func (a *MockAgent) ReceivedActions() (<-chan chan []byte, <-chan chan string) {
	return a.operations, a.operationCommands
}

func (a *MockAgent) Stop() error {
	if !a.running {
		return errors.New("agent is not running")
	}
	a.operations <- nil
	a.running = false
	return nil
}

func (a *MockAgent) SetLogLevel(l int) {}
