package communication

import (
	"fmt"
	"steel-lang/datastructure"
	"testing"

	"github.com/google/uuid"
)

func TestMakeMemberlistAgent(t *testing.T) {
	tests := []struct {
		index int
		set   string
		port  int
		nodes []string
	}{
		//  {_, set, port, nodes},
		{1, "lorem,ipsum10", 0, nil},
		{2, "b", 8100, []string{}},
		{3, "D__987,a,b543_", 8101, []string{"127.0.0.1:8150"}},
		{4, "E_e1,E_e1,C____,d_210_", 8102, []string{"127.0.0.1:8150,127.0.0.1:8151"}},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("TestMakeMemberlistAgent#%d", test.index), func(t *testing.T) {
			agt := MakeMemberlistAgent(datastructure.MakeStringSet(test.set), test.port, test.nodes)
			if agt.IsRunning() {
				t.Error("agent should not be running")
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
			if agt.listPtr != &agt.list {
				t.Error("listPtr should be pointing to list")
			}
			checkCorrectStop(t, agt)
		})
	}
}

func TestStart(t *testing.T) {
	tests := []struct {
		index int
		set   string
		port  int
		nodes []string
	}{
		//  {_, set, port, nodes},
		{1, "enim,minim10", 0, nil},
		{2, "dolor123,sit", 9100, []string{}},
		{3, "a,b,c,d,e,f,f,f", 9101, []string{"127.0.0.1:9150"}},
		{4, "iIi,H_h_h_,G_qwerty_", 9102, []string{"127.0.0.1:9150,127.0.0.1:9151"}},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("TestStart#%d", test.index), func(t *testing.T) {
			agt := MakeMemberlistAgent(datastructure.MakeStringSet(test.set), test.port, test.nodes)
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
		{9, 10101, []string{"127.0.0.1:10150"}, false, false},
		{10, 10103, []string{"127.0.0.1:10101", "127.0.0.1:10102"}, false, false},
		{11, 10105, []string{".,>Z><<-@#00asdfg"}, false, false},
	}
	dummy := make([]*memberlistAgent, 0, len(tests))
	resources := datastructure.MakeStringSet("laboris,nisi,ut,aliquip,ex")
	for i, test := range tests {
		t.Run(fmt.Sprintf("TestJoin#%d", test.index), func(t *testing.T) {
			dummy = append(dummy, MakeMemberlistAgent(resources, test.port, test.nodes))
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
	a := MakeMemberlistAgent(datastructure.MakeStringSet("a______789,B___,C_1_qwerty"), port, nil)
	checkCorrectStop(t, a)
	err := a.ForAll([]byte(`lorem`))
	if err == nil {
		t.Error("ForAll should return error when agent is not running")
	}
	start(t, a, port)
	statusCh := localForAll(t, a, `incididunt ut labore et`, false)
	a.operations <- nil
	if !<-statusCh {
		t.Error("received wrong payload")
	}
	statusCh = localForAll(t, a, `dolore magna aliqua. Ut`, true)
	a.operations <- nil
	if !<-statusCh {
		t.Error("received wrong payload")
	}
}

func TestStop(t *testing.T) {
	const port = 11100
	a := MakeMemberlistAgent(datastructure.MakeStringSet("amet_456"), port, nil)
	if a.Stop() == nil {
		t.Error("should return error when agent is not running")
	}
	start(t, a, port)
	uuid1 := a.list.LocalNode().Name
	startMockInterested("", a.operations, a.operationCommands)
	restart(t, a, port)
	uuid2 := a.list.LocalNode().Name
	if uuid1 == uuid2 {
		t.Error("uuid should be different after restart")
	}
	localForAll(t, a, `consectetur adipiscing elit`, false)
	restart(t, a, port)
	uuid3 := a.list.LocalNode().Name
	if uuid1 == uuid3 || uuid2 == uuid3 {
		t.Error("uuid should be different after restart")
	}
	localForAll(t, a, `, sed do eiusmod tempor`, true)
	stop(t, a)
}

func start(t *testing.T, a *memberlistAgent, p int) {
	t.Helper()
	err := a.Start()
	if err != nil {
		t.Fatal(err.Error())
	}
	checkCorrectStart(t, a, p)
}

func stop(t *testing.T, a *memberlistAgent) {
	t.Helper()
	err := a.Stop()
	if err != nil {
		t.Fatal(err.Error())
	}
	checkCorrectStop(t, a)
}

func restart(t *testing.T, a *memberlistAgent, p int) {
	t.Helper()
	stop(t, a)
	start(t, a, p)
}

func localForAll(t *testing.T, a *memberlistAgent, payload string, interested bool) <-chan bool {
	t.Helper()
	ops, cmds := a.ReceivedActions()
	var res <-chan bool
	if interested {
		res = startMockInterested(payload, ops, cmds)
	} else {
		res = startMockUninterested(payload, ops, cmds)
	}
	err := a.ForAll([]byte(payload))
	if err != nil {
		t.Fatal(err.Error())
	}
	return res
}

func checkCorrectStart(t *testing.T, a *memberlistAgent, p int) {
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
	switch {
	case a.registry == nil:
		t.Error("registry should not be nil")
	case a.terminated == nil:
		t.Error("terminated should not be nil")
	case a.waitingForRegistry == nil:
		t.Error("waitingForRegistry should not be nil")
	case a.haltUpdates == nil:
		t.Error("haltUpdates should not be nil")
	case a.quitTransactions == nil:
		t.Error("quitTransactions should not be nil")
	case a.quitGossip == nil:
		t.Error("quitGossip should not be nil")
	case a.quitDemux == nil:
		t.Error("quitDemux should not be nil")
	case a.pendingUpdates == nil:
		t.Error("pendingUpdates should not be nil")
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

func checkCorrectStop(t *testing.T, a *memberlistAgent) {
	t.Helper()
	if a.IsRunning() {
		t.Error("should not be running")
	}
	switch {
	case a.registry != nil:
		t.Error("registry should be nil")
	case a.terminated != nil:
		t.Error("terminated should be nil")
	case a.list != nil:
		t.Error("list should be nil")
	case a.config != nil:
		t.Error("config should be nil")
	case a.waitingForRegistry != nil:
		t.Error("waitingForRegistry should be nil")
	case a.haltUpdates != nil:
		t.Error("haltUpdates should be nil")
	case a.quitTransactions != nil:
		t.Error("quitTransactions should be nil")
	case a.quitGossip != nil:
		t.Error("quitGossip should be nil")
	case a.quitDemux != nil:
		t.Error("quitDemux should be nil")
	case a.pendingUpdates != nil:
		t.Error("pendingUpdates should be nil")
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

func startMockInterested(payload string, requests <-chan chan []byte, commandRequests <-chan chan string) <-chan bool {
	res := make(chan bool)
	go func(status chan<- bool, payload string, requests <-chan chan []byte, commandRequests <-chan chan string) {
		good := true
		for {
			actionsCh := <-requests
			if actionsCh == nil {
				defer func() { status <- good }()
				return
			}
			commandsCh := <-commandRequests
			if payload != string(<-actionsCh) {
				good = false
			}
			commandsCh <- "interested"
			<-commandsCh
			commandsCh <- "done"
		}
	}(res, payload, requests, commandRequests)
	return res
}

func startMockUninterested(payload string, requests <-chan chan []byte, commandRequests <-chan chan string) <-chan bool {
	res := make(chan bool)
	go func(status chan<- bool, payload string, requests <-chan chan []byte, commandRequests <-chan chan string) {
		good := true
		for {
			actionsCh := <-requests
			if actionsCh == nil {
				defer func() { status <- good }()
				return
			}
			commandsCh := <-commandRequests
			if payload != string(<-actionsCh) {
				good = false
			}
			commandsCh <- "not_interested"
		}
	}(res, payload, requests, commandRequests)
	return res
}
