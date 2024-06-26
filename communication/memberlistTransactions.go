// Copyright 2021 Massimo Comuzzo, Michele Pasqua and Marino Miculan
// SPDX-License-Identifier: Apache-2.0

package communication

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"k8s.io/apimachinery/pkg/util/sets"

	"github.com/hashicorp/memberlist"
	"go.uber.org/zap"
)

// TODO evaluate
const (
	// milliseconds
	timeoutPhaseResend = 6000
	timeoutWakeMonitor = 12000
)

// agentID returns the agent id of node.
func agentID(node *memberlist.Node) string {
	return string(node.Meta)
}

type transactionInfo struct {
	Initiator    string
	Number       int
	Payload      []byte
	Participants []string

	// initiatorID possibly points to the agent id of the initiator or is nil.
	initiatorID *string
	stopMonitor chan bool
	coordinated bool
	commands    chan string
}

func (t *transactionInfo) id() string {
	return fmt.Sprintf("%s->%d", t.Initiator, t.Number)
}

func (t *transactionInfo) buryParticipants(members []*memberlist.Node) {
	alives := sets.New[string]()
	for _, member := range members {
		alives.Insert(member.Name)
	}
	buried := 0
	for i, participant := range t.Participants {
		if !alives.Has(participant) {
			t.Participants[i] = ""
			buried++
		}
	}
	remaining := len(t.Participants) - buried
	j := 1
	for i := 0; i < remaining; i++ {
		if t.Participants[i] == "" {
			for t.Participants[j] == "" {
				j++
			}
			t.Participants[i], t.Participants[j] = t.Participants[j], t.Participants[i]
		}
		j++
	}
	t.Participants = t.Participants[:remaining]
}

type transactionChannels struct {
	Initiator       string
	Number          int
	areInterested   chan string
	areUninterested chan string
	arePrepared     chan string
	haveAborted     chan string
	haveCommitted   chan string
}

func makeTransactionChannels(t transactionInfo) transactionChannels {
	return transactionChannels{
		Initiator:       t.Initiator,
		Number:          t.Number,
		areInterested:   make(chan string, msgBuffLen),
		areUninterested: make(chan string, msgBuffLen),
		arePrepared:     make(chan string, msgBuffLen),
		haveAborted:     make(chan string, msgBuffLen),
		haveCommitted:   make(chan string, msgBuffLen),
	}
}

func (t transactionChannels) id() string {
	return fmt.Sprintf("%s->%d", t.Initiator, t.Number)
}

func (t transactionChannels) line(response string) chan<- string {
	switch response {
	case "interested":
		return t.areInterested
	case "not_interested":
		return t.areUninterested
	case "prepared":
		return t.arePrepared
	case "aborted":
		return t.haveAborted
	case "committed":
		return t.haveCommitted
	}
	return nil
}

func (a *MemberlistAgent) interested(tran transactionInfo) ([]string, error) {
	m := message{
		Type:        "interested?",
		Sender:      a.list.LocalNode(),
		Transaction: tran,
	}
	msg, ok := m.marshal("interested?", a.logger)
	if !ok {
		a.logger.Panic("Could not marshal interested? message",
			zap.String("act", "marshalling"),
			zap.String("obj", "interested?"))
	}
	channels := makeTransactionChannels(tran)
	channelsCh := make(chan transactionChannels)
	a.coordinatedChannels <- channelsCh
	channelsCh <- channels
	nodes, err := a.interestPhase(msg, channels)
	a.testsHaltIf(TestsAfterInterested)
	if len(nodes) == 0 {
		channelsCh := make(chan transactionChannels)
		a.coordinatedChannels <- channelsCh
		channelsCh <- transactionChannels{
			Initiator: tran.Initiator,
			Number:    tran.Number,
		}
	}
	return nodes, err
}

// interestPhase sends msg to all nodes selected by the MemberlistDelegate filterPartecipants method.
// It returns a slice containing the names of the nodes that responded with "interested" if no node
// responded with "aborted" otherwise it aborts the transaction and returns an error.
//
// Testing: If a.test == TestsMidInterested it simulates a crash failure of the agent after having
// received TestsMidSends responses.
func (a *MemberlistAgent) interestPhase(msg []byte, channels transactionChannels) ([]string, error) {
	aborted := ""
	waitFor := sets.New[string]()
	for _, member := range a.adapter.filterParticipants(a.list.Members()) {
		waitFor.Insert(member.Name)
	}
	var interested []string
	for waitFor.Len() > 0 {
		var timeout <-chan time.Time = nil
		waitForCopy := waitFor.Clone()
		receiversCh := make(chan sets.Set[string])
		if a.test == TestsMidInterested {
			go a.testsPhaseSend(waitForCopy, msg, true, receiversCh, TestsMidSends)
		} else {
			go a.phaseSend(waitForCopy, msg, true, channels.id(), receiversCh)
		}
		received := 0
	INTERESTED:
		for waitFor.Len() > 0 {
			select {
			case receivers := <-receiversCh:
				timeout = time.After(time.Millisecond * timeoutPhaseResend)
				waitFor = waitFor.Intersection(receivers)
			case participant := <-channels.areInterested:
				received++
				interested = append(interested, participant)
				delete(waitFor, participant)
			case uninterested := <-channels.areUninterested:
				received++
				delete(waitFor, uninterested)
			case aborted = <-channels.haveAborted:
				received++
				delete(waitFor, aborted)
			case <-timeout:
				break INTERESTED
			}
			if a.test == TestsMidInterested && received >= TestsMidSends {
				a.testsHaltAndBlock()
			}
		}
	}
	if aborted == "" {
		return interested, nil
	}
	order := message{
		Type:   "do_abort",
		Sender: a.list.LocalNode(),
		Transaction: transactionInfo{
			Initiator: channels.Initiator,
			Number:    channels.Number,
		},
	}
	abrt, ok := order.marshal(order.Type, a.logger)
	if !ok {
		a.logger.Panic("Could not marshal "+order.Type+" message", zap.String("act", "marshalling"), zap.String("obj", order.Type))
	}
	dests := sets.New[string]()
	for _, i := range interested {
		dests.Insert(i)
	}
	a.secondPhase(dests, abrt, channels.haveAborted, channels.id())
	return nil, fmt.Errorf("%s has aborted", aborted)
}

func (a *MemberlistAgent) coordinateTransaction(tran transactionInfo) error {
	canCommit := message{
		Type:        "can_commit?",
		Sender:      a.list.LocalNode(),
		Transaction: tran,
	}
	msg, ok := canCommit.marshal("can_commit?", a.logger)
	if !ok {
		return errors.New("could not marshal can_commit? message")
	}
	receivers := sets.New[string]()
	for _, nodeName := range tran.Participants {
		receivers.Insert(nodeName)
	}
	channels := makeTransactionChannels(tran)
	channelsCh := make(chan transactionChannels)
	a.coordinatedChannels <- channelsCh
	channelsCh <- channels
	a.logger.Debug("Started transaction",
		zap.String("subj", a.id),
		zap.String("act", "start_tran"),
		zap.Int("participants", receivers.Len()))
	res := a.firstPhase(receivers, msg, channels)
	a.logger.Debug("Terminated first phase",
		zap.String("subj", a.id),
		zap.String("act", "end_1_phase"),
		zap.Int("participants", receivers.Len()))
	a.testsHaltIf(TestsAfterFirst)
	responses := channels.haveCommitted
	action := "do_commit"
	if res != nil {
		responses = channels.haveAborted
		action = "do_abort"
	}
	order := message{
		Type:   action,
		Sender: a.list.LocalNode(),
		Transaction: transactionInfo{
			Initiator: tran.Initiator,
			Number:    tran.Number,
		},
	}
	msg, ok = order.marshal(order.Type, a.logger)
	if !ok {
		a.logger.Panic("Could not marshal "+order.Type+" message", zap.String("act", "marshalling"), zap.String("obj", order.Type))
	}
	a.secondPhase(receivers, msg, responses, tran.id())
	channelsCh = make(chan transactionChannels)
	a.coordinatedChannels <- channelsCh
	channelsCh <- transactionChannels{
		Initiator: tran.Initiator,
		Number:    tran.Number,
	}
	a.logger.Debug("Terminated transaction", zap.String("subj", a.id), zap.String("act", "end_tran"))
	return res
}

func (a *MemberlistAgent) firstPhase(participants sets.Set[string], msg []byte, channels transactionChannels) error {
	waitFor := participants.Clone()
	for waitFor.Len() > 0 {
		var timeout <-chan time.Time = nil
		waitForCopy := waitFor.Clone()
		receiversCh := make(chan sets.Set[string])
		if a.test == TestsMidFirst {
			go a.testsPhaseSend(waitForCopy, msg, true, receiversCh, TestsMidSends)
		} else {
			go a.phaseSend(waitForCopy, msg, true, channels.id(), receiversCh)
		}
		received := 0
	GET_RESPONSES_1:
		for waitFor.Len() > 0 {
			select {
			case receivers := <-receiversCh:
				timeout = time.After(time.Millisecond * timeoutPhaseResend)
				waitFor = waitFor.Intersection(receivers)
			case prepared := <-channels.arePrepared:
				received++
				delete(waitFor, prepared)
			case <-channels.haveCommitted: // I am substituting initiator
				received++
				if a.test == TestsMidFirst {
					break
				}
				return nil
			case aborted := <-channels.haveAborted:
				received++
				if a.test == TestsMidFirst {
					break
				}
				return fmt.Errorf("%s has aborted", aborted)
			case <-timeout:
				break GET_RESPONSES_1
			}
			if a.test == TestsMidFirst && received >= TestsMidSends {
				a.testsHaltAndBlock()
			}
		}
	}
	return nil
}

// secondPhase sends msg to waitFor with best effort and awaits the responses.
// After having performed the sends if after timeoutPhaseResend milliseconds some node has not
// responded then msg is resended to those nodes and the timeout is restarted.
//
// responses is a channel that must pass the name of a node when a response from that node is received.
func (a *MemberlistAgent) secondPhase(waitFor sets.Set[string], msg []byte, responses <-chan string, tranID string) {
	for waitFor.Len() > 0 {
		var timeout <-chan time.Time = nil
		waitForCopy := waitFor.Clone()
		receiversCh := make(chan sets.Set[string])
		if a.test == TestsMidSecond {
			go a.testsPhaseSend(waitForCopy, msg, false, receiversCh, TestsMidSends)
		} else {
			go a.phaseSend(waitForCopy, msg, false, tranID, receiversCh)
		}
		received := 0
	GET_RESPONSES_2:
		for waitFor.Len() > 0 {
			select {
			case receivers := <-receiversCh:
				timeout = time.After(time.Millisecond * timeoutPhaseResend)
				waitFor = waitFor.Intersection(receivers)
			case responded := <-responses:
				received++
				waitFor.Delete(responded)
			case <-timeout:
				break GET_RESPONSES_2
			}
			if a.test == TestsMidSecond && received >= TestsMidSends {
				a.testsHaltAndBlock()
			}
		}
	}
}

// phaseSend sends msg to all alive nodes in receivers.
//
// reliableSend indicates wheter to use a reliable transport protocol or not.
//
// done will pass the nodes that were alive during the execution of phaseSend.
//
// Testing: if a.test == TestsUnreliable about 10% of the sends aren't performed.
func (a *MemberlistAgent) phaseSend(receivers sets.Set[string], msg []byte, reliableSend bool, tranID string, done chan<- sets.Set[string]) {
	newReceivers := sets.New[string]()
	for _, member := range a.list.Members() {
		if receivers.Has(member.Name) {
			newReceivers.Insert(member.Name)
			if a.test == TestsUnreliableSend && rand.Float32() < 0.1 {
				continue
			}
			if reliableSend {
				a.list.SendReliable(member, msg)
			} else {
				a.list.SendBestEffort(member, msg)
			}
			a.logger.Debug(fmt.Sprintf("Sent message to \"%s\"", agentID(member)),
				zap.String("subj", a.id),
				zap.String("tran", tranID),
				zap.String("act", "send"),
				zap.Int("size", len(msg)),
				zap.String("to", agentID(member)))
		}
	}
	done <- newReceivers
}

func (a *MemberlistAgent) testsPhaseSend(receivers sets.Set[string], msg []byte, reliableSend bool, done chan<- sets.Set[string], haltAfter int) {
	selected := make([]*memberlist.Node, 0, haltAfter)
	sent := 0
	for _, member := range a.list.Members() {
		if sent == haltAfter {
			break
		}
		if receivers.Has(member.Name) {
			selected = append(selected, member)
			if reliableSend {
				a.list.SendReliable(member, msg)
			} else {
				a.list.SendBestEffort(member, msg)
			}
			sent++
		}
	}
	for {
		time.Sleep(time.Millisecond * timeoutPhaseResend)
		for _, member := range selected {
			if reliableSend {
				a.list.SendReliable(member, msg)
			} else {
				a.list.SendBestEffort(member, msg)
			}
		}
	}
}

func demuxResponses(coordinated <-chan chan transactionChannels, responses <-chan message, quit <-chan chan bool, logger *zap.Logger) {
	stopping := false
	lines := make(map[string]transactionChannels)
	for {
		select {
		case tranCh := <-coordinated:
			tran := <-tranCh
			if tran.arePrepared != nil {
				lines[tran.id()] = tran
			} else {
				delete(lines, tran.id())
				if stopping && len(lines) == 0 {
					return
				}
			}
		case c := <-quit:
			stopping = true
			defer func() { c <- true }()
			defer logger.Sync()
			if len(lines) == 0 {
				return
			}
		case response := <-responses:
			channels, present := lines[response.Transaction.id()]
			if !present {
				logger.Debug("Transaction already terminated: discarding response",
					zap.String("act", "discard"),
					zap.String("obj", "transaction response"),
					zap.String("from", agentID(response.Sender)))
				break
			}
			c := channels.line(response.Type)
			select {
			case c <- response.Sender.Name:
				logger.Debug("Received: "+response.Type,
					zap.String("act", "recv"),
					zap.String("obj", response.Type),
					zap.String("from", agentID(response.Sender)))
			default:
				logger.Warn("Discarded interested status",
					zap.String("act", "discard"),
					zap.String("obj", response.Type),
					zap.String("from", agentID(response.Sender)))
			}
		}
	}
}

func (a *MemberlistAgent) handleTransactions() {
	stopping := false
	for {
		select {
		case c := <-a.quitTransactions:
			stopping = true
			defer func() { c <- true }()
			if len(a.transactions) == 0 {
				return
			}
		case msg := <-a.transactionMessages:
			response := message{
				Sender: a.list.LocalNode(),
				Transaction: transactionInfo{
					Initiator: msg.Transaction.Initiator,
					Number:    msg.Transaction.Number,
				},
			}
			respond := true
			id := msg.Transaction.id()
			status := a.getStatus(id)
			switch msg.Type {

			case "__interested__", "__not_interested__", "__aborted__":
				if status != "evaluating" {
					a.logger.Panic("Received evaluation result when not evaluating", zap.String("act", "recv"), zap.String("obj", msg.Type))
				}
				tran := a.transactions[id]
				response.Type = strings.Trim(msg.Type, "_")
				if response.Type == "interested" {
					tran.stopMonitor = make(chan bool)
					go a.monitorTransaction(*tran)
				} else {
					a.terminated[id] = response.Type
				}
				respond = false
				for _, node := range a.list.Members() {
					if node.Name == tran.Initiator {
						msg.Sender = node
						respond = true
						break
					}
				}

			case "interested?":
				switch status {
				case "new":
					if stopping {
						respond = false
						break
					}
					actionsCh := make(chan []byte)
					commandsCh := make(chan string)
					a.operations <- actionsCh
					a.operationCommands <- commandsCh
					actionsCh <- msg.Transaction.Payload
					tran := &transactionInfo{}
					*tran = msg.Transaction
					tran.Payload = nil
					tran.Participants = nil
					go func(tran transactionInfo) {
						a.transactionMessages <- message{
							Sender:      a.list.LocalNode(),
							Type:        "__" + <-commandsCh + "__",
							Transaction: tran,
						}
					}(*tran)
					tran.commands = commandsCh
					for _, member := range a.list.Members() {
						if member.Name == tran.Initiator {
							agtID := agentID(member)
							tran.initiatorID = &agtID
							break
						}
					}
					a.transactions[id] = tran
					respond = false
				case "interested", "not_interested":
					response.Type = status
				default:
					respond = false
				}

			case "can_commit?":
				a.logger.Debug("Received: can_commit?",
					zap.String("subj", a.id),
					zap.String("act", "recv"),
					zap.String("obj", "can_commit?"),
					zap.String("from", agentID(msg.Sender)))
				switch status {
				case "not_interested", "evaluating":
					a.logger.Panic("Received can_commit? but I am evaluating or I wasn't interested",
						zap.String("act", "recv"),
						zap.String("obj", "can_commit?"),
						zap.String("from", agentID(msg.Sender)))
				case "interested":
					if a.test == TestsAbort {
						a.abort(id)
						response.Type = "aborted"
						break
					}
					tran := a.transactions[id]
					tran.commands <- "can_commit?"
					response.Type = <-tran.commands
					if response.Type == "aborted" {
						tran.stopMonitor <- true
						a.terminated[id] = "aborted"
						delete(a.transactions, id)
					} else {
						tran.Participants = msg.Transaction.Participants
					}
				default:
					response.Type = status
				}

			case "do_abort":
				a.logger.Debug("Received: do_abort",
					zap.String("subj", a.id),
					zap.String("act", "recv"),
					zap.String("obj", "do_abort"),
					zap.String("from", agentID(msg.Sender)))
				switch status {
				case "evaluating":
					select {
					case a.transactionMessages <- msg:
					default:
					}
					respond = false
				case "committed", "new", "not_interested":
					a.logger.Panic("Received do_abort for a committed, new or uninteresting transaction",
						zap.String("act", "recv"),
						zap.String("obj", "do_abort"),
						zap.String("from", agentID(msg.Sender)))
				case "prepared", "interested":
					a.abort(id)
				}
				response.Type = "aborted"

			case "do_commit":
				a.logger.Debug("Received: do_commit",
					zap.String("subj", a.id),
					zap.String("act", "recv"),
					zap.String("obj", "do_commit"),
					zap.String("from", agentID(msg.Sender)))
				switch status {
				case "prepared":
					a.commit(id)
				case "committed":
				default:
					a.logger.Panic("Spurious do_commit",
						zap.String("act", "recv"),
						zap.String("obj", "do_commit"),
						zap.String("from", agentID(msg.Sender)))
				}
				response.Type = "committed"

			case "get_decision":
				switch status {
				case "new", "evaluating", "not_interested":
					a.logger.Panic("Received get_decision for a new, evaluating or uninteresting transaction",
						zap.String("act", "recv"),
						zap.String("obj", "get_decision"),
						zap.String("from", agentID(msg.Sender)))
				case "interested", "prepared":
					respond = false
					tran := a.transactions[id]
					initiatorAlive := false
					for _, member := range a.list.Members() {
						if member.Name == tran.Initiator {
							initiatorAlive = true
							break
						}
					}
					if initiatorAlive {
						break
					}
					if status == "interested" {
						a.logger.Debug("Aborting: I was interested but initiator is dead",
							zap.String("subj", a.id),
							zap.String("act", "abort_tran"),
							zap.Stringp("initiator", tran.initiatorID))
						a.abort(id)
						break
					}
					if tran.coordinated {
						msg.Transaction.Payload = nil
						msg.Transaction.Participants = nil
						msg.Type = "prepared"
						select {
						case a.transactionResponses <- msg:
						default:
						}
						break
					}
					tran.buryParticipants(a.list.Members())
					if len(tran.Participants) == 0 {
						a.logger.Panic("Empty participant list")
					}
					head := tran.Participants[0]
					if head == a.list.LocalNode().Name {
						a.logger.Debug("Initiator is dead: becoming new coordinator",
							zap.String("subj", a.id),
							zap.String("act", "coord"),
							zap.Stringp("initiator", tran.initiatorID))
						go a.coordinateTransaction(*tran)
						tran.coordinated = true
						break
					}
					msg.Sender = a.list.LocalNode()
					deflected, ok := msg.marshal("deflected message", a.logger)
					if ok {
						for _, member := range a.list.Members() {
							if member.Name == head {
								a.list.SendReliable(member, deflected)
								a.logger.Debug(fmt.Sprintf("Sent message to \"%s\"", agentID(member)),
									zap.String("subj", a.id),
									zap.String("tran", id),
									zap.String("act", "send"),
									zap.Int("size", len(deflected)),
									zap.String("to", agentID(member)))
								break
							}
						}
					}
				default:
					response.Type = "do_abort"
					if status == "committed" {
						response.Type = "do_commit"
					}
					a.logger.Debug("Received get_decision, responding: "+response.Type,
						zap.String("subj", a.id),
						zap.String("act", "send"),
						zap.String("obj", response.Type),
						zap.String("to", agentID(msg.Sender)))
				}

			default:
				respond = false
				a.logger.DPanic("Unsupported transaction message: "+msg.Type,
					zap.String("act", "recv"),
					zap.String("obj", `"`+msg.Type+`"`),
					zap.String("from", agentID(msg.Sender)))
			}

			if respond {
				responseMsg, ok := response.marshal("response", a.logger)
				if ok {
					a.list.SendBestEffort(msg.Sender, responseMsg)
					a.logger.Debug(fmt.Sprintf("Sent message to \"%s\"", agentID(msg.Sender)),
						zap.String("subj", a.id),
						zap.String("tran", id),
						zap.String("act", "send"),
						zap.Int("size", len(responseMsg)),
						zap.String("to", agentID(msg.Sender)))
				}
			}
			a.logger.Sync()
			if stopping && len(a.transactions) == 0 {
				return
			}
		}
	}
}

func (a *MemberlistAgent) getStatus(id string) string {
	outcome, present := a.terminated[id]
	if present {
		return outcome
	}
	tran, present := a.transactions[id]
	if !present {
		return "new"
	}
	if tran.stopMonitor == nil {
		return "evaluating"
	}
	if len(tran.Participants) == 0 {
		return "interested"
	}
	return "prepared"
}

func (a *MemberlistAgent) abort(id string) {
	tran, present := a.transactions[id]
	if !present {
		a.logger.Panic("Called abort for non-existent transaction")
	}
	tran.stopMonitor <- true
	tran.commands <- "do_abort"
	<-tran.commands
	a.terminated[id] = "aborted"
	delete(a.transactions, id)
}

func (a *MemberlistAgent) commit(id string) {
	tran, present := a.transactions[id]
	if !present {
		a.logger.Panic("Called commit for non-existent transaction")
	}
	tran.stopMonitor <- true
	tran.commands <- "do_commit"
	<-tran.commands
	a.terminated[id] = "committed"
	delete(a.transactions, id)
}

func (a *MemberlistAgent) monitorTransaction(transaction transactionInfo) {
	msg := message{
		Type:        "get_decision",
		Sender:      a.list.LocalNode(),
		Transaction: transaction,
	}
	msg.Transaction.stopMonitor = nil
	msg.Transaction.coordinated = false
	for {
		select {
		case <-transaction.stopMonitor:
			return
		case <-time.After(time.Millisecond * timeoutWakeMonitor):
			a.transactionMessages <- msg
		}
	}
}
