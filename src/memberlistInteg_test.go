package main_test

import (
	"steel-lang/communication"
	"steel-lang/config"
	"steel-lang/datastructure"
	"steel-lang/semantics"
	"testing"
)

func TestSingleNode(t *testing.T) {
	memory := datastructure.MakeResources()
	memory.Bool["start"] = false
	memory.Bool["aliqua"] = false
	memory.Integer["magna"] = 0
	e, err := semantics.NewMuSteelExecuter(memory, nil, communication.MakeMemberlistAgent(memory.ResourceNames(), 8000, nil, config.TestsLogConfig), config.TestsLogConfig)
	if err != nil {
		t.Fatal(err)
	}
	r1 := "rule r1 on start default magna = 123 + this.magna; for this.aliqua do this.magna = -123;"
	r2 := "rule r2 on magna for this.magna >= this.magna do this.magna = 2 * this.magna + this.magna;"
	e.AddRule(r1)
	e.AddRule(r2)
	e.Input("start = true;")
	for i := 0; i < 3; i++ {
		e.Exec()
	}
	if e.IsStable() {
		t.Error("should not be stable")
	}
	state := e.GetState()
	memory = state.Memory
	if memory.Bool["aliqua"] {
		t.Error("aliqua should be false")
	}
	if !memory.Bool["start"] {
		t.Error("start should be true")
	}
	if memory.Integer["magna"] != 1107 {
		t.Error("magna should be 1107")
	}
}

func TestTwoNodes(t *testing.T) {
	memory := datastructure.MakeResources()
	memory.Integer["lorem"] = 5
	r := "rule r on lorem for all this.lorem > ext.lorem do ext.lorem = this.lorem; "
	rules := []string{r}
	t.Run("TestTwoNodes#1", func(t *testing.T) {
		e1, err := semantics.NewMuSteelExecuter(memory, rules, communication.MakeMemberlistAgent(memory.ResourceNames(), 9001, nil, config.TestsLogConfig), config.TestsLogConfig)
		if err != nil {
			t.Fatal(err)
		}
		t.Parallel()
		for e1.IsStable() {
		}
		e1.Exec()
		if !e1.IsStable() {
			t.Error("should be stable")
		}
		mem1 := e1.GetState().Memory
		if mem1.Integer["lorem"] != 10 {
			t.Error("lorem should be 10")
		}
	})
	t.Run("TestTwoNodes#2", func(t *testing.T) {
		t.Parallel()
		e2, err := semantics.NewMuSteelExecuter(memory, rules, communication.MakeMemberlistAgent(memory.ResourceNames(), 9002, []string{"127.0.0.1:9001"}, config.TestsLogConfig), config.TestsLogConfig)
		if err != nil {
			t.Fatal(err)
		}
		e2.Input("lorem = 10; ")
		if !e2.IsStable() {
			t.Error("should be stable")
		}
		mem2 := e2.GetState().Memory
		if mem2.Integer["lorem"] != 10 {
			t.Error("lorem should be 10")
		}
	})
}

func TestThreeNodes(t *testing.T) {
	memory := datastructure.MakeResources()
	memory.Float["ipsum"] = 3.0
	memory.Bool["involved"] = false
	r1 := "rule r1 on ipsum default involved = false; for all this.ipsum != ext.ipsum do ext.involved = true ; "
	r2 := "rule r2 on involved for all ext.involved && this.ipsum > ext.ipsum do ext.ipsum = this.ipsum;"
	rules := []string{r1, r2}
	t.Run("TestThreeNodes#1", func(t *testing.T) {
		e1, err := semantics.NewMuSteelExecuter(memory, rules, communication.MakeMemberlistAgent(memory.ResourceNames(), 10001, nil, config.TestsLogConfig), config.TestsLogConfig)
		if err != nil {
			t.Fatal(err)
		}
		t.Parallel()
		for e1.GetState().Memory.Float["ipsum"] != 6.5 {
			for e1.IsStable() {
			}
			e1.Exec()
		}
		mem1 := e1.GetState().Memory
		if !mem1.Bool["involved"] {
			t.Error("involved should be true")
		}
	})
	t.Run("TestThreeNodes#2", func(t *testing.T) {
		memory.Float["ipsum"] = 6.5
		e2, err := semantics.NewMuSteelExecuter(memory, rules, communication.MakeMemberlistAgent(memory.ResourceNames(), 10002, []string{"127.0.0.1:10001"}, config.TestsLogConfig), config.TestsLogConfig)
		if err != nil {
			t.Fatal(err)
		}
		t.Parallel()
		for e2.IsStable() {
		}
		for !e2.IsStable() {
			e2.Exec()
		}
		mem2 := e2.GetState().Memory
		if !mem2.Bool["involved"] {
			t.Error("involved should be true")
		}
		if mem2.Float["ipsum"] != 6.5 {
			t.Error("ipsum hould be 6.5")
		}
	})
	t.Run("TestThreeNodes#3", func(t *testing.T) {
		t.Parallel()
		memory.Float["ipsum"] = 3.0
		e3, err := semantics.NewMuSteelExecuter(memory, rules, communication.MakeMemberlistAgent(memory.ResourceNames(), 10003, []string{"127.0.0.1:10001"}, config.TestsLogConfig), config.TestsLogConfig)
		if err != nil {
			t.Fatal(err)
		}
		e3.Input("ipsum = 6.0;")
		e3.Exec()
		for e3.IsStable() {
		}
		e3.Exec()
		if !e3.IsStable() {
			t.Error("should be stable")
		}
		mem3 := e3.GetState().Memory
		if !mem3.Bool["involved"] {
			t.Error("involved should be true")
		}
		if mem3.Float["ipsum"] != 6.0 {
			t.Error("ipsum hould be 6.0")
		}
	})
}
