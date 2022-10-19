// Copyright 2021 Massimo Comuzzo, Michele Pasqua and Marino Miculan
// SPDX-License-Identifier: Apache-2.0

package goabu_test

import (
	"fmt"
	"testing"

	"github.com/abu-lang/goabu"
	"github.com/abu-lang/goabu/communication"
	"github.com/abu-lang/goabu/config"
	"github.com/abu-lang/goabu/memory"
)

func TestSingleNode(t *testing.T) {
	memory := memory.MakeResources()
	memory.Bool["start"] = false
	memory.Bool["aliqua"] = false
	memory.Integer["magna"] = 0
	e, err := goabu.NewExecuter(memory, nil, communication.NewMemberlistAgent(t.Name(), 8000, config.TestsLogConfig), config.TestsLogConfig)
	if err != nil {
		t.Fatal(err)
	}
	e.SetOptimisticExec(*goabu.Optimistic)
	e.SetOptimisticInput(*goabu.Optimistic)
	r1 := "rule r1 on start default magna = 123 + this.magna; for this.aliqua do this.magna = -123;"
	r2 := "rule r2 on magna for this.magna >= this.magna do this.magna = 2 * this.magna + this.magna;"
	e.AddRules(r1)
	e.AddRules(r2)
	e.Input("start = true;")
	for i := 0; i < 3; i++ {
		e.Exec()
	}
	if e.DoIfStable(func() {}) {
		t.Error("should not be stable")
	}
	state := e.TakeState()
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
	memory := memory.MakeResources()
	memory.Integer["lorem"] = 5
	r := "rule r on lorem for all this.lorem > ext.lorem do ext.lorem = this.lorem; "
	rules := []string{r}
	t.Run("TestTwoNodes#1", func(t *testing.T) {
		e1, err := goabu.NewExecuter(memory, rules, communication.NewMemberlistAgent(t.Name(), 9001, config.TestsLogConfig), config.TestsLogConfig)
		if err != nil {
			t.Fatal(err)
		}
		e1.SetOptimisticExec(*goabu.Optimistic)
		e1.SetOptimisticInput(*goabu.Optimistic)
		t.Parallel()
		for e1.DoIfStable(func() {}) {
		}
		e1.Exec()
		if !e1.DoIfStable(func() {}) {
			t.Error("should be stable")
		}
		mem1 := e1.TakeState().Memory
		if mem1.Integer["lorem"] != 10 {
			t.Error("lorem should be 10")
		}
	})
	t.Run("TestTwoNodes#2", func(t *testing.T) {
		t.Parallel()
		e2, err := goabu.NewExecuter(memory, rules, communication.NewMemberlistAgent(t.Name(), 9002, config.TestsLogConfig, "127.0.0.1:9001"), config.TestsLogConfig)
		if err != nil {
			t.Fatal(err)
		}
		e2.SetOptimisticExec(*goabu.Optimistic)
		e2.SetOptimisticInput(*goabu.Optimistic)
		e2.Input("lorem = 10; ")
		if !e2.DoIfStable(func() {}) {
			t.Error("should be stable")
		}
		mem2 := e2.TakeState().Memory
		if mem2.Integer["lorem"] != 10 {
			t.Error("lorem should be 10")
		}
	})
}

func TestThreeNodes(t *testing.T) {
	memory := memory.MakeResources()
	memory.Float["ipsum"] = 3.0
	memory.Bool["involved"] = false
	r1 := "rule r1 on ipsum default involved = true for all ipsum != ext.ipsum do involved = true"
	r2 := "rule r2 on involved for all involved && ipsum > ext.ipsum do ipsum = this.ipsum"
	rules := []string{r1, r2}
	t.Run("TestThreeNodes#1", func(t *testing.T) {
		e1, err := goabu.NewExecuter(memory, rules, communication.NewMemberlistAgent(t.Name(), 10001, config.TestsLogConfig), config.TestsLogConfig)
		if err != nil {
			t.Fatal(err)
		}
		e1.SetOptimisticExec(*goabu.Optimistic)
		e1.SetOptimisticInput(*goabu.Optimistic)
		t.Parallel()
		for e1.TakeState().Memory.Float["ipsum"] != 6.5 {
			e1.Exec()
		}
		if !e1.TakeState().Memory.Bool["involved"] {
			t.Error("involved should be true")
		}
	})
	t.Run("TestThreeNodes#2", func(t *testing.T) {
		memory.Float["ipsum"] = 6.5
		e2, err := goabu.NewExecuter(memory, rules, communication.NewMemberlistAgent(t.Name(), 10002, config.TestsLogConfig, "127.0.0.1:10001"), config.TestsLogConfig)
		if err != nil {
			t.Fatal(err)
		}
		e2.SetOptimisticExec(*goabu.Optimistic)
		e2.SetOptimisticInput(*goabu.Optimistic)
		t.Parallel()
		for e2.DoIfStable(func() {}) {
		}
		e2.Exec()
		mem2 := e2.TakeState().Memory
		if !mem2.Bool["involved"] {
			t.Error("involved should be true")
		}
		if mem2.Float["ipsum"] != 6.5 {
			t.Error("ipsum should be 6.5")
		}
	})
	t.Run("TestThreeNodes#3", func(t *testing.T) {
		t.Parallel()
		memory.Float["ipsum"] = 3.0
		e3, err := goabu.NewExecuter(memory, rules, communication.NewMemberlistAgent(t.Name(), 10003, config.TestsLogConfig, "127.0.0.1:10001"), config.TestsLogConfig)
		if err != nil {
			t.Fatal(err)
		}
		e3.SetOptimisticExec(*goabu.Optimistic)
		e3.SetOptimisticInput(*goabu.Optimistic)
		e3.Input("ipsum = 6.0;")
		for e3.TakeState().Memory.Float["ipsum"] != 6.5 {
			e3.Exec()
		}
		if !e3.TakeState().Memory.Bool["involved"] {
			t.Error("involved should be true")
		}
	})
}

func TestInc(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}

	const k = 5
	const m = 4
	memory := memory.MakeResources()
	memory.Integer["A"] = 0
	rules := []string{fmt.Sprintf("rule INC on A for all A < %d && ext.A < %d do A = this.A + 1", m, m)}
	agts := make([]*goabu.Executer, 0, k)
	for i := 0; i < k; i++ {
		t.Run(fmt.Sprintf("TestInc#%d", i+1), func(t *testing.T) {
			cfg := config.TestsLogConfig
			cfg.Level = config.LogInfo
			var prec []string
			if i > 0 {
				cfg.Level = config.LogWarning
				prec = []string{fmt.Sprintf("127.0.0.1:%d", 11000+i)}
			}
			e, err := goabu.NewExecuter(memory, rules, communication.NewMemberlistAgent(t.Name(), 11000+i+1, cfg, prec...), cfg)
			if err != nil {
				t.Fatal(err)
			}
			e.SetOptimisticExec(*goabu.Optimistic)
			e.SetOptimisticInput(*goabu.Optimistic)
			agts = append(agts, e)
		})
	}
	done := make([]<-chan struct{}, 0, k)
	for i := 0; i < k; i++ {
		done = append(done, incNode(m, agts[i]))
	}
	for _, d := range done {
		<-d
	}
}

func incNode(m int64, e *goabu.Executer) <-chan struct{} {
	res := make(chan struct{})
	go func() {
		e.Input("A = 1")
		state := e.TakeState()
		for state.Memory.Integer["A"] != m {
			if state.Memory.Integer["A"] > m {
				panic(fmt.Sprintf("A should be <= %d", m))
			}
			e.Exec()
			state = e.TakeState()
		}
		for !e.DoIfStable(func() {}) {
			e.Exec()
			if e.TakeState().Memory.Integer["A"] > m {
				panic(fmt.Sprintf("A should be <= %d", m))
			}
		}
		res <- struct{}{}
	}()
	return res
}

func TestInvariants(t *testing.T) {
	m1 := memory.MakeResources()
	m1.Integer["odd"] = 5
	r1 := "rule A on odd for all true do ext.even = ext.even + this.odd; "
	t.Run("TestInvariants#1", func(t *testing.T) {
		e1, err := goabu.NewExecuter(m1, []string{r1}, communication.NewMemberlistAgent(t.Name(), 12001, config.TestsLogConfig),
			config.TestsLogConfig, "odd % 2 == 1")
		if err != nil {
			t.Fatal(err)
		}
		e1.SetOptimisticExec(*goabu.Optimistic)
		e1.SetOptimisticInput(*goabu.Optimistic)
		t.Parallel()
		for e1.DoIfStable(func() {}) {
		}
		e1.Exec()
		if !e1.DoIfStable(func() {}) {
			t.Error("should be stable")
		}
		mem1 := e1.TakeState().Memory
		if mem1.Integer["odd"] != 17 {
			t.Error("odd should be 17")
		}
	})
	m2 := memory.MakeResources()
	m2.Integer["even"] = 10
	r2 := "rule B on even for all true do ext.odd = this.even + ext.odd; "
	t.Run("TestInvariants#2", func(t *testing.T) {
		t.Parallel()
		e2, err := goabu.NewExecuter(m2, []string{r2},
			communication.NewMemberlistAgent(t.Name(), 12002, config.TestsLogConfig, "127.0.0.1:12001"), config.TestsLogConfig,
			"even > 0", "even % 2 == 0")
		if err != nil {
			t.Fatal(err)
		}
		e2.SetOptimisticExec(*goabu.Optimistic)
		e2.SetOptimisticInput(*goabu.Optimistic)
		e2.Input("even = even + 2")
		for e2.DoIfStable(func() {}) {
		}
		e2.Exec()
		if !e2.DoIfStable(func() {}) {
			t.Error("should be stable")
		}
		mem2 := e2.TakeState().Memory
		if mem2.Integer["even"] != 12 {
			t.Error("even should be 12")
		}
	})
}
