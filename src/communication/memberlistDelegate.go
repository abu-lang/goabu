package communication

import (
	"github.com/hashicorp/memberlist"
	"go.uber.org/zap"
)

type BaseMembers struct {
	List            *memberlist.Memberlist
	Config          *memberlist.Config
	Terminated      map[string]string
	Transactions    map[string]*transactionInfo
	ListeningPort   int
	ReceivedActions func() (<-chan chan []byte, <-chan chan string)
	Logger          *zap.Logger
}

type MemberlistDelegate interface {
	Start(BaseMembers)
	Stop(BaseMembers)
	FilterPartecipants(BaseMembers, []*memberlist.Node) []*memberlist.Node

	NodeMeta(BaseMembers, int) []byte
	NotifyMsg(BaseMembers, []byte)
	GetBroadcasts(BaseMembers, int, int) [][]byte
	LocalState(BaseMembers, bool) []byte
	MergeRemoteState(BaseMembers, []byte, bool)
	NotifyJoin(BaseMembers, *memberlist.Node)
	NotifyLeave(BaseMembers, *memberlist.Node)
	NotifyUpdate(BaseMembers, *memberlist.Node)
}
