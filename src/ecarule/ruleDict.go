package ecarule

type RuleDict map[string]*Rule

func MakeRuleDict() RuleDict {
	return make(map[string]*Rule)
}

func (rules RuleDict) Contains(name string) bool {
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
