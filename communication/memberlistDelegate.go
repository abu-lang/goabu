// Copyright 2021 Massimo Comuzzo, Michele Pasqua and Marino Miculan
// SPDX-License-Identifier: Apache-2.0

package communication

import (
	"github.com/hashicorp/memberlist"
	"go.uber.org/zap"
)

type BaseMembers struct {
	// AgentID is the agent's identifier.
	AgentID         string
	ListeningPort   int
	List            *memberlist.Memberlist
	Config          *memberlist.Config
	ReceivedActions func() (<-chan chan []byte, <-chan chan string)
	Logger          *zap.Logger
}

type MemberlistDelegate interface {
	Start(BaseMembers)
	Stop(BaseMembers)
	FilterParticipants(BaseMembers, []*memberlist.Node) []*memberlist.Node

	NotifyMsg(BaseMembers, []byte)
	GetBroadcasts(BaseMembers, int, int) [][]byte
	LocalState(BaseMembers, bool) []byte
	MergeRemoteState(BaseMembers, []byte, bool)
	NotifyJoin(BaseMembers, *memberlist.Node)
	NotifyLeave(BaseMembers, *memberlist.Node)
	NotifyUpdate(BaseMembers, *memberlist.Node)
}
