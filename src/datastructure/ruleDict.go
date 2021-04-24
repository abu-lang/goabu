package datastructure

type RuleDict map[string]*ParsedRule

func (s *RuleDict) Contains(name string) bool {
	_, present := (*s)[name]
	return present
}

func (s *RuleDict) Insert(rule *ParsedRule) {
	(*s)[rule.Name] = rule
}

func (s *RuleDict) Empty() bool {
	return len(*s) == 0
}

func (s *RuleDict) Add(other *RuleDict) {
	if other != nil {
		for name, rule := range *other {
			(*s)[name] = rule
		}
	}
}
