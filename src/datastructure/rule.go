package datastructure

import (
	"fmt"
)

type Rule struct {
	Name           string
	Events         []string
	DefaultActions []Action
	Task           Task
}

type Task struct {
	Mode      string
	Condition string
	Actions   []Action
}

type Action struct {
	Resource   string
	Expression string
}

func (r Rule) String() string {
	str := fmt.Sprintf("%s on %v\n", r.Name, r.Events)
	if len(r.DefaultActions) > 0 {
		str = str + "Default:\n"
		for _, act := range r.DefaultActions {
			str = str + act.String() + "\n"
		}
	}
	return str + r.Task.String()
}

func (t Task) String() string {
	str := fmt.Sprintln(t.Mode + " " + t.Condition + " do:")
	for _, act := range t.Actions {
		str = str + act.String() + "\n"
	}
	return str
}

func (a Action) String() string {
	return "(" + a.Resource + "," + a.Expression + ")"
}
