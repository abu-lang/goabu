package datastructure

type RuleDict map[string]*ParsedRule

func MakeRuleDict() RuleDict {
	return make(map[string]*ParsedRule)
}

func (rules RuleDict) Contains(name string) bool {
	_, present := rules[name]
	return present
}

// Precondition: rules != nil
func (rules RuleDict) Insert(rule *ParsedRule) {
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
