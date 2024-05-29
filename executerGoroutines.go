// Copyright 2021 Massimo Comuzzo, Michele Pasqua and Marino Miculan
// SPDX-License-Identifier: Apache-2.0

package goabu

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/abu-lang/goabu/memory"
	"github.com/abu-lang/goabu/parser"
	"github.com/abu-lang/goabu/stringset"

	"go.uber.org/zap"
)

// TODO evaluate
const inputsRate float64 = 0.5

// milliseconds, TODO evaluate
const inputsFlush = 100

// preparedUpdates represents a list of Updates that will possibly be
// appended to the execution pool.
type preparedUpdates struct {
	confirm chan bool
	updates []Update
}

func (m *Executer) receiveInputs() {
	inputs := m.memory.Inputs()
	errors := m.memory.Errors()
	bufferSize := int(math.RoundToEven(float64(len(m.types)) * inputsRate))
	var buffer string = ""
	var l int = 0
	var timeout <-chan time.Time = nil
	var inBuffer stringset.Set = stringset.Make()
	flush := func() {
		err := m.Input(buffer)
		if err != nil {
			m.logger.Panic("Error in parsing I/O input actions: "+err.Error(),
				zap.String("act", "io_parse"), zap.String("obj", buffer))
		}
		buffer = ""
		l = 0
		inBuffer = stringset.Make()
		timeout = nil
	}
	for {
		select {
		case err := <-errors:
			m.logger.Error("I/O error: "+err.Error(), zap.String("act", "io"))
		case input := <-inputs:
			resource := strings.TrimSpace(strings.Split(input, "=")[0])
			if inBuffer.Has(resource) {
				flush()
			}
			buffer += input
			l++
			inBuffer.Insert(resource)
			if l < bufferSize {
				if l == 1 {
					timeout = time.After(inputsFlush * time.Millisecond)
				}
				continue
			}
		case <-timeout:
		}
		flush()
	}
}

func (m *Executer) receiveExternalActions() {
	requests, commandRequests := m.agent.ReceivedActions()
	for {
		actionsCh := <-requests
		if actionsCh == nil {
			return
		}
		commandsCh := <-commandRequests
		go m.serveTransaction(actionsCh, commandsCh)
	}
}

// serveTransaction interacts with the Agent in order to possibly receive and append a list of Updates to m.pool.
func (m *Executer) serveTransaction(actionsCh <-chan []byte, commandsCh chan string) {
	defer m.logger.Sync()
	wTasks, err := unmarshalWireTasks(<-actionsCh)
	if err != nil {
		m.logger.Error("Error during external actions unmarshalling: "+err.Error(),
			zap.String("act", "unmarshalling"),
			zap.String("obj", "received tasks"))
		commandsCh <- "aborted"
		return
	}
	var updates []Update
	workingSet := stringset.Make(wTasks.getRemoteResources()...)
	k := m.coordinator.requestRead(workingSet)
	m.lockMemory.RLock()
	context, workMem, err := newEmptyGruleStructures(map[string]memory.Resources{"this": m.memory.GetResources(), "ext": wTasks.Resources})
	m.lockMemory.RUnlock()
	if err != nil {
		m.logger.Panic(err.Error())
	}
	p := parser.New(m.types, workMem)
	lTasks, errs := p.ParseRemoteTasks(wTasks.Resources.Types(), wTasks.Tasks...)
	if len(errs) > 0 {
		for _, err := range errs {
			m.logger.Error("error during parsing: "+err.Error(),
				zap.String("act", "parse"),
				zap.String("obj", "received tasks"))
		}
		m.logger.Sync()
		commandsCh <- "aborted"
		return
	}
	for _, task := range lTasks {
		m.lockMemory.RLock()
		update, err := condEvalActions(task.Condition, task.Actions, context, workMem)
		if err != nil {
			m.logger.Panic("Error during received task evaluation: "+err.Error(),
				zap.String("act", "eval"),
				zap.String("obj", "received tasks"))
		}
		updates = appendNonempty(updates, update)
		m.lockMemory.RUnlock()
	}
	if len(updates) == 0 {
		if m.coordinator.confirmRead(k) {
			m.coordinator.closeRead(k)
			commandsCh <- "not_interested"
		} else {
			m.coordinator.closeRead(k)
			commandsCh <- "aborted"
		}
		return
	}
	commandsCh <- "interested"
	confirm := make(chan bool)
	switch <-commandsCh {
	case "can_commit?":
		if m.coordinator.confirmRead(k) {
			m.updateReceiver <- preparedUpdates{updates: updates, confirm: confirm}
			commandsCh <- "prepared"
		} else {
			m.coordinator.closeRead(k)
			commandsCh <- "aborted"
			return
		}
	case "do_abort":
		m.coordinator.confirmRead(k)
		m.coordinator.closeRead(k)
		commandsCh <- "done"
		return
	}
	ok := false
	switch <-commandsCh {
	case "do_commit":
		ok = true
		fallthrough
	case "do_abort":
		m.coordinator.closeRead(k)
		commandsCh <- "done"
	}
	confirm <- ok
	<-confirm
}

// startUpdateReceiver starts a goroutine responsible for appending received Updates to m.pool.
// This goroutine takes preparedUpdates over the channel returned by startUpdateReceiver and
// appends their Updates to m.pool following their arrival order but waits for a bool on their
// confirm channel before proceeding. If true is sent over the confirm channel then the related
// Updates are appended otherwise they are discarded. The received bool is then re-sent over the
// confirm channel when the operation is concluded.
func (m *Executer) startUpdateReceiver() chan<- preparedUpdates {
	res := make(chan preparedUpdates)
	go func(updates <-chan preparedUpdates) {
		var queue []preparedUpdates
		var confirm chan bool = nil
		for {
			select {
			case ok := <-confirm:
				if ok {
					m.lockPool.Lock()
					m.pool = append(m.pool, queue[0].updates...)
					m.lockPool.Unlock()
					m.logger.Info(fmt.Sprintf("Added %d updates to the pool", len(queue[0].updates)),
						zap.String("act", "add_updates"),
						zapUpdates("updates", queue[0].updates))
				}
				confirm <- ok
				queue = queue[1:]
			case u := <-updates:
				queue = append(queue, u)
			}
			if len(queue) == 0 {
				confirm = nil
			} else {
				confirm = queue[0].confirm
			}
		}
	}(res)
	return res
}
