package semantics

import (
	"math"
	"steel-lang/misc"
	"strings"
	"time"
)

const inputsRate float64 = 5.0

// milliseconds
const inputsFlush = 100

func (m *MuSteelExecuter) receiveInputs() {
	inputs := m.memory.Inputs()
	queueSize := int(math.RoundToEven(float64(m.memory.InputsNumber()) * inputsRate))
	var queue string = ""
	var l int = 0
	var timeout <-chan time.Time = nil
	var queued misc.StringSet = misc.MakeStringSet("")
	for {
		select {
		case action := <-inputs:
			resource := strings.TrimSpace(strings.Split(action, "=")[0])
			if queued.Contains(resource) {
				err := m.Input(queue)
				if err != nil {
					panic(err)
				}
				queue = ""
				l = 0
				timeout = nil
				queued = misc.MakeStringSet("")
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
			panic(err)
		}
		queue = ""
		l = 0
		timeout = nil
		queued = misc.MakeStringSet("")
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
	eActions, err := unmarshalExternalActions(<-actionsCh, m.types)
	if err != nil {
		panic(err)
	}
	var sActions [][]SemanticAction
	localResources := misc.MakeStringSet("")
	for r := range m.types {
		localResources.Insert(r)
	}
	workingSet := misc.MakeStringSet("")
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
			sActions = appendNonempty(sActions, condEvalActions(eAction.Condition, actions, context, workMem))
			m.lockMemory.RUnlock()
		}
	}
	if len(sActions) == 0 {
		if m.coordinator.confirmRead(k) {
			commandsCh <- "not_interested"
		} else {
			commandsCh <- "aborted"
		}
		m.coordinator.closeRead(k)
		return
	}
	commandsCh <- "interested"
	switch <-commandsCh {
	case "can_commit?":
		if m.coordinator.confirmRead(k) {
			commandsCh <- "prepared"
		} else {
			commandsCh <- "aborted"
			m.coordinator.closeRead(k)
			return
		}
	case "do_abort":
		commandsCh <- "done"
		m.coordinator.confirmRead(k)
		m.coordinator.closeRead(k)
		return
	}
	switch <-commandsCh {
	case "do_commit":
		m.lockPool.Lock()
		m.pool = append(m.pool, sActions...)
		m.lockPool.Unlock()
		fallthrough
	case "do_abort":
		commandsCh <- "done"
		m.coordinator.closeRead(k)
	}
}
