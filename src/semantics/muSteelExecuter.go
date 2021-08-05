package semantics

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"steel-lang/datastructure"
	"steel-lang/misc"
	"steel-lang/parser"
	antlr_parser "steel-lang/parser/antlr"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
)

const inputsRate float64 = 5.0

// milliseconds
const inputsFlush = 100

type State struct {
	Memory datastructure.Resources
	Pool   [][]SemanticAction
}

type MuSteelExecuter struct {
	memory        datastructure.ResourceController
	types         map[string]string
	pool          [][]SemanticAction
	lockPool      sync.Mutex
	localLibrary  map[string]datastructure.RuleDict
	globalLibrary map[string]datastructure.RuleDict

	workingMemory *ast.WorkingMemory
	dataContext   ast.IDataContext

	agent ISteelAgent
}

func NewMuSteelExecuter(mem datastructure.ResourceController, rules []string, agt ISteelAgent) (*MuSteelExecuter, error) {
	res := &MuSteelExecuter{
		memory:        mem.Clone(),
		pool:          make([][]SemanticAction, 0),
		localLibrary:  make(map[string]datastructure.RuleDict),
		globalLibrary: make(map[string]datastructure.RuleDict),
		agent:         agt,
	}
	if !res.memory.IsValid() {
		return nil, errors.New("invalid Resources argument")
	}
	res.types = res.memory.GetTypes()
	var err error
	res.dataContext, res.workingMemory, err = res.NewEmptyGruleStructures("this")
	if err != nil {
		return nil, err
	}
	err = res.AddRules(rules)
	if err != nil {
		return nil, err
	}
	err = mem.Start()
	if err != nil {
		return nil, err
	}
	go res.receiveInputs()
	err = res.StartAgent()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (m *MuSteelExecuter) StartAgent() error {
	err := m.agent.Start()
	if err != nil {
		return err
	}
	go m.receiveExternalActions()
	err = m.agent.Join()
	if err != nil {
		return err
	}
	return nil
}

func (m *MuSteelExecuter) StopAgent() error {
	return m.agent.Stop()
}

func (m *MuSteelExecuter) SetAgent(agt ISteelAgent) error {
	if m.agent.IsRunning() {
		return errors.New("agent is still running")
	}
	m.agent = agt
	return nil
}

func (m *MuSteelExecuter) GetState() State {
	memCopy := m.memory.Clone().GetResources()
	m.lockPool.Lock()
	poolCopy := make([][]SemanticAction, 0, len(m.pool))
	for _, acts := range m.pool {
		actsCopy := make([]SemanticAction, len(acts))
		copy(actsCopy, acts)
		poolCopy = append(poolCopy, actsCopy)
	}
	m.lockPool.Unlock()
	return State{Memory: memCopy, Pool: poolCopy}
}

func (m *MuSteelExecuter) IsStable() bool {
	m.lockPool.Lock()
	defer m.lockPool.Unlock()
	return len(m.pool) == 0
}

func (m *MuSteelExecuter) HasRule(name string) bool {
	for _, d := range m.localLibrary {
		if d.Contains(name) {
			return true
		}
	}
	for _, d := range m.globalLibrary {
		if d.Contains(name) {
			return true
		}
	}
	return false
}

func (m *MuSteelExecuter) AddRule(r string) error {
	rule, err := m.parseRule(r)
	if err != nil {
		return err
	}
	if m.HasRule(rule.Name) {
		return fmt.Errorf("there is already a rule named %s", rule.Name)
	}

	library := m.localLibrary
	if rule.Task.Mode != "for" {
		library = m.globalLibrary
	}
	for _, evt := range rule.Events {
		if library[evt] == nil {
			var dict datastructure.RuleDict = datastructure.MakeRuleDict()
			library[evt] = dict
		}
		library[evt].Insert(rule)
	}
	return nil
}

func (m *MuSteelExecuter) AddRules(rules []string) error {
	return addList(rules, m.AddRule)
}

func (m *MuSteelExecuter) Exec() {
	m.lockPool.Lock()
	if len(m.pool) == 0 {
		m.lockPool.Unlock()
		return
	}
	actions, index := m.chooseActions()
	m.removeActions(index)
	m.lockPool.Unlock()
	fmt.Print("Exec: ")
	m.execActions(actions)
}

func (m *MuSteelExecuter) Input(actions string) error {
	parsed, err := m.parseActions(actions)
	if err != nil {
		return err
	}
	sActions := evalActions(parsed, m.dataContext, m.workingMemory)
	fmt.Print("Input: ")
	m.execActions(sActions)
	return nil
}

func (m *MuSteelExecuter) receiveInputs() {
	inputs := m.memory.Inputs()
	queueSize := int(math.RoundToEven(float64(m.memory.InputsNumber()) * inputsRate))
	var queue string = ""
	var l int = 0
	var timeout <-chan time.Time = nil
	var queued misc.StringSet = misc.MakeStringSet("")
	for {
		select {
		case action := <-inputs:
			resource := strings.TrimSpace(strings.Split(action, "=")[0])
			if queued.Contains(resource) {
				err := m.Input(queue)
				if err != nil {
					panic(err)
				}
				queue = ""
				l = 0
				timeout = nil
				queued = misc.MakeStringSet("")
			}
			queue += action
			l++
			queued.Insert(resource)
			if l == 1 {
				timeout = time.After(inputsFlush * time.Millisecond)
			}
			if l < queueSize {
				continue
			}
		case <-timeout:
		}
		err := m.Input(queue)
		if err != nil {
			panic(err)
		}
		queue = ""
		l = 0
		timeout = nil
		queued = misc.MakeStringSet("")
	}
}

func (m *MuSteelExecuter) receiveExternalActions() {
	requests, commandRequests := m.agent.ReceivedActions()
	for {
		actionsCh := <-requests
		if actionsCh == nil {
			return
		}
		commandsCh := <-commandRequests
		eActions, err := unmarshalExternalActions(<-actionsCh, m.types)
		if err != nil {
			panic(err)
		}
		var sActions [][]SemanticAction
		m.lockPool.Lock()
		localResources := m.memory.ResourceNames()
		context, workMem, err := m.NewEmptyGruleStructures("ext")
		if err != nil {
			panic(err)
		}
		for _, eAction := range eActions {
			if localResources.ContainsSet(eAction.CondWorkingSet) {
				actions := eAction.cullActions(localResources)
				if len(actions) == 0 {
					continue
				}
				sActions = appendNonempty(sActions, condEvalActions(eAction.Condition, actions, context, workMem))
			}
		}
		if len(sActions) == 0 {
			commandsCh <- "not_interested"
			m.lockPool.Unlock()
			continue
		}
		commandsCh <- "interested"
		switch <-commandsCh {
		case "do_commit":
			m.pool = append(m.pool, sActions...)
			fallthrough
		case "do_abort":
			commandsCh <- "done"
		}
		m.lockPool.Unlock()
	}
}

func (m *MuSteelExecuter) addActions(actions string) error {
	parsed, err := m.parseActions(actions)
	if err != nil {
		return err
	}
	m.lockPool.Lock()
	m.pool = append(m.pool, evalActions(parsed, m.dataContext, m.workingMemory))
	m.lockPool.Unlock()
	return nil
}

func (m *MuSteelExecuter) addPool(pl []string) error {
	return addList(pl, m.addActions)
}

func (m *MuSteelExecuter) chooseActions() ([]SemanticAction, int) {
	// TODO: implement other strategies
	return m.pool[0], 0
}

func (m *MuSteelExecuter) execActions(actions []SemanticAction) {
	m.lockPool.Lock()
	var Xset []SemanticAction
	for _, action := range actions {
		variable := action.Variable
		variable = m.workingMemory.AddVariable(variable)
		currentVal, err := variable.Evaluate(m.dataContext, m.workingMemory)
		if err != nil {
			panic(err)
		}
		diff := false
		if currentVal.Kind() == reflect.Interface || action.Value.Kind() == reflect.Interface {
			diff = true
		} else {
			eq, err := pkg.EvaluateEqual(currentVal, action.Value)
			if err != nil {
				panic(err)
			}
			if !eq.Bool() {
				diff = true
				ltype := currentVal.Type()
				rtype := action.Value.Type()
				if !rtype.AssignableTo(ltype) {
					panic(fmt.Errorf("cannot assign a %v to a %v", rtype, ltype))
				}
			}
		}
		if diff {
			err := variable.Assign(action.Value, m.dataContext, m.workingMemory)
			if err != nil {
				panic(err)
			}
			m.memory.Modified(action.Resource)
			Xset = append(Xset, action)
			fmt.Print(action)
		}
	}
	fmt.Println()
	sActions, eActions := m.discovery(Xset)
	m.pool = append(m.pool, sActions...)
	m.lockPool.Unlock()
	if len(eActions) > 0 {
		payload, err := marshalExternalActions(eActions)
		if err == nil {
			err = m.agent.ForAll(payload)
		}
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func (m *MuSteelExecuter) removeActions(index int) {
	m.pool = append(m.pool[:index], m.pool[index+1:len(m.pool)]...)
}

func (m *MuSteelExecuter) discovery(Xset []SemanticAction) ([][]SemanticAction, []externalAction) {
	var newpool [][]SemanticAction
	var extActions []externalAction
	localRules, globalRules := m.activeRules(Xset)
	for _, rule := range localRules {
		if len(rule.DefaultActions) > 0 {
			newpool = append(newpool, evalActions(rule.DefaultActions, m.dataContext, m.workingMemory))
		}
		newpool = appendNonempty(newpool, condEvalActions(rule.Task.Condition, rule.Task.Actions, m.dataContext, m.workingMemory))
	}
	for _, rule := range globalRules {
		if len(rule.DefaultActions) > 0 {
			newpool = append(newpool, evalActions(rule.DefaultActions, m.dataContext, m.workingMemory))
		}
		ext := m.preEvaluated(rule)
		extActions = append(extActions, ext)
	}
	return newpool, extActions
}

func (m *MuSteelExecuter) activeRules(Xset []SemanticAction) (local, global datastructure.RuleDict) {
	local = datastructure.MakeRuleDict()
	global = datastructure.MakeRuleDict()
	for _, act := range Xset {
		local.Add(m.localLibrary[act.Resource])
		global.Add(m.globalLibrary[act.Resource])
	}
	return local, global
}

// Precondition: rule.Task.Mode != "for"
func (m *MuSteelExecuter) preEvaluated(rule *datastructure.Rule) externalAction {
	res := externalAction{
		CondWorkingSet: misc.MakeStringSet(""),
		Constants:      make(map[string]interface{}),
		IntConstants:   make(map[string]int64),
		dataContext:    m.dataContext,
		workingMemory:  m.workingMemory,
	}
	res.WorkingSets = make([]misc.StringSet, 0, len(rule.Task.Actions))
	for _, action := range rule.Task.Actions {
		res.WorkingSets = append(res.WorkingSets, misc.MakeStringSet(action.Resource))
	}
	res.Condition = res.preEvaluatedExpression(rule.Task.Condition, res.CondWorkingSet)
	res.Actions = res.preEvaluatedActions(rule.Task.Actions)
	return res
}

func (m *MuSteelExecuter) parseRule(r string) (*datastructure.Rule, error) {
	var err error
	listener := parser.NewEcaruleParserListener(m.types, m.workingMemory, func(e error) {
		err = e
	})

	ts := parser.TokenStream(r)
	p := antlr_parser.NewEcaruleParser(ts)
	p.BuildParseTrees = true
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Prule())

	// update WorkingMemory
	m.workingMemory.IndexVariables()

	if err != nil {
		return nil, err
	}
	return listener.Rule, nil
}

func (m *MuSteelExecuter) parseActions(actions string) ([]datastructure.Action, error) {
	var err error
	listener := parser.NewEcaruleParserListener(m.types, m.workingMemory, func(e error) {
		err = e
	})

	ts := parser.TokenStream(actions)
	p := antlr_parser.NewEcaruleParser(ts)
	p.BuildParseTrees = true

	antlr.ParseTreeWalkerDefault.Walk(listener, p.Actions())

	// update WorkingMemory
	m.workingMemory.IndexVariables()

	if err != nil {
		return nil, err
	}
	return listener.Rule.DefaultActions, nil
}

func (m *MuSteelExecuter) NewEmptyGruleStructures(name string) (ast.IDataContext, *ast.WorkingMemory, error) {
	dataContext := ast.NewDataContext()
	resources := m.memory.GetResources()
	err := dataContext.Add(name, &(resources))
	if err != nil {
		return dataContext, nil, err
	}
	kbName := "dummy_" + name
	version := "0.0.0"
	knowledgeBase := &ast.KnowledgeBase{
		Name:          kbName,
		Version:       version,
		RuleEntries:   make(map[string]*ast.RuleEntry),
		WorkingMemory: ast.NewWorkingMemory(kbName, version),
	}
	defunc := &ast.BuiltInFunctions{
		Knowledge:     knowledgeBase,
		WorkingMemory: knowledgeBase.WorkingMemory,
		DataContext:   dataContext,
	}
	err = dataContext.Add("DEFUNC", defunc)
	if err != nil {
		return dataContext, nil, err
	}
	knowledgeBase.InitializeContext(dataContext)
	return dataContext, knowledgeBase.WorkingMemory, nil
}

func (m *MuSteelExecuter) PrintState() string {
	return fmt.Sprintf("Memory: %v\nPool: %v\n", m.memory, m.printPool())
}

func (m *MuSteelExecuter) printPool() string {
	m.lockPool.Lock()
	defer m.lockPool.Unlock()
	if len(m.pool) == 0 {
		return "{}"
	} else {
		str := "{"
		for _, list := range m.pool {
			str = str + "\n  "
			for _, action := range list {
				str = str + action.String()
			}
		}
		return str + "\n}"
	}
}

func addList(strs []string, add func(string) error) error {
	var fstErr error
	failed := ""
	for i, s := range strs {
		err := add(s)
		if err != nil {
			failed += strconv.Itoa(i) + ", "
			if fstErr == nil {
				fstErr = err
			}
		}
	}
	if fstErr != nil {
		return fmt.Errorf("could not add elements with indexes %s as %s", failed[:len(failed)-2], fstErr.Error())
	}
	return nil
}
