package communication

import (
	"errors"
	"fmt"
	"steel-lang/config"
	"sync"

	"github.com/google/uuid"
	"github.com/hashicorp/memberlist"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	// milliseconds
	timeoutRegister = 1000
	msgBuffLen      = 10
)

const (
	TestsNothing = iota
	TestsAbort
	TestsUnreliableSend
	// Halting
	TestsMidInterested
	TestsAfterInterested
	TestsMidFirst
	TestsAfterFirst
	TestsMidSecond
)

var TestsMidSends = 2

type messageUnion struct {
	Type   string
	Sender *memberlist.Node

	Transaction transactionInfo
}

type memberlistAgent struct {
	initialNodes []string
	terminated   map[string]string
	transactions map[string]*transactionInfo
	logLevel     zap.AtomicLevel
	logger       *zap.Logger

	running               bool
	initialConfig         *memberlist.Config
	config                *memberlist.Config
	list                  *memberlist.Memberlist
	delegate              MemberlistDelegate
	adapter               delegateAdapter
	quitTransactions      chan chan bool
	quitGossip            chan chan bool
	quitDemux             chan chan bool
	transactionMessages   chan messageUnion
	transactionResponses  chan messageUnion
	coordinatedChannels   chan chan transactionChannels
	trackGossip           chan chan *sync.WaitGroup
	initiatedTransactions int

	listeningPort     int
	operations        chan chan []byte
	operationCommands chan chan string
	// testing
	test       int
	halted     bool
	lockHalted *sync.Mutex
}

func MakeMemberlistAgent(port int, lc config.LogConfig, nodes ...string) *memberlistAgent {
	return MakeMemberlistAgentAdvanced(port, nil, nil, lc, nodes...)
}

func MakeMemberlistAgentAdvanced(port int, cfg *memberlist.Config, delegate *MemberlistDelegate, lc config.LogConfig, nodes ...string) *memberlistAgent {
	res := &memberlistAgent{
		running:               false,
		listeningPort:         port,
		config:                &memberlist.Config{},
		initialNodes:          nodes,
		initiatedTransactions: 0,
		operations:            make(chan chan []byte),
		operationCommands:     make(chan chan string),
	}
	if cfg != nil {
		res.initialConfig = cfg
	} else {
		res.initialConfig = memberlist.DefaultLocalConfig()
	}
	if delegate != nil {
		res.delegate = *delegate
	} else {
		res.delegate = delegateDefault{}
	}
	if lc.Encoding == "" {
		lc.Encoding = "console"
	}
	zapCfg, ok := config.LogPreset(lc.Encoding).(zap.Config)
	if !ok {
		zapCfg = zap.NewProductionConfig()
	}
	res.logLevel = zapCfg.Level
	var err error
	res.logger, err = zapCfg.Build()
	if err != nil {
		if ok { // fallback to zap default
			ok = false
			zapCfg = zap.NewProductionConfig()
			res.logLevel = zapCfg.Level
			res.logger, err = zapCfg.Build()
		}
		if err != nil {
			panic("could not create memberlistAgent logger")
		}
	}
	res.SetLogLevel(lc.Level)
	if !ok {
		res.logger.Warn("could not load memberlistAgent logger config")
	}
	return res
}

func (a *memberlistAgent) IsRunning() bool {
	return a.running
}

func (a *memberlistAgent) Start() error {
	if a.running {
		return errors.New("agent is already running")
	}
	uuid, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	a.terminated = make(map[string]string)
	a.transactions = make(map[string]*transactionInfo)
	a.quitTransactions = make(chan chan bool)
	a.quitGossip = make(chan chan bool)
	a.quitDemux = make(chan chan bool)
	a.transactionMessages = make(chan messageUnion, msgBuffLen)
	a.transactionResponses = make(chan messageUnion, msgBuffLen)
	a.coordinatedChannels = make(chan chan transactionChannels)
	a.trackGossip = make(chan chan *sync.WaitGroup)

	*a.config = *a.initialConfig
	stdLog, err := zap.NewStdLogAt(a.logger, zapcore.DebugLevel)
	if err != nil {
		return err
	}
	a.config.Logger = stdLog
	a.config.BindPort = a.listeningPort
	a.config.Name = uuid.String()

	a.adapter = a.makeAdapter(a.delegate)
	a.config.Delegate = a.adapter
	a.config.Events = a.adapter

	go joiner(a.trackGossip, a.quitGossip)

	// start listening
	a.list, err = memberlist.Create(a.config)
	if err != nil {
		replyCh := make(chan bool)
		a.quitGossip <- replyCh
		<-replyCh
		return err
	}

	a.running = true
	a.adapter.start()
	go demuxResponses(a.coordinatedChannels, a.transactionResponses, a.quitDemux)
	go a.handleTransactions()
	return nil
}

func (a *memberlistAgent) Join() error {
	if !a.running {
		return errors.New("agent is not running")
	}
	if len(a.initialNodes) > 0 {
		_, err := a.list.Join(a.initialNodes)
		return err
	}
	return nil
}

func (a *memberlistAgent) ForAll(payload []byte) error {
	if !a.running {
		return errors.New("agent is not running")
	}
	info := transactionInfo{
		Initiator: a.list.LocalNode().Name,
		Number:    a.initiatedTransactions,
		Payload:   payload,
	}
	a.initiatedTransactions++
	info.Partecipants = a.interested(info)
	if len(info.Partecipants) == 0 {
		fmt.Println("terminated transaction: none interested")
		return nil
	}
	return a.coordinateTransaction(info)
}

func (a *memberlistAgent) ReceivedActions() (<-chan chan []byte, <-chan chan string) {
	return a.operations, a.operationCommands
}

func (a *memberlistAgent) Stop() error {
	if !a.running {
		return errors.New("agent is not running")
	}

	fmt.Println("Stopping transaction handling...")
	replyCh := make(chan bool)
	a.quitTransactions <- replyCh
	<-replyCh
	a.operations <- nil
	fmt.Println("Stopped transaction handling")
	fmt.Println("Stopping response demultiplexing...")
	replyCh = make(chan bool)
	a.quitDemux <- replyCh
	<-replyCh
	fmt.Println("Stopped response demultiplexing")
	fmt.Println("Stopping update handling...")
	a.adapter.stop()
	fmt.Println("Stopped update handling")
	fmt.Println("Gossiping leave...")
	err := a.list.Leave(a.config.PushPullInterval)
	if err != nil {
		fmt.Println("error in gossiping leave:", err.Error())
	} else {
		fmt.Println("Gossiped leave")
	}
	fmt.Println("Detaching from group...")
	err = a.list.Shutdown() // v0.2.4 always returns nil
	if err != nil {
		fmt.Println("error in leaving group:", err.Error())
	} else {
		fmt.Println("Left group")
	}
	fmt.Println("Stopping gossip handling...")
	replyCh = make(chan bool)
	a.quitGossip <- replyCh
	<-replyCh
	fmt.Println("Stopped gossip handling")

	// preserve delegate
	a.delegate = a.adapter.delegate

	// clean up
	a.terminated = nil
	a.transactions = nil
	a.list = nil
	a.config = &memberlist.Config{}
	a.adapter = delegateAdapter{}
	a.transactionMessages = nil
	a.quitTransactions = nil
	a.quitGossip = nil
	a.quitDemux = nil
	a.transactionResponses = nil
	a.coordinatedChannels = nil
	a.trackGossip = nil
	a.running = false
	return nil
}

func (a *memberlistAgent) SetLogLevel(l int) {
	if l < config.LogDebug {
		l = config.LogDebug
	} else if l > config.LogFatal {
		l = config.LogFatal
	}
	zapLevel := zapcore.InfoLevel
	switch l {
	case config.LogDebug:
		zapLevel = zapcore.DebugLevel
	case config.LogWarning:
		zapLevel = zapcore.WarnLevel
	case config.LogError:
		zapLevel = zapcore.ErrorLevel
	case config.LogFatal:
		zapLevel = zapcore.DPanicLevel
	}
	a.logLevel.SetLevel(zapLevel)
}

//---------------------------------DELEGATES----------------------------------

func joiner(track <-chan chan *sync.WaitGroup, quit <-chan chan bool) {
	var waitGroup sync.WaitGroup
	for {
		select {
		case t := <-track:
			waitGroup.Add(1)
			t <- &waitGroup
		case q := <-quit:
			defer func() { q <- true }()
			waitGroup.Wait()
			return
		}
	}
}

func (a *memberlistAgent) makeAdapter(d MemberlistDelegate) delegateAdapter {
	return delegateAdapter{
		listPtr:              &a.list,
		trackGossip:          a.trackGossip,
		transactionMessages:  a.transactionMessages,
		transactionResponses: a.transactionResponses,
		delegate:             d,
		members: BaseMembers{
			Config:          a.config,
			Terminated:      a.terminated,
			Transactions:    a.transactions,
			ListeningPort:   a.listeningPort,
			Logger:          a.logger,
			ReceivedActions: a.ReceivedActions,
		},
	}
}

//----------------------------------TESTING-----------------------------------

func TestsMakeMemberlistAgent(port int, test int, nodes ...string) *memberlistAgent {
	res := MakeMemberlistAgent(port, config.TestsLogConfig, nodes...)
	res.test = test
	res.lockHalted = &sync.Mutex{}
	return res
}

func (a *memberlistAgent) TestsHaltReturn() {
	a.list.Shutdown()
	a.lockHalted.Lock()
	a.halted = true
	a.lockHalted.Unlock()
}

func (a *memberlistAgent) testsHalt() {
	a.TestsHaltReturn()
	var block chan bool = nil
	<-block
}

func (a *memberlistAgent) testsHaltIf(t int) {
	if a.test == t {
		a.testsHalt()
	}
}
