package datastructure

import (
	"fmt"
	"strconv"
)

type Action struct {
	Resource string
	External bool
	Expression string
}

func PrintAction(syntaction Action) string {
	return "(" + syntaction.Resource + "," + strconv.FormatBool(syntaction.External) + "," + syntaction.Expression + ")"
}

type Task struct {
	Mode string
	Exp  string
	Actions []Action
}

func PrintTask(task Task) string {
	str := fmt.Sprintln(task.Mode + " " + task.Exp + " do:")
	for _, act := range task.Actions {
		str = str + fmt.Sprintln(PrintAction(act))
	}
	return str
}

type Rule struct {
	Name    string
	Event     []string
	DefaultActions []Action
	Task    Task
}

func PrintRule(rule Rule) string {
	str := fmt.Sprintf("%s on %v\n", rule.Name, rule.Event)
	if rule.DefaultActions != nil {
		str = str + "Default:\n"
		for _, act := range rule.DefaultActions {
			str = str + fmt.Sprintln(PrintAction(act))
		}
	}
	str = str + PrintTask(rule.Task)
	return str
}
