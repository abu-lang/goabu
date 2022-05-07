package communication

import (
	"errors"
	"sync"
	"time"

	"github.com/hashicorp/memberlist"
	"go.uber.org/zap"
)

type delegateAdapter struct {
	listPtr              **memberlist.Memberlist
	trackGossip          chan chan *sync.WaitGroup
	transactionMessages  chan messageUnion
	transactionResponses chan messageUnion
	members              BaseMembers
	delegate             MemberlistDelegate
}

func (d delegateAdapter) delegateMembers() BaseMembers {
	res := d.members
	res.List = *d.listPtr
	return res
}

func (d delegateAdapter) start() {
	d.delegate.Start(d.delegateMembers())
}

func (d delegateAdapter) stop() {
	d.delegate.Stop(d.delegateMembers())
}

func (d delegateAdapter) filterParticipants(nodes []*memberlist.Node) []*memberlist.Node {
	return d.delegate.FilterParticipants(d.delegateMembers(), nodes)
}

func (d delegateAdapter) register() (*sync.WaitGroup, error) {
	replyCh := make(chan *sync.WaitGroup)
	select {
	case d.trackGossip <- replyCh:
		return <-replyCh, nil
	case <-time.After(time.Millisecond * timeoutRegister):
		return nil, errors.New("timeout in waiting from joiner")
	}
}

// NodeMeta implements memberlist.Delegate.NodeMeta.
// It returns the id of the agent as a possibly truncated []byte.
func (d delegateAdapter) NodeMeta(limit int) []byte {
	group, err := d.register()
	if err != nil {
		d.members.Logger.Error(err.Error())
		return []byte{}
	}
	defer group.Done()

	res := []byte(d.members.AgentID)
	if len(res) > limit {
		res = res[0:limit:limit]
	}
	return res
}

// NotifyMsg implements memberlist.Delegate.NotifyMsg.
// If the agent is running it sends m to the agent if m is a message of the
// transaction handling protocol otherwise if the agent is running and m is
// a different message then the delegate's NotifyMsg is called.
func (d delegateAdapter) NotifyMsg(m []byte) {
	group, err := d.register()
	if err != nil {
		d.members.Logger.Error(err.Error())
		return
	}
	defer group.Done()

	var message messageUnion
	ok := message.unmarshal(m)
	if ok {
		switch message.Type { // intercept transaction messages
		case "interested", "not_interested", "prepared", "aborted", "committed":
			select {
			case d.transactionResponses <- message:
			default:
				d.members.Logger.Warn("Dicarded transaction response",
					zap.String("act", "discard"),
					zap.String("obj", "transaction response"),
					zap.String("from", agentID(message.Sender)))
			}
			return
		case "interested?", "can_commit?", "do_commit", "do_abort", "get_decision":
			select {
			case d.transactionMessages <- message:
			default:
				d.members.Logger.Warn("Dicarded incoming transaction message",
					zap.String("act", "discard"),
					zap.String("obj", "transaction message"),
					zap.String("from", agentID(message.Sender)))
			}
			return
		}
	}

	d.delegate.NotifyMsg(d.delegateMembers(), m)
}

// GetBroadcasts implements memberlist.Delegate.GetBroadcasts.
// If the agent is still running it returns the result of the invocation
// of the delegate's GetBroadcast otherwise [][]byte{} is returned.
func (d delegateAdapter) GetBroadcasts(overhead, limit int) [][]byte {
	group, err := d.register()
	if err != nil {
		d.members.Logger.Error(err.Error())
		return [][]byte{}
	}
	defer group.Done()

	return d.delegate.GetBroadcasts(d.delegateMembers(), overhead, limit)
}

// LocalState implements memberlist.Delegate.LocalState.
// If the agent is still running it returns the result of the invocation
// of the delegate's LocalState otherwise []byte{} is returned.
func (d delegateAdapter) LocalState(join bool) []byte {
	group, err := d.register()
	if err != nil {
		d.members.Logger.Error(err.Error())
		return []byte{}
	}
	defer group.Done()

	return d.delegate.LocalState(d.delegateMembers(), join)
}

// MergeRemoteState implements memberlist.Delegate.MergeRemoteState.
// If the agent is still running it calls the delegate's MergeRemoteState.
func (d delegateAdapter) MergeRemoteState(buf []byte, join bool) {
	group, err := d.register()
	if err != nil {
		d.members.Logger.Error(err.Error())
		return
	}
	defer group.Done()

	d.delegate.MergeRemoteState(d.delegateMembers(), buf, join)
}

// NotifyJoin implements memberlist.EventDelegate.NotifyJoin.
// If the agent is still running it calls the delegate's NotifyJoin.
func (d delegateAdapter) NotifyJoin(node *memberlist.Node) {
	group, err := d.register()
	if err != nil {
		d.members.Logger.Error(err.Error())
		return
	}
	defer group.Done()

	d.delegate.NotifyJoin(d.delegateMembers(), node)
}

// NotifyLeave implements memberlist.EventDelegate.NotifyLeave.
// If the agent is still running it calls the delegate's NotifyLeave.
func (d delegateAdapter) NotifyLeave(node *memberlist.Node) {
	group, err := d.register()
	if err != nil {
		d.members.Logger.Error(err.Error())
		return
	}
	defer group.Done()

	d.delegate.NotifyLeave(d.delegateMembers(), node)
}

// NotifyUpdate implements memberlist.EventDelegate.NotifyUpdate.
// If the agent is still running it calls the delegate's NotifyUpdate.
func (d delegateAdapter) NotifyUpdate(node *memberlist.Node) {
	group, err := d.register()
	if err != nil {
		d.members.Logger.Error(err.Error())
		return
	}
	defer group.Done()

	d.delegate.NotifyUpdate(d.delegateMembers(), node)
}
