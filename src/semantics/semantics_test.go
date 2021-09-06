package semantics

import (
	"flag"
	"fmt"
	"steel-lang/config"
	"steel-lang/memory"
	"testing"
	"time"
)

var optimistic = flag.Bool("opt", false, "set optimistic concurrency control")

func TestInvalidNames(t *testing.T) {
	names := []string{"", "  abc", "def ", "ip sum", "this", "ext", "rule", "on", "default", "for", "FoR", "all", "do", "10sit",
		"a,met", "=", "123", "."}
	for _, n := range names {
		test := fmt.Sprintf("TestInvalidNames#\"%s\"", n)
		t.Run(test, func(t *testing.T) {
			mem := memory.MakeResources()
			mem.Integer[n] = 42
			_, err := NewExecuter(mem, nil, MakeMockAgent(), config.TestsLogConfig)
			if err == nil {
				t.Error(test + " failed")
			}
		})
	}
}

func TestNewExecuter(t *testing.T) {
	empty := memory.MakeResources()
	_, err := NewExecuter(empty, nil, MakeMockAgent(), config.TestsLogConfig)
	if err == nil {
		t.Error("should return error with no resources")
	}
	invalid := memory.MakeResources()
	invalid.Bool["lorem42"] = false
	invalid.Float["lorem42"] = 3.14
	_, err = NewExecuter(invalid, nil, MakeMockAgent(), config.TestsLogConfig)
	if err == nil {
		t.Error("should return error with duplicated resources")
	}
	memory := memory.MakeResources()
	memory.Integer["dolor"] = 42
	memory.Text["sit"] = "amet"
	started := MakeMockAgent()
	started.Start()
	_, err = NewExecuter(invalid, nil, started, config.TestsLogConfig)
	if err == nil {
		t.Error("should return error with started agent")
	}
	e, err := NewExecuter(memory, nil, MakeMockAgent(), config.TestsLogConfig)
	if err != nil {
		t.Fatal(err.Error())
	}
	if !e.DoIfStable(func() {}) {
		t.Error("should be stable")
	}
	if !e.agent.IsRunning() {
		t.Error("agent should be running")
	}
}

func TestAddRules(t *testing.T) {
	local := `rule local on trigger executed
		for !this.executed do
		trigger = "activated";`

	global := `rule global on trigger
		default executed = this.executed && false;
		for all this.trigger != ext.trigger && this.trigger == "activated"
		do ext.trigger = this.trigger;`

	memory := memory.MakeResources()
	memory.Bool["executed"] = false
	memory.Text["trigger"] = "activable"
	e, err := NewExecuter(memory, nil, MakeMockAgent(), config.TestsLogConfig)
	if err != nil {
		t.Fatal(err.Error())
	}
	e.AddRules(local, global)
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
	if !e.localLibrary["trigger"].Has("local") {
		t.Error("trigger should contain local")
	}
	if !e.localLibrary["executed"].Has("local") {
		t.Error("executed should contain local")
	}
	if !e.globalLibrary["trigger"].Has("global") {
		t.Error("trigger should contain global")
	}
}

func TestAddPool(t *testing.T) {
	memory := memory.MakeResources()
	memory.Float["elit"] = 5.0
	memory.Integer["consectetur"] = 5
	memory.Text["adipiscing"] = "eiusmod"
	memory.Time["tempor"] = time.Unix(0, 0)
	e, err := NewExecuter(memory, nil, MakeMockAgent(), config.TestsLogConfig)
	if err != nil {
		t.Fatal(err.Error())
	}
	e.SetOptimisticExec(*optimistic)
	e.SetOptimisticInput(*optimistic)
	e.addPool([]string{
		"elit = 2.71828;",
		`consectetur = this.consectetur * 7; adipiscing = "";`,
		`tempor = MakeTime(2021, 6, 5, 0, 0, 0) ;`,
	})
	poolLength := len(e.pool)
	for i := 0; i < 3; i++ {
		if poolLength != len(e.pool) {
			t.Errorf("pool should have length %d", poolLength)
		}
		e.Exec()
		poolLength--
	}
	if !e.DoIfStable(func() {}) {
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
	startCooling := `rule startCooling on temperature
		for "hihj".Replace("hj", "gh") == this.temperature
		do  cooling = true;
			counter = 3 + 2 * 1 - 1 * 3`

	counter := `rule counter on counter cooling
		for this.counter > 0 && this.cooling
		do this.counter = this.counter - 1;`

	stopCooling := `rule stopCooling on counter
		for this.counter == 0 && this.cooling
		do  this.cooling = !this.cooling;
			this.temperature = "NORMAL".ToLower()`

	memory := memory.MakeResources()
	memory.Integer["counter"] = 42
	memory.Bool["cooling"] = false
	memory.Text["temperature"] = "low"
	e, err := NewExecuter(memory, nil, MakeMockAgent(), config.TestsLogConfig)
	if err != nil {
		t.Fatal(err.Error())
	}
	e.SetOptimisticExec(*optimistic)
	e.SetOptimisticInput(*optimistic)
	err = e.StopAgent()
	if err != nil {
		t.Fatal(err.Error())
	}
	e.AddRules(startCooling, counter, stopCooling)
	e.addActions(`temperature = "high"`)
	execs := 0
	for !e.DoIfStable(func() {}) {
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
	r1 := `rule r1 on elit
		for all ext.elit > 0 || ext.labore
		do  ext.elit = 0;
			ext.consectetur = "-10";`

	r2 := `rule r2 on consectetur
		for all ext.consectetur < 0
		do  ext.elit = ext.elit * 2 + 3.14;
			ext.adipiscing = ext.incididunt;
			ext.tempor = MakeTime(2000, 1, 1, 0, 0, 0);
			ext.labore = false `

	memory := memory.MakeResources()
	memory.Float["elit"] = -100.0
	memory.Integer["consectetur"] = 30
	memory.Text["adipiscing"] = "sed"
	memory.Time["tempor"] = time.Unix(0, 0)

	memory.Text["incididunt"] = "ut"
	memory.Bool["labore"] = true
	e, err := NewExecuter(memory, nil, MakeMockAgent(), config.TestsLogConfig)
	if err != nil {
		t.Fatal(err.Error())
	}
	e.SetOptimisticExec(*optimistic)
	e.SetOptimisticInput(*optimistic)
	e.AddRules(r1)
	e.AddRules(r2)

	// remove some resources
	mem := e.memory.GetResources()
	delete(mem.Bool, "labore")
	delete(mem.Text, "incididunt")
	e.types = e.memory.Types()

	e.addActions("elit = 100.0;")
	e.addActions("consectetur = -2;")
	execs := 0
	for !e.DoIfStable(func() {}) {
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
	memory := memory.MakeResources()
	memory.Bool["start"] = false
	memory.Bool["aliqua"] = false
	memory.Integer["magna"] = 0
	e, err := NewExecuter(memory, nil, MakeMockAgent(), config.TestsLogConfig)
	if err != nil {
		t.Fatal(err)
	}
	e.SetOptimisticExec(*optimistic)
	e.SetOptimisticInput(*optimistic)
	r1 := "rule r1 on start default magna = 123 + this.magna; for all ext.aliqua do ext.magna = -123;"
	r2 := "rule r2 on magna for all this.magna >= ext.magna do ext.magna = 2 * this.magna + ext.magna;"
	e.AddRules(r1)
	e.AddRules(r2)
	e.Input("start = true;")
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
