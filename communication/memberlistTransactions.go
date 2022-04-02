package communication

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/abu-lang/goabu/stringset"

	"github.com/hashicorp/memberlist"
	"go.uber.org/zap"
)

// TODO evaluate
const (
	// milliseconds
	timeoutPhaseResend = 6000
	timeoutWakeMonitor = 12000
)

type transactionInfo struct {
	Initiator    string
	Number       int
	Payload      []byte
	Participants []string
	stopMonitor  chan bool
	coordinated  bool
	commands     chan string
}

func (t *transactionInfo) id() string {
	return fmt.Sprintf("%s->%d", t.Initiator, t.Number)
}

func (t *transactionInfo) buryParticipants(members []*memberlist.Node) {
	alives := stringset.Make()
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
	message := messageUnion{
		Type:        "interested?",
		Sender:      a.list.LocalNode(),
		Transaction: tran,
	}
	msg, ok := message.marshal("interested?", a.logger)
	if !ok {
		a.logger.Panic("Could not marshal interested? message",
			zap.String("act", "marshalling"),
			zap.String("obj", "interested?"))
	}
	channels := makeTransactionChannels(tran)
	channelsCh := make(chan transactionChannels)
	a.coordinatedChannels <- channelsCh
	channelsCh <- channels
	nodes, err := a.interestPhase(msg, channels, tran.Number)
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

func (a *MemberlistAgent) interestPhase(msg []byte, channels transactionChannels, number int) ([]string, error) {
	aborted := ""
	waitFor := stringset.Make()
	for _, member := range a.adapter.filterParticipants(a.list.Members()) {
		waitFor.Insert(member.Name)
	}
	var interested []string
	for !waitFor.Empty() {
		var timeout <-chan time.Time = nil
		waitForCopy := waitFor.Clone()
		receiversCh := make(chan stringset.Set)
		if a.test == TestsMidInterested {
			go a.testsPhaseSend(waitForCopy, msg, true, receiversCh, TestsMidSends)
		} else {
			go a.phaseSend(waitForCopy, msg, true, receiversCh)
		}
		received := 0
	INTERESTED:
		for !waitFor.Empty() {
			select {
			case receivers := <-receiversCh:
				timeout = time.After(time.Millisecond * timeoutPhaseResend)
				waitFor.Intersect(receivers)
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
	order := messageUnion{
		Type:   "do_abort",
		Sender: a.list.LocalNode(),
		Transaction: transactionInfo{
			Initiator: a.list.LocalNode().Name,
			Number:    number,
		},
	}
	abrt, ok := order.marshal(order.Type, a.logger)
	if !ok {
		a.logger.Panic("Could not marshal "+order.Type+" message", zap.String("act", "marshalling"), zap.String("obj", order.Type))
	}
	dests := stringset.Make()
	for _, i := range interested {
		dests.Insert(i)
	}
	a.secondPhase(dests, abrt, channels.haveAborted)
	return nil, fmt.Errorf("%s has aborted", aborted)
}

func (a *MemberlistAgent) coordinateTransaction(tran transactionInfo) error {
	canCommit := messageUnion{
		Type:        "can_commit?",
		Sender:      a.list.LocalNode(),
		Transaction: tran,
	}
	msg, ok := canCommit.marshal("can_commit?", a.logger)
	if !ok {
		return errors.New("could not marshal can_commit? message")
	}
	receivers := stringset.Make()
	for _, nodeName := range tran.Participants {
		receivers.Insert(nodeName)
	}
	channels := makeTransactionChannels(tran)
	channelsCh := make(chan transactionChannels)
	a.coordinatedChannels <- channelsCh
	channelsCh <- channels
	a.logger.Debug("Started transaction",
		zap.String("subj", canCommit.Sender.Name),
		zap.String("act", "start_tran"),
		zap.Int("participants", receivers.Size()))
	res := a.firstPhase(receivers, msg, channels)
	a.logger.Debug("Terminated first phase",
		zap.String("subj", canCommit.Sender.Name),
		zap.String("act", "end_1_phase"),
		zap.Int("participants", receivers.Size()))
	a.testsHaltIf(TestsAfterFirst)
	responses := channels.haveCommitted
	action := "do_commit"
	if res != nil {
		responses = channels.haveAborted
		action = "do_abort"
	}
	order := messageUnion{
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
	a.secondPhase(receivers, msg, responses)
	channelsCh = make(chan transactionChannels)
	a.coordinatedChannels <- channelsCh
	channelsCh <- transactionChannels{
		Initiator: tran.Initiator,
		Number:    tran.Number,
	}
	a.logger.Debug("Terminated transaction", zap.String("subj", a.list.LocalNode().Name), zap.String("act", "end_tran"))
	return res
}

func (a *MemberlistAgent) firstPhase(participants stringset.Set, msg []byte, channels transactionChannels) error {
	waitFor := participants.Clone()
	for !waitFor.Empty() {
		var timeout <-chan time.Time = nil
		waitForCopy := waitFor.Clone()
		receiversCh := make(chan stringset.Set)
		if a.test == TestsMidFirst {
			go a.testsPhaseSend(waitForCopy, msg, true, receiversCh, TestsMidSends)
		} else {
			go a.phaseSend(waitForCopy, msg, true, receiversCh)
		}
		received := 0
	GET_RESPONSES_1:
		for !waitFor.Empty() {
			select {
			case receivers := <-receiversCh:
				timeout = time.After(time.Millisecond * timeoutPhaseResend)
				waitFor.Intersect(receivers)
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

func (a *MemberlistAgent) secondPhase(waitFor stringset.Set, msg []byte, responses <-chan string) {
	for !waitFor.Empty() {
		var timeout <-chan time.Time = nil
		waitForCopy := waitFor.Clone()
		receiversCh := make(chan stringset.Set)
		if a.test == TestsMidSecond {
			go a.testsPhaseSend(waitForCopy, msg, false, receiversCh, TestsMidSends)
		} else {
			go a.phaseSend(waitForCopy, msg, false, receiversCh)
		}
		received := 0
	GET_RESPONSES_2:
		for !waitFor.Empty() {
			select {
			case receivers := <-receiversCh:
				timeout = time.After(time.Millisecond * timeoutPhaseResend)
				waitFor.Intersect(receivers)
			case responded := <-responses:
				received++
				waitFor.Remove(responded)
			case <-timeout:
				break GET_RESPONSES_2
			}
			if a.test == TestsMidSecond && received >= TestsMidSends {
				a.testsHaltAndBlock()
			}
		}
	}
}

func (a *MemberlistAgent) phaseSend(receivers stringset.Set, msg []byte, reliableSend bool, done chan<- stringset.Set) {
	newReceivers := stringset.Make()
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
		}
	}
	done <- newReceivers
}

func (a *MemberlistAgent) testsPhaseSend(receivers stringset.Set, msg []byte, reliableSend bool, done chan<- stringset.Set, haltAfter int) {
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

func demuxResponses(coordinated <-chan chan transactionChannels, responses <-chan messageUnion, quit <-chan chan bool, logger *zap.Logger) {
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
					zap.String("from", response.Sender.Name))
				break
			}
			c := channels.line(response.Type)
			select {
			case c <- response.Sender.Name:
				logger.Debug("Received: "+response.Type,
					zap.String("act", "recv"),
					zap.String("obj", response.Type),
					zap.String("from", response.Sender.Name))
			default:
				logger.Warn("Discarded interested status",
					zap.String("act", "discard"),
					zap.String("obj", response.Type),
					zap.String("from", response.Sender.Name))
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
		case message := <-a.transactionMessages:
			response := messageUnion{
				Sender: a.list.LocalNode(),
				Transaction: transactionInfo{
					Initiator: message.Transaction.Initiator,
					Number:    message.Transaction.Number,
				},
			}
			respond := true
			id := message.Transaction.id()
			status := a.getStatus(id)
			switch message.Type {

			case "__interested__", "__not_interested__", "__aborted__":
				if status != "evaluating" {
					a.logger.Panic("Received evaluation result when not evaluating", zap.String("act", "recv"), zap.String("obj", message.Type))
				}
				tran := a.transactions[id]
				response.Type = strings.Trim(message.Type, "_")
				if response.Type == "interested" {
					tran.stopMonitor = make(chan bool)
					go a.monitorTransaction(*tran)
				} else {
					a.terminated[id] = response.Type
				}
				respond = false
				for _, node := range a.list.Members() {
					if node.Name == tran.Initiator {
						message.Sender = node
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
					actionsCh <- message.Transaction.Payload
					tran := &transactionInfo{}
					*tran = message.Transaction
					tran.Payload = nil
					tran.Participants = nil
					go func(tran transactionInfo) {
						a.transactionMessages <- messageUnion{
							Sender:      a.list.LocalNode(),
							Type:        "__" + <-commandsCh + "__",
							Transaction: tran,
						}
					}(*tran)
					tran.commands = commandsCh
					a.transactions[id] = tran
					respond = false
				case "interested", "not_interested":
					response.Type = status
				default:
					respond = false
				}

			case "can_commit?":
				a.logger.Debug("Received: can_commit?",
					zap.String("subj", a.list.LocalNode().Name),
					zap.String("act", "recv"),
					zap.String("obj", "can_commit?"),
					zap.String("from", message.Sender.Name))
				switch status {
				case "not_interested", "evaluating":
					a.logger.Panic("Received can_commit? but I am evaluating or I wasn't interested",
						zap.String("act", "recv"),
						zap.String("obj", "can_commit?"),
						zap.String("from", message.Sender.Name))
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
						tran.Participants = message.Transaction.Participants
					}
				default:
					response.Type = status
				}

			case "do_abort":
				a.logger.Debug("Received: do_abort",
					zap.String("subj", a.list.LocalNode().Name),
					zap.String("act", "recv"),
					zap.String("obj", "do_abort"),
					zap.String("from", message.Sender.Name))
				switch status {
				case "evaluating":
					select {
					case a.transactionMessages <- message:
					default:
					}
					respond = false
				case "committed", "new", "not_interested":
					a.logger.Panic("Received do_abort for a committed, new or uninteresting transaction",
						zap.String("act", "recv"),
						zap.String("obj", "do_abort"),
						zap.String("from", message.Sender.Name))
				case "prepared", "interested":
					a.abort(id)
				}
				response.Type = "aborted"

			case "do_commit":
				a.logger.Debug("Received: do_commit",
					zap.String("subj", a.list.LocalNode().Name),
					zap.String("act", "recv"),
					zap.String("obj", "do_commit"),
					zap.String("from", message.Sender.Name))
				switch status {
				case "prepared":
					a.commit(id)
				case "committed":
				default:
					a.logger.Panic("Spurious do_commit",
						zap.String("act", "recv"),
						zap.String("obj", "do_commit"),
						zap.String("from", message.Sender.Name))
				}
				response.Type = "committed"

			case "get_decision":
				switch status {
				case "new", "evaluating", "not_interested":
					a.logger.Panic("Received get_decision for a new, evaluating or uninteresting transaction",
						zap.String("act", "recv"),
						zap.String("obj", "get_decision"),
						zap.String("from", message.Sender.Name))
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
							zap.String("subj", a.list.LocalNode().Name),
							zap.String("act", "abort_tran"),
							zap.String("initiator", tran.Initiator))
						a.abort(id)
						break
					}
					if tran.coordinated {
						message.Transaction.Payload = nil
						message.Transaction.Participants = nil
						message.Type = "prepared"
						select {
						case a.transactionResponses <- message:
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
							zap.String("subj", head),
							zap.String("act", "coord"),
							zap.String("initiator", tran.Initiator))
						go a.coordinateTransaction(*tran)
						tran.coordinated = true
						break
					}
					message.Sender = a.list.LocalNode()
					deflected, ok := message.marshal("deflected message", a.logger)
					if ok {
						for _, member := range a.list.Members() {
							if member.Name == head {
								a.list.SendReliable(member, deflected)
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
						zap.String("subj", a.list.LocalNode().Name),
						zap.String("act", "send"),
						zap.String("obj", response.Type),
						zap.String("to", message.Sender.Name))
				}

			default:
				respond = false
				a.logger.DPanic("Unsupported transaction message: "+message.Type,
					zap.String("act", "recv"),
					zap.String("obj", `"`+message.Type+`"`),
					zap.String("from", message.Sender.Name))
			}

			if respond {
				responseMsg, ok := response.marshal("response", a.logger)
				if ok {
					a.list.SendBestEffort(message.Sender, responseMsg)
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
	message := messageUnion{
		Type:        "get_decision",
		Sender:      a.list.LocalNode(),
		Transaction: transaction,
	}
	message.Transaction.stopMonitor = nil
	message.Transaction.coordinated = false
	for {
		select {
		case <-transaction.stopMonitor:
			return
		case <-time.After(time.Millisecond * timeoutWakeMonitor):
			a.transactionMessages <- message
		}
	}
}
