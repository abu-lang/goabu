// Copyright 2023 Massimo Comuzzo, Michele Pasqua and Marino Miculan
// SPDX-License-Identifier: Apache-2.0

package parser

import (
	"testing"

	"github.com/hyperjumptech/grule-rule-engine/ast"
)

// isClean checks if the [*ecaruleParserListener] was correctly reset to a clean state.
func (l *ruleParser) isClean() bool {
	if l.rules != nil || len(l.local.PreviousNode) > 0 || l.local.Stack.Len() > 0 || l.local.StopParse || len(l.local.ErrorCallback.Errors) > 0 {
		return false
	}
	return true
}

// TestReset tests [*ecaruleParserListener]'s reset method.
func TestReset(t *testing.T) {
	tests := []struct {
		idx    int
		name   string
		events []string
		rule   string
	}{
		//  {_, name, events, rule},
		{1, "MyGlobalRule", []string{"foo"}, "rule MyGlobalRule on foo for all foo >= ext.foo do foo = ext.foo + foo"},
		{2, "MyLocalRule", []string{"foo", "bar"}, "rule MyLocalRule on foo bar for \"qux\" == bar do foo = foo * 2, bar = \"grault\""},
	}
	types := map[string]string{
		"foo": "Integer",
		"bar": "Integer",
	}
	wm := ast.NewWorkingMemory("", "")
	p := New(types, wm).(*goabuParser)
	if !p.listener.isClean() {
		t.Error("listener should be clean")
	}
	for _, test := range tests {
		p.reset(test.rule)
		if !p.listener.isClean() {
			t.Error(test.idx, "->", "listener should be clean")
		}
		rules, errs := p.Parse(test.rule)
		if len(errs) > 0 {
			t.Error(test.idx, "->", "error in parsing rule")
		}
		if len(rules) != 1 {
			t.Fatal(test.idx, "->", "error in parsing rule")
		}
		rule := rules[0]
		if rule.Name != test.name || len(rule.Events) != len(test.events) {
			t.Error(test.idx, "->", "error in parsing rule")
		}
		for _, e := range test.events {
			present := false
			for _, v := range rule.Events {
				if v == e {
					present = true
				}
			}
			if !present {
				t.Error(test.idx, "->", "error in parsing rule")
			}
		}
	}
}

// TestDefault tests parsing of default actions.
func TestDefault(t *testing.T) {
	tests := []struct {
		idx  int
		rule string
	}{
		//  {_, rule},
		{1, "rule r1 on start default magna = 123 + this.magna, for all ext.aliqua do ext.magna = -123,"},
		{2, "rule r2 on ipsum default involved = true for all ipsum != ext.ipsum do involved = true"},
		{3, "rule R on foo default baz = 0.0, bar = \"octocat\" for all ext.foo < 0 do foo = ext.foo * -1"},
	}
	types := map[string]string{
		"start":    "Bool",
		"magna":    "Integer",
		"ipsum":    "Bool",
		"involved": "Bool",
		"bar":      "Text",
		"foo":      "Integer",
		"baz":      "Float",
	}
	wm := ast.NewWorkingMemory("", "")
	p := New(types, wm).(*goabuParser)
	exp, err := newBooleanLiteralExpression(p.listener.local.KnowledgeBase.WorkingMemory, true)
	if err != nil {
		t.Fatal(err)
	}
	for _, test := range tests {
		rules, errs := p.Parse(test.rule)
		if len(errs) > 0 {
			t.Error(test.idx, "->", "error in parsing rule", errs)
		}
		if len(rules) != 1 {
			t.Fatal(test.idx, "->", "error in parsing rule")
		}
		rule := rules[0]
		found := 0
		for _, t := range rule.LocalTasks {
			if t.Condition.AstID == exp.AstID {
				found++
			}
		}
		if found != 1 {
			t.Error(test.idx, "->", "error in default action task")
		}
	}
}

// TestMultipleRules tests the parsing of multiple rules in the same input stream.
func TestMultipleRules(t *testing.T) {
	rules := []string{
		"rule A on odd for all true do ext.even = ext.even + this.odd,",
		"rule r on lorem for all this.lorem > ext.lorem do ext.lorem = this.lorem,",
		`rule batteryCheck on battery
			for all (battery < 5 && ext.battery > 80)
			do ext.help_lat = position_lat,
				ext.help_lon = position_lon`,
		`rule setRescue on help_lat help_lon
			for (AbsInt((position_lat - help_lat)) < threshold && AbsInt((position_lon - help_lon)) < threshold)
			do mode = "rescue"`,
		"rule R1 on motor for this.motor > 0 && this.motor < 255 do motor = this.motor + 60",
		"rule R2 on button2 for all this.button1 && this.button2 do ext.led = !ext.led",
	}
	names := []string{"A", "r", "batteryCheck", "setRescue", "R1", "R2"}
	types := map[string]string{
		"odd":          "Integer",
		"even":         "Integer",
		"threshold":    "Integer",
		"battery":      "Integer",
		"help_lat":     "Integer",
		"help_lon":     "Integer",
		"position_lat": "Integer",
		"position_lon": "Integer",
		"motor":        "Integer",
		"lorem":        "Float",
		"mode":         "Text",
		"button1":      "Bool",
		"button2":      "Bool",
		"led":          "Bool",
	}
	tests := []struct {
		idx      int
		ruleIdxs []int
	}{
		//  {_, ruleIdxs},
		{1, []int{0, 1}},
		{2, []int{2, 5, 3}},
		{3, []int{4, 2}},
		{4, []int{0, 1, 3, 5}},
		{5, []int{3}},
		{6, []int{0, 2, 4}},
	}
	wm := ast.NewWorkingMemory("", "")
	p := New(types, wm).(*goabuParser)
	for _, test := range tests {
		inp := ""
		for _, i := range test.ruleIdxs {
			inp = inp + " " + rules[i]
		}
		parsed, errs := p.Parse(inp)
		if len(errs) > 0 {
			t.Fatal(test.idx, "->", "error in parsing rules", errs)
		}
		if len(parsed) != len(test.ruleIdxs) {
			t.Error(test.idx, "->", "mismatched parsed rules number")
		}
		for _, i := range test.ruleIdxs {
			name := names[i]
			found := false
			for _, rule := range parsed {
				if rule.Name == name {
					found = true
					break
				}
			}
			if !found {
				t.Error(test.idx, "->", "missing rule: ", name)
			}
		}
	}
}
