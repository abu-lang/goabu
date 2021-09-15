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

func (d delegateAdapter) filterPartecipants(nodes []*memberlist.Node) []*memberlist.Node {
	return d.delegate.FilterPartecipants(d.delegateMembers(), nodes)
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

func (d delegateAdapter) NodeMeta(limit int) []byte {
	group, err := d.register()
	if err != nil {
		d.members.Logger.Error(err.Error())
		return []byte{}
	}
	defer group.Done()

	return d.delegate.NodeMeta(d.delegateMembers(), limit)
}

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
					zap.String("from", message.Sender.Name))
			}
			return
		case "interested?", "can_commit?", "do_commit", "do_abort", "get_decision":
			select {
			case d.transactionMessages <- message:
			default:
				d.members.Logger.Warn("Dicarded incoming transaction message",
					zap.String("act", "discard"),
					zap.String("obj", "transaction message"),
					zap.String("from", message.Sender.Name))
			}
			return
		}
	}

	d.delegate.NotifyMsg(d.delegateMembers(), m)
}

func (d delegateAdapter) GetBroadcasts(overhead, limit int) [][]byte {
	group, err := d.register()
	if err != nil {
		d.members.Logger.Error(err.Error())
		return [][]byte{}
	}
	defer group.Done()

	return d.delegate.GetBroadcasts(d.delegateMembers(), overhead, limit)
}

func (d delegateAdapter) LocalState(join bool) []byte {
	group, err := d.register()
	if err != nil {
		d.members.Logger.Error(err.Error())
		return []byte{}
	}
	defer group.Done()

	return d.delegate.LocalState(d.delegateMembers(), join)
}

func (d delegateAdapter) MergeRemoteState(buf []byte, join bool) {
	group, err := d.register()
	if err != nil {
		d.members.Logger.Error(err.Error())
		return
	}
	defer group.Done()

	d.delegate.MergeRemoteState(d.delegateMembers(), buf, join)
}

func (d delegateAdapter) NotifyJoin(node *memberlist.Node) {
	group, err := d.register()
	if err != nil {
		d.members.Logger.Error(err.Error())
		return
	}
	defer group.Done()

	d.delegate.NotifyJoin(d.delegateMembers(), node)
}

func (d delegateAdapter) NotifyLeave(node *memberlist.Node) {
	group, err := d.register()
	if err != nil {
		d.members.Logger.Error(err.Error())
		return
	}
	defer group.Done()

	d.members.Logger.Info("Node "+node.Name+" has left",
		zap.String("act", "leave"),
		zap.String("subj", node.Name))
	defer d.members.Logger.Sync()

	d.delegate.NotifyLeave(d.delegateMembers(), node)
}

func (d delegateAdapter) NotifyUpdate(node *memberlist.Node) {
	group, err := d.register()
	if err != nil {
		d.members.Logger.Error(err.Error())
		return
	}
	defer group.Done()

	d.delegate.NotifyUpdate(d.delegateMembers(), node)
}
