package semantics

import (
	"math"
	"steel-lang/stringset"
	"strings"
	"time"

	"go.uber.org/zap"
)

const inputsRate float64 = 5.0

// milliseconds
const inputsFlush = 100

func (m *MuSteelExecuter) receiveInputs() {
	inputs := m.memory.Inputs()
	errors := m.memory.Errors()
	queueSize := int(math.RoundToEven(float64(m.memory.InputsNumber()) * inputsRate))
	var queue string = ""
	var l int = 0
	var timeout <-chan time.Time = nil
	var queued stringset.StringSet = stringset.Make("")
	for {
		select {
		case err := <-errors:
			m.logger.Error("I/O error: "+err.Error(), zap.String("act", "io"))
		case action := <-inputs:
			resource := strings.TrimSpace(strings.Split(action, "=")[0])
			if queued.Contains(resource) {
				err := m.Input(queue)
				if err != nil {
					m.logger.Panic("Error in parsing I/O input actions: "+err.Error(),
						zap.String("act", "io_parse"), zap.String("obj", queue))
				}
				queue = ""
				l = 0
				timeout = nil
				queued = stringset.Make("")
			}
			queue += action
			l++
			queued.Insert(resource)
			if l == 1 {
				timeout = time.After(inputsFlush * time.Millisecond)
			}
			if l < queueSize {
				continue
			}
		case <-timeout:
		}
		err := m.Input(queue)
		if err != nil {
			m.logger.Panic("Error in parsing I/O input actions: "+err.Error(),
				zap.String("act", "io_parse"), zap.String("obj", queue))
		}
		queue = ""
		l = 0
		timeout = nil
		queued = stringset.Make("")
	}
}

func (m *MuSteelExecuter) receiveExternalActions() {
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

func (m *MuSteelExecuter) serveTransaction(actionsCh <-chan []byte, commandsCh chan string) {
	defer m.logger.Sync()
	eActions, err := unmarshalExternalActions(<-actionsCh, m.types)
	if err != nil {
		m.logger.Error("Error during external actions unmarshalling: "+err.Error(),
			zap.String("act", "unmarshalling"),
			zap.String("obj", "external actions"))
		commandsCh <- "aborted"
		return
	}
	var updates []Update
	localResources := stringset.Make("")
	for r := range m.types {
		localResources.Insert(r)
	}
	workingSet := stringset.Make("")
	for _, eAction := range eActions {
		if localResources.ContainsSet(eAction.CondWorkingSet) {
			workingSet.Add(eAction.CondWorkingSet)
			for _, ws := range eAction.WorkingSets {
				if localResources.ContainsSet(ws) {
					workingSet.Add(ws)
				}
			}
		}
	}
	k := m.coordinator.requestRead(workingSet)
	m.lockMemory.RLock()
	context, workMem, err := m.newEmptyGruleStructures("ext")
	m.lockMemory.RUnlock()
	if err != nil {
		m.logger.Panic(err.Error())
	}
	for _, eAction := range eActions {
		if localResources.ContainsSet(eAction.CondWorkingSet) {
			actions := eAction.cullActions(localResources)
			if len(actions) == 0 {
				continue
			}
			m.lockMemory.RLock()
			update, err := condEvalActions(eAction.Condition, actions, context, workMem)
			if err != nil {
				m.logger.Panic("Error during external actions evaluation: "+err.Error(),
					zap.String("act", "eval"),
					zap.String("obj", "external actions"))
			}
			updates = appendNonempty(updates, update)
			m.lockMemory.RUnlock()
		}
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
	switch <-commandsCh {
	case "can_commit?":
		if m.coordinator.confirmRead(k) {
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
	switch <-commandsCh {
	case "do_commit":
		m.lockPool.Lock()
		m.pool = append(m.pool, updates...)
		m.lockPool.Unlock()
		m.logger.Info("Added external actions", zap.String("act", "add_updates"), zap.Array("updates", poolLogger(updates)))
		fallthrough
	case "do_abort":
		m.coordinator.closeRead(k)
		commandsCh <- "done"
	}
}
