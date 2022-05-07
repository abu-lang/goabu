// Package communication provides means for transactional communication.
package communication

import (
	"encoding/json"
	"errors"
	"sync"

	"github.com/abu-lang/goabu/config"

	"github.com/google/uuid"
	"github.com/hashicorp/memberlist"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// TODO evaluate
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

func (m *messageUnion) marshal(obj string, logger *zap.Logger) ([]byte, bool) {
	res, err := json.Marshal(*m)
	if err != nil {
		logger.Error("Error during marshalling: "+err.Error(),
			zap.String("act", "marshalling"),
			zap.String("obj", obj))
		return nil, false
	}
	return res, true
}

func (m *messageUnion) unmarshal(bs []byte) bool {
	err := json.Unmarshal(bs, m)
	return err == nil
}

type MemberlistAgent struct {
	id           string
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

// NewMemberlistAgent creates a stopped MemberlistAgent which implements the goabu.Agent interface.
//
// id must be a string uniquely identifying the new MemberlistAgent or "", if id == "" then a random
// identifier is arbitrarily chosen in its place.
//
// port specifies on which port should the created MemberlistAgent listen for join request coming from
// other MemberlistAgents.
//
// nodes is a list of strings of the form "host:port" and indicates to whom the created MemberlistAgent
// should send join request when the Join method is called.
func NewMemberlistAgent(id string, port int, lc config.LogConfig, nodes ...string) *MemberlistAgent {
	return NewMemberlistAgentAdvanced(id, port, nil, nil, lc, nodes...)
}

// NewMemberlistAgentAdvanced creates a stopped MemberlistAgent which implements the goabu.Agent interface.
// It is a more verbose and configurable version of NewMemberlistAgent.
//
// cfg specifies the configuration of the underlying memberlist.Memberlist.
//
// delegate consents to override the handling of the memberlist.Memberlist events, see file delegateDefault.go
// for the default implementation.
func NewMemberlistAgentAdvanced(id string, port int, cfg *memberlist.Config, delegate *MemberlistDelegate,
	lc config.LogConfig, nodes ...string) *MemberlistAgent {
	res := &MemberlistAgent{
		id:                    id,
		running:               false,
		listeningPort:         port,
		config:                &memberlist.Config{},
		initialNodes:          nodes,
		initiatedTransactions: 0,
		operations:            make(chan chan []byte),
		operationCommands:     make(chan chan string),
	}
	if res.id == "" {
		res.id = uuid.New().String() + "/agent"
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
			panic("could not create MemberlistAgent logger")
		}
	}
	res.SetLogLevel(lc.Level)
	if !ok {
		res.logger.Warn("Could not load MemberlistAgent logger config",
			zap.String("act", "load"),
			zap.String("obj", "agent logger config"))
	}
	return res
}

func (a *MemberlistAgent) IsRunning() bool {
	return a.running
}

func (a *MemberlistAgent) Start() error {
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
	go demuxResponses(a.coordinatedChannels, a.transactionResponses, a.quitDemux, a.logger)
	go a.handleTransactions()
	return nil
}

func (a *MemberlistAgent) Join() error {
	if !a.running {
		return errors.New("agent is not running")
	}
	if len(a.initialNodes) > 0 {
		_, err := a.list.Join(a.initialNodes)
		return err
	}
	return nil
}

func (a *MemberlistAgent) ForAll(payload []byte) error {
	if !a.running {
		return errors.New("agent is not running")
	}
	info := transactionInfo{
		Initiator: a.list.LocalNode().Name,
		Number:    a.initiatedTransactions,
		Payload:   payload,
	}
	a.initiatedTransactions++
	var err error
	info.Participants, err = a.interested(info)
	if err != nil {
		return err
	}
	if len(info.Participants) == 0 {
		a.logger.Debug("Terminated transaction: none interested", zap.String("act", "end_tran"), zap.Int("participants", 0))
		return nil
	}
	return a.coordinateTransaction(info)
}

func (a *MemberlistAgent) ReceivedActions() (<-chan chan []byte, <-chan chan string) {
	return a.operations, a.operationCommands
}

func (a *MemberlistAgent) Stop() error {
	if !a.running {
		return errors.New("agent is not running")
	}

	a.logStopping("transaction handling")
	replyCh := make(chan bool)
	a.quitTransactions <- replyCh
	<-replyCh
	a.operations <- nil
	a.logStopped("transaction handling")

	a.logStopping("response demultiplexing")
	replyCh = make(chan bool)
	a.quitDemux <- replyCh
	<-replyCh
	a.logStopped("response demultiplexing")

	a.logStopping("delegate")
	a.adapter.stop()
	a.logStopped("delegate")

	a.logger.Debug("Gossiping leave...", zap.String("act", "gossip"), zap.String("obj", "leave"))
	err := a.list.Leave(a.config.PushPullInterval)
	if err != nil {
		a.logger.Warn("Error in gossiping leave: "+err.Error(), zap.String("act", "gossip"), zap.String("obj", "leave"))
	} else {
		a.logger.Debug("Gossiped leave", zap.String("act", "gossip"), zap.String("obj", "leave"))
	}

	a.logger.Debug("Leaving group...", zap.String("act", "leave"))
	err = a.list.Shutdown() // always returns nil in memberlist v0.3.1
	if err != nil {
		a.logger.Error("Error in leaving group: "+err.Error(), zap.String("act", "leave"))
	} else {
		a.logger.Debug("Left group", zap.String("act", "leave"))
	}

	a.logStopping("gossip handling")
	replyCh = make(chan bool)
	a.quitGossip <- replyCh
	<-replyCh
	a.logStopped("gossip handling")
	a.logger.Info("Stopped agent", zap.String("act", "stop"), zap.String("subj", a.id))
	a.logger.Sync()

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

func (a *MemberlistAgent) SetLogLevel(l int) {
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

func (a *MemberlistAgent) logStopping(proc string) {
	a.logger.Debug("Stopping "+proc+"...",
		zap.String("act", "stop"),
		zap.String("obj", proc))
}

func (a *MemberlistAgent) logStopped(proc string) {
	a.logger.Debug("Stopped "+proc,
		zap.String("act", "stop"),
		zap.String("obj", proc))
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

func (a *MemberlistAgent) makeAdapter(d MemberlistDelegate) delegateAdapter {
	return delegateAdapter{
		listPtr:              &a.list,
		trackGossip:          a.trackGossip,
		transactionMessages:  a.transactionMessages,
		transactionResponses: a.transactionResponses,
		delegate:             d,
		members: BaseMembers{
			AgentID:         a.id,
			ListeningPort:   a.listeningPort,
			List:            nil,
			Config:          a.config,
			ReceivedActions: a.ReceivedActions,
			Logger:          a.logger,
		},
	}
}

//----------------------------------TESTING-----------------------------------

// TestsNewMemberlistAgent is used for testing purposes.
//
// It behaves like NewMemberlistAgent when config.TestsLogConfig is passed as the lc
// argument with the difference that the returned MemberlistAgent simulates a crash failure
// upon the happening of a particular event specified through the test argument.
func TestsNewMemberlistAgent(id string, port int, test int, nodes ...string) *MemberlistAgent {
	res := NewMemberlistAgent(id, port, config.TestsLogConfig, nodes...)
	res.test = test
	res.lockHalted = &sync.Mutex{}
	return res
}

func (a *MemberlistAgent) TestsHalt() {
	a.list.Shutdown()
	a.lockHalted.Lock()
	a.halted = true
	a.lockHalted.Unlock()
}

func (a *MemberlistAgent) testsHaltAndBlock() {
	a.TestsHalt()
	var block chan bool = nil
	<-block
}

func (a *MemberlistAgent) testsHaltIf(t int) {
	if a.test == t {
		a.testsHaltAndBlock()
	}
}
