package communication

import (
	"github.com/hashicorp/memberlist"
	"go.uber.org/zap"
)

type delegateDefault struct{}

func (d delegateDefault) Start(b BaseMembers) {}

func (d delegateDefault) Stop(b BaseMembers) {}

func (d delegateDefault) FilterParticipants(b BaseMembers, nodes []*memberlist.Node) []*memberlist.Node {
	res := make([]*memberlist.Node, 0, len(nodes))
	myName := b.List.LocalNode().Name
	for _, n := range nodes {
		if n.Name != myName {
			res = append(res, n)
		}
	}
	return res
}

func (d delegateDefault) NodeMeta(b BaseMembers, limit int) []byte {
	return []byte{}
}

func (d delegateDefault) NotifyMsg(b BaseMembers, m []byte) {
	b.Logger.Error("Unsupported message",
		zap.String("act", "recv"),
		zap.Binary("obj", m))
}

func (d delegateDefault) GetBroadcasts(b BaseMembers, overhead, limit int) [][]byte {
	return [][]byte{}
}

func (d delegateDefault) LocalState(b BaseMembers, join bool) []byte {
	return []byte{}
}

func (d delegateDefault) MergeRemoteState(b BaseMembers, buf []byte, join bool) {}

func (d delegateDefault) NotifyJoin(b BaseMembers, node *memberlist.Node) {}

func (d delegateDefault) NotifyLeave(b BaseMembers, node *memberlist.Node) {}

func (d delegateDefault) NotifyUpdate(b BaseMembers, node *memberlist.Node) {}
