package main_test

import (
	"steel-lang/communication"
	"steel-lang/datastructure"
	"steel-lang/semantics"
	"testing"
)

func TestSingleNode(t *testing.T) {
	memory := datastructure.MakeResources()
	memory.Bool["start"] = false
	memory.Bool["aliqua"] = false
	memory.Integer["magna"] = 0
	e, err := semantics.NewMuSteelExecuter(memory, communication.MakeMemberlistAgent(memory.ResourceNames(), 8000, nil))
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
	e.AddRule(&r1)
	e.AddRule(&r2)
	e.Input([]datastructure.Action{{Resource: "start", Expression: "true"}})
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
	r := datastructure.Rule{
		Name:           "r",
		Events:         []string{"lorem"},
		DefaultActions: nil,
		Task: datastructure.Task{
			Mode:      "for all",
			Condition: `this.Integer["lorem"] > ext.Integer["lorem"]`,
			Actions: []datastructure.Action{
				{Resource: "lorem",
					Expression: `this.Integer["lorem"]`,
				},
			},
		},
	}
	t.Run("TestTwoNodes#1", func(t *testing.T) {
		e1, err := semantics.NewMuSteelExecuter(memory, communication.MakeMemberlistAgent(memory.ResourceNames(), 9001, nil))
		if err != nil {
			t.Fatal(err)
		}
		e1.AddRule(&r)
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
		e2, err := semantics.NewMuSteelExecuter(memory, communication.MakeMemberlistAgent(memory.ResourceNames(), 9002, []string{"127.0.0.1:9001"}))
		if err != nil {
			t.Fatal(err)
		}
		e2.AddRule(&r)
		e2.Input([]datastructure.Action{{Resource: "lorem", Expression: "10"}})
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
	r1 := datastructure.Rule{
		Name:   "r1",
		Events: []string{"ipsum"},
		DefaultActions: []datastructure.Action{
			{Resource: "involved", Expression: "false"},
		},
		Task: datastructure.Task{
			Mode:      "for all",
			Condition: `this.Float["ipsum"] != ext.Float["ipsum"]`,
			Actions: []datastructure.Action{
				{Resource: "involved",
					Expression: `true`,
				},
			},
		},
	}
	r2 := datastructure.Rule{
		Name:           "r2",
		Events:         []string{"involved"},
		DefaultActions: nil,
		Task: datastructure.Task{
			Mode:      "for all",
			Condition: `ext.Bool["involved"] && this.Float["ipsum"] > ext.Float["ipsum"]`,
			Actions: []datastructure.Action{
				{Resource: "ipsum",
					Expression: `this.Float["ipsum"]`,
				},
			},
		},
	}
	t.Run("TestThreeNodes#1", func(t *testing.T) {
		e1, err := semantics.NewMuSteelExecuter(memory, communication.MakeMemberlistAgent(memory.ResourceNames(), 10001, nil))
		if err != nil {
			t.Fatal(err)
		}
		e1.AddRules([]datastructure.Rule{r1, r2})
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
		e2, err := semantics.NewMuSteelExecuter(memory, communication.MakeMemberlistAgent(memory.ResourceNames(), 10002, []string{"127.0.0.1:10001"}))
		if err != nil {
			t.Fatal(err)
		}
		e2.AddRules([]datastructure.Rule{r1, r2})
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
		e3, err := semantics.NewMuSteelExecuter(memory, communication.MakeMemberlistAgent(memory.ResourceNames(), 10003, []string{"127.0.0.1:10001"}))
		if err != nil {
			t.Fatal(err)
		}
		e3.AddRules([]datastructure.Rule{r1, r2})
		e3.Input([]datastructure.Action{{Resource: "ipsum", Expression: "6.0"}})
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
