package communication

import (
	"encoding/json"
	"errors"
	"fmt"
	"steel-lang/datastructure"
	"steel-lang/semantics"
	"time"

	"github.com/hashicorp/memberlist"
)

const (
	// milliseconds
	timeoutPhaseResend = 6000
	timeoutWakeMonitor = 12000
)

type transactionInfo struct {
	Initiator    string
	Number       int
	Actions      []semantics.ExternalAction
	Partecipants []string
	stopMonitor  chan bool
	coordinated  bool
}

func (t *transactionInfo) id() string {
	return fmt.Sprintf("%s->%d", t.Initiator, t.Number)
}

func (t *transactionInfo) buryPartecipants(members []*memberlist.Node) {
	alives := datastructure.MakeStringSet("")
	for _, member := range members {
		alives.Insert(member.Name)
	}
	buried := 0
	for i, partecipant := range t.Partecipants {
		if !alives.Contains(partecipant) {
			t.Partecipants[i] = ""
			buried++
		}
	}
	remaining := len(t.Partecipants) - buried
	j := 1
	for i := 0; i < remaining; i++ {
		if t.Partecipants[i] == "" {
			for t.Partecipants[j] == "" {
				j++
			}
			t.Partecipants[i], t.Partecipants[j] = t.Partecipants[j], t.Partecipants[i]
		}
		j++
	}
	t.Partecipants = t.Partecipants[:remaining]
}

type transactionChannels struct {
	Initiator     string
	Number        int
	arePrepared   chan string
	haveAborted   chan string
	haveCommitted chan string
}

func makeTransactionChannels(t transactionInfo) transactionChannels {
	return transactionChannels{
		Initiator:     t.Initiator,
		Number:        t.Number,
		arePrepared:   make(chan string, msgBuffLen),
		haveAborted:   make(chan string, msgBuffLen),
		haveCommitted: make(chan string, msgBuffLen),
	}
}

func (t transactionChannels) id() string {
	return fmt.Sprintf("%s->%d", t.Initiator, t.Number)
}

func (a *memberlistAgent) possiblyInterested(actions []semantics.ExternalAction) []string {
	var res []string
	for _, member := range a.list.Members() {
		a.lockRegistry.RLock()
		resources, present := a.registry[member.Name]
		a.lockRegistry.RUnlock()
		if !present { // I do not know the resources of member
			res = append(res, member.Name)
			continue
		}
		if resources == nil { // member is leaving
			continue
		}
		for _, action := range actions { // member should have at least the resources for executing one action
			if resources.ContainsSet(action.WorkingSet) {
				res = append(res, member.Name)
				continue
			}
		}
	}
	return res
}

func (a *memberlistAgent) coordinateTransaction(tran transactionInfo) error {
	canCommit := messageUnion{
		Type:        "can_commit?",
		Sender:      a.list.LocalNode(),
		Transaction: tran,
	}
	msg, err := json.Marshal(canCommit)
	if err != nil {
		return err
	}
	receivers := datastructure.MakeStringSet("")
	for _, nodeName := range tran.Partecipants {
		receivers.Insert(nodeName)
	}
	channels := makeTransactionChannels(tran)
	channelsCh := make(chan transactionChannels)
	a.coordinatedChannels <- channelsCh
	channelsCh <- channels
	fmt.Printf("started transaction with %d partecipants\n", receivers.Size())
	res := a.firstPhase(receivers, msg, channels)
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
	msg, err = json.Marshal(order)
	if err != nil {
		panic(err)
	}
	a.secondPhase(receivers, msg, responses)
	channelsCh = make(chan transactionChannels)
	a.coordinatedChannels <- channelsCh
	channelsCh <- transactionChannels{
		Initiator: tran.Initiator,
		Number:    tran.Number,
	}
	fmt.Println("terminated transaction")
	return res
}

func (a *memberlistAgent) firstPhase(partecipants datastructure.StringSet, msg []byte, channels transactionChannels) error {
	waitFor := partecipants.Clone()
	for !waitFor.Empty() {
		var timeout <-chan time.Time = nil
		waitForCopy := waitFor.Clone()
		receiversCh := make(chan datastructure.StringSet)
		go a.phaseSend(waitForCopy, msg, true, receiversCh)
	GET_RESPONSES_1:
		for !waitFor.Empty() {
			select {
			case receivers := <-receiversCh:
				timeout = time.After(time.Millisecond * timeoutPhaseResend)
				waitFor.Intersect(receivers)
			case prepared := <-channels.arePrepared:
				delete(waitFor, prepared)
			case <-channels.haveCommitted: // I am substituting initiator
				return nil
			case aborted := <-channels.haveAborted:
				return fmt.Errorf("%s has aborted", aborted)
			case <-timeout:
				break GET_RESPONSES_1
			}
		}
	}
	return nil
}

func (a *memberlistAgent) secondPhase(waitFor datastructure.StringSet, msg []byte, responses <-chan string) {
	for !waitFor.Empty() {
		var timeout <-chan time.Time = nil
		waitForCopy := waitFor.Clone()
		receiversCh := make(chan datastructure.StringSet)
		go a.phaseSend(waitForCopy, msg, false, receiversCh)
	GET_RESPONSES_2:
		for !waitFor.Empty() {
			select {
			case receivers := <-receiversCh:
				timeout = time.After(time.Millisecond * timeoutPhaseResend)
				waitFor.Intersect(receivers)
			case responded := <-responses:
				waitFor.Remove(responded)
			case <-timeout:
				break GET_RESPONSES_2
			}
		}
	}
}

func (a *memberlistAgent) phaseSend(receivers datastructure.StringSet, msg []byte, reliableSend bool, done chan<- datastructure.StringSet) {
	newReceivers := datastructure.MakeStringSet("")
	for _, member := range a.list.Members() {
		if receivers.Contains(member.Name) {
			newReceivers.Insert(member.Name)
			if reliableSend {
				a.list.SendReliable(member, msg)
			} else {
				a.list.SendBestEffort(member, msg)
			}
		}
	}
	done <- newReceivers
}

func demuxResponses(coordinated <-chan chan transactionChannels, responses <-chan messageUnion, quit <-chan chan bool) {
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
			if len(lines) == 0 {
				return
			}
		case response := <-responses:
			channels, present := lines[response.Transaction.id()]
			if !present {
				fmt.Println("transaction already terminated discarding response from", response.Sender.Name)
				break
			}
			switch response.Type {
			case "prepared":
				select {
				case channels.arePrepared <- response.Sender.Name:
					fmt.Printf("%s is prepared\n", response.Sender.Name)
				default:
					fmt.Println("discarded prepared status from", response.Sender.Name)
				}
			case "aborted":
				select {
				case channels.haveAborted <- response.Sender.Name:
					fmt.Printf("%s has aborted\n", response.Sender.Name)
				default:
					fmt.Println("discarded aborted status from", response.Sender.Name)
				}
			case "committed":
				select {
				case channels.haveCommitted <- response.Sender.Name:
					fmt.Printf("%s has committed\n", response.Sender.Name)
				default:
					fmt.Println("discarded committed status from", response.Sender.Name)
				}
			}
		}
	}

}

func (a *memberlistAgent) handleTransactions() {
	stopping := false
	for {
		select {
		case c := <-a.quitTransactions:
			stopping = true
			defer func() { c <- true }()
			if a.transaction.Initiator == "" {
				return
			}
		case message := <-a.transactionMessages:
			switch message.Type {
			case "can_commit?":
				status := a.getStatus(message.Transaction)
				if status == "new" {
					if a.transaction.Initiator != "" {
						break
					}
					a.transaction = message.Transaction
					a.transaction.stopMonitor = make(chan bool)
					go a.monitorTransaction(a.transaction)
					status = "prepared"
				}
				response := messageUnion{
					Type:   status,
					Sender: a.list.LocalNode(),
					Transaction: transactionInfo{
						Initiator: message.Transaction.Initiator,
						Number:    message.Transaction.Number,
					},
				}
				responseMsg, err := json.Marshal(response)
				if err != nil {
					fmt.Println("can_commit?: error during response marshalling,", err.Error())
					break
				}
				a.list.SendBestEffort(message.Sender, responseMsg)
			case "do_abort":
				status := a.getStatus(message.Transaction)
				if status == "commited" {
					panic(errors.New("received do abort for a committed tran"))
				}
				if status == "new" {
					fmt.Println("do_abort on new transaction")
					a.terminated[message.Transaction.id()] = "aborted"
				}
				if status == "prepared" {
					a.transaction.stopMonitor <- true
					a.terminated[message.Transaction.id()] = "aborted"
					a.transaction = transactionInfo{Initiator: ""}
				}
				response := messageUnion{
					Type:   "aborted",
					Sender: a.list.LocalNode(),
					Transaction: transactionInfo{
						Initiator: message.Transaction.Initiator,
						Number:    message.Transaction.Number,
					},
				}
				responseMsg, err := json.Marshal(response)
				if err != nil {
					fmt.Println("do_abort: error during response marshalling,", err.Error())
					break
				}
				a.list.SendBestEffort(message.Sender, responseMsg)
				if stopping && a.transaction.Initiator == "" {
					return
				}
			case "do_commit":
				status := a.getStatus(message.Transaction)
				if status != "prepared" && status != "committed" {
					panic(errors.New("spurious do_commit"))
				}
				if status == "prepared" {
					a.transaction.stopMonitor <- true
					actionsCh := make(chan []semantics.ExternalAction)
					a.committedOperations <- actionsCh
					actionsCh <- a.transaction.Actions
					<-actionsCh
					a.terminated[message.Transaction.id()] = "committed"
					a.transaction = transactionInfo{Initiator: ""}
				}
				response := messageUnion{
					Type:   "committed",
					Sender: a.list.LocalNode(),
					Transaction: transactionInfo{
						Initiator: message.Transaction.Initiator,
						Number:    message.Transaction.Number,
					},
				}
				responseMsg, err := json.Marshal(response)
				if err != nil {
					fmt.Println("do_commit: error during response marshalling,", err.Error())
					break
				}
				a.list.SendBestEffort(message.Sender, responseMsg)
				if stopping && a.transaction.Initiator == "" {
					return
				}
			case "get_decision":
				status := a.getStatus(message.Transaction)
				if status == "new" {
					fmt.Println("get_decision on new transaction")
					if a.transaction.Initiator == "" {
						a.transaction = message.Transaction
						a.transaction.stopMonitor = make(chan bool)
						go a.monitorTransaction(a.transaction)
						select {
						case a.transactionMessages <- message:
						default:
						}
					}
					break
				}
				if status == "prepared" {
					if a.transaction.coordinated {
						message.Transaction.Actions = nil
						message.Transaction.Partecipants = nil
						message.Type = "prepared"
						select {
						case a.transactionResponses <- message:
						default:
						}
						break
					}
					a.transaction.buryPartecipants(a.list.Members())
					if len(a.transaction.Partecipants) == 0 {
						panic(errors.New("empty partecipant list"))
					}
					head := a.transaction.Partecipants[0]
					if head == a.list.LocalNode().Name && a.transaction.Initiator != a.list.LocalNode().Name {
						fmt.Println("initiator", a.transaction.Initiator, "is dead: becoming new coordinator")
						go a.coordinateTransaction(a.transaction)
						a.transaction.coordinated = true
						break
					}
					message.Sender = a.list.LocalNode()
					deflected, err := json.Marshal(message)
					if err != nil {
						fmt.Println("get_decision: error during deflected message marshalling,", err.Error())
						break
					}
					for _, member := range a.list.Members() {
						if member.Name == head {
							a.list.SendReliable(member, deflected)
							break
						}
					}
					break
				}
				response := messageUnion{
					Sender: a.list.LocalNode(),
					Transaction: transactionInfo{
						Initiator: message.Transaction.Initiator,
						Number:    message.Transaction.Number,
					},
				}
				response.Type = "do_abort"
				if status == "committed" {
					response.Type = "do_commit"
				}
				fmt.Println("get_decision from", message.Sender.Name, "responding", response.Type)
				responseMsg, err := json.Marshal(response)
				if err != nil {
					fmt.Println("get_decision: error during response marshalling,", err.Error())
					break
				}
				a.list.SendBestEffort(message.Sender, responseMsg)
			default:
				fmt.Println("unsupported message:", message.Type)
			}
		}
	}
}

func (a *memberlistAgent) getStatus(tran transactionInfo) string {
	outcome, present := a.terminated[tran.id()]
	if present {
		return outcome
	}
	if a.transaction.id() == tran.id() {
		return "prepared"
	}
	return "new"
}

func (a *memberlistAgent) monitorTransaction(transaction transactionInfo) {
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
