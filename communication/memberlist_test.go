// Copyright 2021 Massimo Comuzzo, Michele Pasqua and Marino Miculan
// SPDX-License-Identifier: Apache-2.0

package communication

import (
	"bytes"
	"errors"
	"fmt"
	"testing"

	"github.com/abu-lang/goabu/config"

	"github.com/google/uuid"
)

const (
	TestResAgree = iota
	TestResCommit
	TestResAbort
)

func TestNewMemberlistAgent(t *testing.T) {
	tests := []struct {
		index int
		id    string
		port  int
		nodes []string
	}{
		//  {_, id, port, nodes},
		{1, "agent", 0, nil},
		{2, "Alice", 8100, []string{}},
		{3, "bob", 8101, []string{"127.0.0.1:8150"}},
		{4, "12345678", 8102, []string{"127.0.0.1:8150,127.0.0.1:8151"}},
		{5, "", 8103, []string{"127.0.0.1:8151"}},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("TestNewMemberlistAgent#%d", test.index), func(t *testing.T) {
			agt := NewMemberlistAgent(test.id, test.port, config.TestsLogConfig, test.nodes...)
			if agt.IsRunning() {
				t.Error("agent should not be running")
			}
			if agt.id == "" {
				t.Error("id should not be \"\"")
			}
			if agt.initiatedTransactions != 0 {
				t.Error("initiatedTransactions should be 0")
			}
			if agt.listeningPort != test.port {
				t.Errorf("listeningPort should be %d", test.port)
			}
			if agt.operations == nil {
				t.Error("operations should not be nil")
			}
			if agt.operationCommands == nil {
				t.Error("opeartionCommands should not be nil")
			}
			if agt.test != TestsNothing {
				t.Error("test should be TestNothing")
			}
			if agt.halted {
				t.Error("halted should be false")
			}
			if agt.lockHalted != nil {
				t.Error("lockHalted should be nil")
			}
			checkCorrectStop(t, agt)
		})
	}
}

func TestStart(t *testing.T) {
	tests := []struct {
		index int
		port  int
		nodes []string
	}{
		//  {_, port, nodes},
		{1, 0, nil},
		{2, 9100, []string{}},
		{3, 9101, []string{"127.0.0.1:9150"}},
		{4, 9102, []string{"127.0.0.1:9150,127.0.0.1:9151"}},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("TestStart#%d", test.index), func(t *testing.T) {
			agt := NewMemberlistAgent(t.Name(), test.port, config.TestsLogConfig, test.nodes...)
			start(t, agt, test.port)
			err := agt.Start()
			if err == nil {
				t.Error("start should return error when agent is already running")
			}
		})
	}
}

func TestJoin(t *testing.T) {
	tests := []struct {
		index int
		port  int
		nodes []string
		start bool
		good  bool
	}{
		//  {_, set, port, nodes, start, good},
		{1, 0, nil, true, true},
		{2, 10100, []string{}, true, true},
		{3, 0, nil, false, false},
		{4, 10101, []string{"127.0.0.1:10150"}, true, false},
		{5, 10102, []string{"127.0.0.1:10101"}, true, true},
		{6, 10103, []string{"127.0.0.1:10101", "127.0.0.1:10102"}, true, true},
		{7, 10104, []string{"exercitation,ullamco"}, true, false},
		{8, 10105, []string{"! #\"($/(,.123456nostrud"}, true, false},
		{9, 10106, []string{"127.0.0.1:10150"}, false, false},
		{10, 10107, []string{"127.0.0.1:10101", "127.0.0.1:10102"}, false, false},
		{11, 10108, []string{".,>Z><<-@#00asdfg"}, false, false},
	}
	dummy := make([]*MemberlistAgent, 0, len(tests))
	for i, test := range tests {
		t.Run(fmt.Sprintf("TestJoin#%d", test.index), func(t *testing.T) {
			dummy = append(dummy, NewMemberlistAgent(t.Name(), test.port, config.TestsLogConfig, test.nodes...))
			agt := dummy[i]
			if test.start {
				start(t, agt, test.port)
			}
			err := agt.Join()
			switch {
			case test.good && err != nil:
				t.Error(err.Error())
			case !test.good && err == nil:
				t.Error("Join should return error")
			}
			if !test.start {
				return
			}
			if test.good && len(test.nodes) > 0 {
				if agt.list.NumMembers() <= 1 {
					t.Error("list should have at least 2 members")
				}
			} else {
				if agt.list.NumMembers() != 1 {
					t.Error("list should have 1 member")
				}
			}
		})
	}
}

func TestForAll(t *testing.T) {
	const port = 0
	a := NewMemberlistAgent("TestForAll", port, config.TestsLogConfig)
	checkCorrectStop(t, a)
	err := a.ForAll([]byte(`lorem`))
	if err == nil {
		t.Error("ForAll should return error when agent is not running")
	}
	start(t, a, port)
	statusCh := localForAll(t, a, []byte(`incididunt ut labore et`), false)
	a.operations <- nil
	if !<-statusCh {
		t.Error("received wrong payload")
	}
	statusCh = localForAll(t, a, []byte(`dolore magna aliqua. Ut`), true)
	a.operations <- nil
	if !<-statusCh {
		t.Error("received wrong payload")
	}
}

func TestStop(t *testing.T) {
	const port = 11100
	a := NewMemberlistAgent("TestStop", port, config.TestsLogConfig)
	if a.Stop() == nil {
		t.Error("should return error when agent is not running")
	}
	start(t, a, port)
	uuid1 := a.list.LocalNode().Name
	startMockInterested(nil, a.operations, a.operationCommands)
	restart(t, a, port)
	uuid2 := a.list.LocalNode().Name
	if uuid1 == uuid2 {
		t.Error("uuid should be different after restart")
	}
	localForAll(t, a, []byte(`consectetur adipiscing elit`), false)
	restart(t, a, port)
	uuid3 := a.list.LocalNode().Name
	if uuid1 == uuid3 || uuid2 == uuid3 {
		t.Error("uuid should be different after restart")
	}
	localForAll(t, a, []byte(`, sed do eiusmod tempor`), true)
	stop(t, a)
}

func TestAborted(t *testing.T) {
	argsList := []struct {
		port int
		join []int
		test int
	}{
		{port: 12100},
		{port: 12101, join: []int{12100}},
		{port: 12102, join: []int{12101}, test: TestsAbort},
		{port: 12103, join: []int{12101}},
	}

	transactionHelper(t, makeAgents(t.Name(), argsList), []byte("commodo"), TestResAbort)
}

func TestUnreliable(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}

	argsList := []struct {
		port int
		join []int
		test int
	}{
		{port: 13100, test: TestsUnreliableSend},
		{port: 13101, join: []int{13100}},
		{port: 13102},
		{port: 13103, join: []int{13100, 13102}},
		{port: 13104, join: []int{13102}},
	}

	transactionHelper(t, makeAgents(t.Name(), argsList), []byte(". Duis aute123"), TestResCommit)
}

func TestInterestedMid(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}

	argsList := []struct {
		port int
		join []int
		test int
	}{
		{port: 14100, join: []int{14103}, test: TestsMidInterested},
		{port: 14101},
		{port: 14102},
		{port: 14103, join: []int{14101, 14102}},
		{port: 14104, join: []int{14100}},
	}

	transactionHelper(t, makeAgents(t.Name(), argsList), []byte("456reprehenderit in"), TestResAbort)
}

func TestInterestedAfter(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}

	argsList := []struct {
		port int
		join []int
		test int
	}{
		{port: 15100, join: []int{15102}, test: TestsAfterInterested},
		{port: 15101, join: []int{15100}},
		{port: 15102},
		{port: 15103, join: []int{15101}},
	}

	transactionHelper(t, makeAgents(t.Name(), argsList), []byte("velit esse.....@#"), TestResAbort)
}

func TestFirstMid(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}

	argsList := []struct {
		port int
		join []int
		test int
	}{
		{port: 16100, test: TestsMidFirst},
		{port: 16101, join: []int{16103}},
		{port: 16102},
		{port: 16103, join: []int{16100, 16102}},
	}

	transactionHelper(t, makeAgents(t.Name(), argsList), []byte("nulla pariatur. +-+-"), TestResAgree)
}

func TestFirstAfter(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}

	argsList := []struct {
		port int
		join []int
		test int
	}{
		{port: 17100, test: TestsAfterFirst},
		{port: 17101, join: []int{17102}},
		{port: 17102, join: []int{17100}},
		{port: 17103, join: []int{17100}},
	}

	transactionHelper(t, makeAgents(t.Name(), argsList), []byte("**!sint occaecat"), TestResCommit)
}

func TestSecondMid(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}

	argsList := []struct {
		port int
		join []int
		test int
	}{
		{port: 18100, test: TestsMidSecond},
		{port: 18101, join: []int{18103}},
		{port: 18102, join: []int{18100}},
		{port: 18103, join: []int{18102}},
	}

	transactionHelper(t, makeAgents(t.Name(), argsList), []byte("proident, sunt in"), TestResCommit)
}

func TestDeadlockExample(t *testing.T) {
	payload := []byte("deadlock_example")

	argsList := []struct {
		port int
		join []int
		test int
	}{
		{port: 19100},
		{port: 19101, join: []int{19100}},
		{port: 19102, join: []int{19101}},
		{port: 19103, join: []int{19102}},
		{port: 19104, join: []int{19103}},
	}
	agents := makeAgents(t.Name(), argsList)
	for i, agt := range agents {
		t.Run(fmt.Sprintf("ClusterMemberStart#%d", i+1), func(t *testing.T) {
			start(t, agt, agt.listeningPort)
		})
		startMockExec(agt.ReceivedActions())
	}
	for i, agt := range agents {
		t.Run(fmt.Sprintf("ClusterMemberJoin#%d", i+1), func(t *testing.T) {
			err := agt.Join()
			if err != nil {
				t.Fatal(err.Error())
			}
		})
	}
	for _, agt := range agents {
		for agt.list.NumMembers() != len(agents) {
		}
	}
	t.Log("Cluster is up, starting 2 transactions")
	waitTran := func(agt *MemberlistAgent) <-chan bool {
		res := make(chan bool)
		go func() {
			err := agt.ForAll(payload)
			res <- err == nil
		}()
		return res
	}
	t1 := waitTran(agents[0])
	t2 := waitTran(agents[1])
	select {
	case <-t1:
		t1 = nil
	case <-t2:
		t2 = nil
	}
	select {
	case <-t1:
	case <-t2:
	}
	fmt.Printf("%d agents: no deadlock happened\n", len(agents)) // gc
}

func start(t *testing.T, a *MemberlistAgent, p int) {
	t.Helper()
	err := a.Start()
	if err != nil {
		t.Fatal(err.Error())
	}
	checkCorrectStart(t, a, p)
}

func stop(t *testing.T, a *MemberlistAgent) {
	t.Helper()
	err := a.Stop()
	if err != nil {
		t.Fatal(err.Error())
	}
	checkCorrectStop(t, a)
}

func restart(t *testing.T, a *MemberlistAgent, p int) {
	t.Helper()
	stop(t, a)
	start(t, a, p)
}

func localForAll(t *testing.T, a *MemberlistAgent, payload []byte, interested bool) <-chan bool {
	t.Helper()
	ops, cmds := a.ReceivedActions()
	var res <-chan bool
	if interested {
		res = startMockInterested(payload, ops, cmds)
	} else {
		res = startMockUninterested(payload, ops, cmds)
	}
	err := a.ForAll(payload)
	if err != nil {
		t.Fatal(err.Error())
	}
	return res
}

func checkCorrectStart(t *testing.T, a *MemberlistAgent, p int) {
	t.Helper()
	if !a.IsRunning() {
		t.Error("should be running")
	}
	if p != 0 && a.config.BindPort != p {
		t.Errorf("BindPort should be %d", p)
	}
	_, err := uuid.Parse(a.config.Name)
	if err != nil {
		t.Error("Name should be a valid uuid")
	}
	if a.list.LocalNode().Name != a.config.Name || a.list.NumMembers() < 1 {
		t.Error("memberlist should be created")
	}
	if a.config.Name == a.id {
		t.Error("memberlist name shuld be different from agent id")
	}
	switch {
	case a.terminated == nil:
		t.Error("terminated should not be nil")
	case len(a.terminated) > 0:
		t.Error("terminated should be empty")
	case a.transactions == nil:
		t.Error("transactions should not be nil")
	case len(a.transactions) > 0:
		t.Error("transactions should be empty")
	case a.quitTransactions == nil:
		t.Error("quitTransactions should not be nil")
	case a.quitGossip == nil:
		t.Error("quitGossip should not be nil")
	case a.quitDemux == nil:
		t.Error("quitDemux should not be nil")
	case a.transactionMessages == nil:
		t.Error("transactionMessages should not be nil")
	case a.transactionResponses == nil:
		t.Error("transactionResponses should not be nil")
	case a.coordinatedChannels == nil:
		t.Error("coordinatedChannels should not be nil")
	case a.trackGossip == nil:
		t.Error("trackGossip should not be nil")
	}
}

func checkCorrectStop(t *testing.T, a *MemberlistAgent) {
	t.Helper()
	if a.IsRunning() {
		t.Error("should not be running")
	}
	switch {
	case a.terminated != nil:
		t.Error("terminated should be nil")
	case a.transactions != nil:
		t.Error("transactions should be nil")
	case a.list != nil:
		t.Error("list should be nil")
	case a.quitTransactions != nil:
		t.Error("quitTransactions should be nil")
	case a.quitGossip != nil:
		t.Error("quitGossip should be nil")
	case a.quitDemux != nil:
		t.Error("quitDemux should be nil")
	case a.transactionMessages != nil:
		t.Error("transactionMessages should be nil")
	case a.transactionResponses != nil:
		t.Error("transactionResponses should be nil")
	case a.coordinatedChannels != nil:
		t.Error("coordinatedChannels should be nil")
	case a.trackGossip != nil:
		t.Error("trackGossip should be nil")
	}
}

func makeAgents(test string, argsList []struct {
	port int
	join []int
	test int
}) []*MemberlistAgent {
	res := make([]*MemberlistAgent, 0, len(argsList))
	for i, args := range argsList {
		nodes := make([]string, 0, len(args.join))
		for _, p := range args.join {
			nodes = append(nodes, fmt.Sprintf("127.0.0.1:%d", p))
		}
		res = append(res, TestsNewMemberlistAgent(fmt.Sprintf("%s#%d", test, i+1), args.port, args.test, nodes...))
	}
	return res
}

func transactionHelper(t *testing.T, agents []*MemberlistAgent, payload []byte, outcome int) {
	t.Helper()
	if len(agents) == 0 {
		return
	}
	results := make([]<-chan int, 0, len(agents)-1)
	detectors := make([]<-chan bool, 0, len(agents)-1)
	agentIds := make(map[<-chan int]int)
	for i, agt := range agents {
		t.Run(fmt.Sprintf("ParticipantStart#%d", i+1), func(t *testing.T) {
			start(t, agt, agt.listeningPort)
		})
		if i == 0 {
			continue
		}
		if agt.test < TestsMidInterested {
			ops, cmds := agt.ReceivedActions()
			r := startMockCommit(payload, ops, cmds)
			agentIds[r] = i
			results = append(results, r)
		} else {
			d := startParticipantDetector(payload, agt)
			detectors = append(detectors, d)
		}
	}
	for i, agt := range agents {
		t.Run(fmt.Sprintf("ParticipantJoin#%d", i+1), func(t *testing.T) {
			err := agt.Join()
			if err != nil {
				t.Fatal(err.Error())
			}
		})
	}
	for _, agt := range agents {
		for agt.list.NumMembers() != len(agents) {
		}
	}
	go agents[0].ForAll(payload)
	detected := 0
	for _, detector := range detectors {
		if <-detector {
			detected++
		}
	}
	terminating := len(results)
	if agents[0].test == TestsMidInterested {
		terminating = TestsMidSends - detected
	}
	for terminating != 0 {
		for i := 0; i < len(results); i++ {
			select {
			case result := <-results[i]:
				if outcome == TestResAgree {
					outcome = result
				}
				if result != outcome {
					idx := agentIds[results[i]]
					if outcome == TestResCommit {
						t.Errorf("(%s) agent #%d should have committed", agents[idx].list.LocalNode().Name, idx+1)
					} else {
						t.Errorf("(%s) agent #%d should have aborted", agents[idx].list.LocalNode().Name, idx+1)
					}
				}
				l := len(results) - 1
				if i < l {
					results[i] = results[l]
				}
				results = results[:l]
				terminating--
			default:
			}
		}
	}
}

func startParticipantDetector(payload []byte, a *MemberlistAgent) <-chan bool {
	res := make(chan bool)
	go func() {
		halted := make(chan bool)
		requests, commandRequests := a.ReceivedActions()
		go func() {
			for {
				a.lockHalted.Lock()
				if a.halted {
					a.lockHalted.Unlock()
					break
				}
				a.lockHalted.Unlock()
			}
			halted <- true
		}()
		select {
		case actionsCh := <-requests:
			commandsCh := <-commandRequests
			if !bytes.Equal(<-actionsCh, payload) {
				panic(errors.New("received wrong payload"))
			}
			commandsCh <- "interested"
			res <- true
			switch <-commandsCh {
			case "can_commit?":
				commandsCh <- "prepared"
			case "do_abort":
				commandsCh <- "done"
				return
			}
			<-commandsCh
			commandsCh <- "done"
		case <-halted:
			res <- false
		}
	}()
	return res
}

func startMockCommit(payload []byte, requests <-chan chan []byte, commandRequests <-chan chan string) <-chan int {
	res := make(chan int)
	go func() {
		actionsCh := <-requests
		commandsCh := <-commandRequests
		if !bytes.Equal(<-actionsCh, payload) {
			panic(errors.New("received wrong payload"))
		}
		commandsCh <- "interested"
		switch <-commandsCh {
		case "can_commit?":
			commandsCh <- "prepared"
		case "do_abort":
			commandsCh <- "done"
			defer func() { res <- TestResAbort }()
			return
		default:
			panic(errors.New("illegal command"))
		}
		switch <-commandsCh {
		case "do_commit":
			defer func() { res <- TestResCommit }()
		case "do_abort":
			defer func() { res <- TestResAbort }()
		default:
			panic(errors.New("illegal command"))
		}
		commandsCh <- "done"
	}()
	return res
}

func startMockInterested(payload []byte, requests <-chan chan []byte, commandRequests <-chan chan string) <-chan bool {
	res := make(chan bool)
	go func() {
		good := true
		for {
			actionsCh := <-requests
			if actionsCh == nil {
				defer func() { res <- good }()
				return
			}
			commandsCh := <-commandRequests
			if !bytes.Equal(<-actionsCh, payload) {
				good = false
			}
			commandsCh <- "interested"
			switch <-commandsCh {
			case "can_commit?":
				commandsCh <- "prepared"
			case "do_abort":
				commandsCh <- "done"
				return
			}
			<-commandsCh
			commandsCh <- "done"
		}
	}()
	return res
}

func startMockUninterested(payload []byte, requests <-chan chan []byte, commandRequests <-chan chan string) <-chan bool {
	res := make(chan bool)
	go func() {
		good := true
		for {
			actionsCh := <-requests
			if actionsCh == nil {
				defer func() { res <- good }()
				return
			}
			commandsCh := <-commandRequests
			if !bytes.Equal(<-actionsCh, payload) {
				good = false
			}
			commandsCh <- "not_interested"
		}
	}()
	return res
}

func startMockExec(requests <-chan chan []byte, commandRequests <-chan chan string) {
	go func() {
		for {
			actionsCh := <-requests
			commandsCh := <-commandRequests
			go func() {
				<-actionsCh
				commandsCh <- "interested"
				switch <-commandsCh {
				case "can_commit?":
					commandsCh <- "prepared"
				case "do_abort":
					commandsCh <- "done"
					return
				default:
					panic(errors.New("illegal command"))
				}
				switch <-commandsCh {
				case "do_commit", "do_abort":
				default:
					panic(errors.New("illegal command"))
				}
				commandsCh <- "done"
			}()
		}
	}()
}
