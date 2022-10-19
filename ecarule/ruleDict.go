// Copyright 2021 Massimo Comuzzo, Michele Pasqua and Marino Miculan
// SPDX-License-Identifier: Apache-2.0

package ecarule

type RuleDict map[string]*Rule

func MakeRuleDict() RuleDict {
	return make(map[string]*Rule)
}

func (rules RuleDict) Has(name string) bool {
	_, present := rules[name]
	return present
}

// Precondition: rules != nil
func (rules RuleDict) Insert(rule *Rule) {
	rules[rule.Name] = rule
}

func (rules RuleDict) Empty() bool {
	return len(rules) == 0
}

// Precondition: dst != nil
func (dst RuleDict) Add(src RuleDict) {
	for name, rule := range src {
		dst[name] = rule
	}
}
