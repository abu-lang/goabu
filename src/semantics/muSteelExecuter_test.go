package semantics

import (
	"steel-lang/datastructure"
	"testing"
	"time"
)

func TestNewMuSteelExecuter(t *testing.T) {
	invalid := datastructure.MakeResources()
	invalid.Bool["lorem42"] = false
	invalid.Float["lorem42"] = 3.14
	_, err := NewMuSteelExecuter(invalid, nil, MakeMockAgent())
	if err == nil {
		t.Error("should return error with invalid memory")
	}
	memory := datastructure.MakeResources()
	memory.Integer["dolor"] = 42
	memory.Text["sit"] = "amet"
	started := MakeMockAgent()
	started.Start()
	_, err = NewMuSteelExecuter(invalid, nil, started)
	if err == nil {
		t.Error("should return error with started agent")
	}
	e, err := NewMuSteelExecuter(memory, nil, MakeMockAgent())
	if err != nil {
		t.Fatal(err.Error())
	}
	if !e.IsStable() {
		t.Error("should be stable")
	}
	if !e.agent.IsRunning() {
		t.Error("agent should be running")
	}
}

func TestAddRules(t *testing.T) {
	local := datastructure.Rule{
		Name:           "local",
		Events:         []string{"trigger", "executed"},
		DefaultActions: nil,
		Task: datastructure.Task{
			Mode:      "for",
			Condition: `!this.Bool["executed"]`,
			Actions: []datastructure.Action{
				{Resource: "trigger",
					Expression: `"activated"`,
				},
			},
		},
	}
	global := datastructure.Rule{
		Name:   "global",
		Events: []string{"trigger"},
		DefaultActions: []datastructure.Action{
			{Resource: "executed",
				Expression: `this.Bool["executed"] && false`,
			},
		},
		Task: datastructure.Task{
			Mode:      "for all",
			Condition: `this.Text["trigger"] != ext.Text["trigger"] && this.Text["trigger"] == "activated"`,
			Actions: []datastructure.Action{
				{Resource: "trigger",
					Expression: `this.Text["trigger"]`,
				},
			},
		},
	}
	memory := datastructure.MakeResources()
	memory.Bool["executed"] = false
	memory.Text["trigger"] = "activable"
	e, err := NewMuSteelExecuter(memory, nil, MakeMockAgent())
	if err != nil {
		t.Fatal(err.Error())
	}
	e.AddRules([]datastructure.Rule{local, global})
	if len(e.localLibrary) != 2 {
		t.Error("localLibrary should have 2 dicts")
	}
	if len(e.globalLibrary) != 1 {
		t.Error("localLibrary should have 1 dict")
	}
	if len(e.localLibrary["trigger"]) != 1 {
		t.Error("trigger local dict should have 1 rule")
	}
	if len(e.localLibrary["executed"]) != 1 {
		t.Error("executed local dict should have 1 rule")
	}
	if len(e.globalLibrary["trigger"]) != 1 {
		t.Error("trigger global dict should have 1 rule")
	}
	if !e.localLibrary["trigger"].Contains("local") {
		t.Error("trigger should contain local")
	}
	if !e.localLibrary["executed"].Contains("local") {
		t.Error("executed should contain local")
	}
	if !e.globalLibrary["trigger"].Contains("global") {
		t.Error("trigger should contain global")
	}
}

func TestAddPool(t *testing.T) {
	memory := datastructure.MakeResources()
	memory.Float["elit"] = 5.0
	memory.Integer["consectetur"] = 5
	memory.Text["adipiscing"] = "eiusmod"
	memory.Time["tempor"] = time.Unix(0, 0)
	e, err := NewMuSteelExecuter(memory, nil, MakeMockAgent())
	if err != nil {
		t.Fatal(err.Error())
	}
	e.AddPool([][]datastructure.Action{
		{{Resource: "elit", Expression: `2.71828`}},
		{{Resource: "consectetur", Expression: `this.Integer["consectetur"]*7`}, {Resource: "adipiscing", Expression: `""`}},
		{{Resource: "tempor", Expression: `MakeTime(2021, 6, 5, 0, 0, 0)`}},
	})
	poolLength := len(e.pool)
	for i := 0; i < 3; i++ {
		if poolLength != len(e.pool) {
			t.Errorf("pool should have length %d", poolLength)
		}
		e.Exec()
		poolLength--
	}
	if !e.IsStable() {
		t.Error("should be stable")
	}
	mem := e.memory.GetResources()
	if mem.Float["elit"] != 2.71828 {
		t.Error("elit should be 2.71828")
	}
	if mem.Integer["consectetur"] != 35 {
		t.Error("consectetur should be 35")
	}
	if mem.Text["adipiscing"] != "" {
		t.Error("adipiscing should be \"\"")
	}
	if mem.Time["tempor"] != time.Date(2021, 6, 5, 0, 0, 0, 0, time.Local) {
		t.Error("tempor should be 2021-06-05 00:00:00")
	}
}

func TestLocal(t *testing.T) {
	startCooling := datastructure.Rule{
		Name:           "startCooling",
		Events:         []string{"temperature"},
		DefaultActions: nil,
		Task: datastructure.Task{
			Mode:      "for",
			Condition: `"hihj".Replace("hj", "gh") == this.Text["temperature"]`,
			Actions: []datastructure.Action{
				{Resource: "cooling",
					Expression: `true`,
				},
				{Resource: "counter",
					Expression: `3 + 2 * 1 - 1 * 3`,
				},
			},
		},
	}
	counter := datastructure.Rule{
		Name:           "counter",
		Events:         []string{"counter", "cooling"},
		DefaultActions: nil,
		Task: datastructure.Task{
			Mode:      "for",
			Condition: `this.Integer["counter"] > 0 && this.Bool["cooling"]`,
			Actions: []datastructure.Action{
				{Resource: "counter",
					Expression: `this.Integer["counter"] - 1`,
				},
			},
		},
	}
	stopCooling := datastructure.Rule{
		Name:           "stopCooling",
		Events:         []string{"counter"},
		DefaultActions: nil,
		Task: datastructure.Task{
			Mode:      "for",
			Condition: `this.Integer["counter"] == 0 && this.Bool["cooling"]`,
			Actions: []datastructure.Action{
				{Resource: "cooling",
					Expression: `!this.Bool["cooling"]`,
				},
				{Resource: "temperature",
					Expression: `"NORMAL".ToLower()`,
				},
			},
		},
	}
	memory := datastructure.MakeResources()
	memory.Integer["counter"] = 42
	memory.Bool["cooling"] = false
	memory.Text["temperature"] = "low"
	e, err := NewMuSteelExecuter(memory, nil, MakeMockAgent())
	if err != nil {
		t.Fatal(err.Error())
	}
	err = e.StopAgent()
	if err != nil {
		t.Fatal(err.Error())
	}
	e.AddRules([]datastructure.Rule{startCooling, counter, stopCooling})
	e.AddActions([]datastructure.Action{{Resource: "temperature", Expression: `"high"`}})
	execs := 0
	for !e.IsStable() {
		if len(e.pool) != 1 {
			t.Error("pool should have length 1")
		}
		e.Exec()
		execs++
	}
	if execs != 5 {
		t.Error("should be stable after 5 calls to Exec")
	}
	mem := e.memory.GetResources()
	if mem.Bool["cooling"] {
		t.Error("cooling should be false")
	}
	if mem.Integer["counter"] != 0 {
		t.Error("counter should be 0")
	}
	if mem.Text["temperature"] != "normal" {
		t.Error("temperature should be \"normal\"")
	}
}

func TestReceiveExternalActions(t *testing.T) {
	r1 := datastructure.Rule{
		Name:           "r1",
		Events:         []string{"elit"},
		DefaultActions: nil,
		Task: datastructure.Task{
			Mode:      "for all",
			Condition: `ext.Float["elit"] > 0 || ext.Bool["labore"]`,
			Actions: []datastructure.Action{
				{Resource: "elit",
					Expression: `0`,
				},
				{Resource: "consectetur",
					Expression: `"-10"`,
				},
			},
		},
	}
	r2 := datastructure.Rule{
		Name:           "r2",
		Events:         []string{"consectetur"},
		DefaultActions: nil,
		Task: datastructure.Task{
			Mode:      "for all",
			Condition: `ext.Integer["consectetur"] < 0`,
			Actions: []datastructure.Action{
				{Resource: "elit",
					Expression: `ext.Float["elit"] * 2 + 3.14`,
				},
				{Resource: "adipiscing",
					Expression: `ext.Text["incididunt"]`,
				},
				{Resource: "tempor",
					Expression: `MakeTime(2000, 1, 1, 0, 0, 0)`,
				},
				{Resource: "labore",
					Expression: `false`,
				},
			},
		},
	}
	memory := datastructure.MakeResources()
	memory.Float["elit"] = -100.0
	memory.Integer["consectetur"] = 30
	memory.Text["adipiscing"] = "sed"
	memory.Time["tempor"] = time.Unix(0, 0)

	memory.Text["incididunt"] = "ut"
	memory.Bool["labore"] = true
	e, err := NewMuSteelExecuter(memory, nil, MakeMockAgent())
	if err != nil {
		t.Fatal(err.Error())
	}
	e.AddRule(r1)
	e.AddRule(r2)

	// remove some resources
	mem := e.memory.GetResources()
	delete(mem.Bool, "labore")
	delete(mem.Text, "incididunt")
	e.types = e.memory.GetTypes()

	e.AddActions([]datastructure.Action{{Resource: "elit", Expression: `100.0`}})
	e.AddActions([]datastructure.Action{{Resource: "consectetur", Expression: `-2`}})
	execs := 0
	for !e.IsStable() {
		e.Exec()
		execs++
		if mem.Text["adipiscing"] != "sed" {
			t.Error("adipiscing should be \"sed\"")
		}
	}
	if execs != 3 {
		t.Error("should be stable after 3 calls to Exec")
	}
	if mem.Float["elit"] != 203.14 {
		t.Error("elit should be 203.14")
	}
	if mem.Integer["consectetur"] != -2 {
		t.Error("counter should be 0")
	}
	if mem.Time["tempor"] != time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local) {
		t.Error("tempor should be 2000-01-01 00:00:00")
	}
}

func TestForall(t *testing.T) {
	memory := datastructure.MakeResources()
	memory.Bool["start"] = false
	memory.Bool["aliqua"] = false
	memory.Integer["magna"] = 0
	e, err := NewMuSteelExecuter(memory, nil, MakeMockAgent())
	if err != nil {
		t.Fatal(err)
	}
	r1 := datastructure.Rule{
		Name:   "r1",
		Events: []string{"start"},
		DefaultActions: []datastructure.Action{
			{Resource: "magna",
				Expression: `123 + this.Integer["magna"]`,
			},
		},
		Task: datastructure.Task{
			Mode:      "for all",
			Condition: `ext.Bool["aliqua"]`,
			Actions: []datastructure.Action{
				{Resource: "magna",
					Expression: `-123`,
				},
			},
		},
	}
	r2 := datastructure.Rule{
		Name:           "r2",
		Events:         []string{"magna"},
		DefaultActions: nil,
		Task: datastructure.Task{
			Mode:      "for all",
			Condition: `this.Integer["magna"] >= ext.Integer["magna"]`,
			Actions: []datastructure.Action{
				{Resource: "magna",
					Expression: `2 * this.Integer["magna"] + ext.Integer["magna"]`,
				},
			},
		},
	}
	e.AddRule(r1)
	e.AddRule(r2)
	e.Input([]datastructure.Action{{Resource: "start", Expression: "true"}})
	magnas := []int64{123, 369, 1107}
	mem := e.memory.GetResources()
	for i := 0; i < 3; i++ {
		e.Exec()
		if len(e.pool) != 1 {
			t.Error("pool should have length 1")
		}
		if mem.Bool["aliqua"] {
			t.Error("aliqua should be false")
		}
		if !mem.Bool["start"] {
			t.Error("start should be true")
		}
		if mem.Integer["magna"] != magnas[i] {
			t.Errorf("magna should be %d", magnas[i])
		}
	}
}
