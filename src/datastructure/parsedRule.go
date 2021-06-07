// TODO redo this ugly makeshift (maybe putting together the rule antlr4 parser with the one from grule) and add support for other modifiers

package datastructure

import (
	"strconv"

	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
)

type ParsedRule struct {
	Name           string
	Events         []string
	DefaultActions []ParsedAction
	Task           ParsedTask
}

type ParsedAction struct {
	Resource   string
	Expression *ast.Assignment
}

type ParsedTask struct {
	Mode      string
	Condition *ast.Expression
	Actions   []ParsedAction
}

func (a ParsedAction) String() string {
	return a.Expression.GetGrlText()
}

func ActionsToStr(actions []ParsedAction) string {
	res := ""
	for _, action := range actions {
		res += action.String() + "; "
	}
	return res
}

func NewParsedRule(rule *Rule, kl *ast.KnowledgeLibrary, types map[string]string) *ParsedRule {
	res := &ParsedRule{
		Name:           rule.Name,
		Events:         make([]string, len(rule.Events)),
		DefaultActions: NewParsedActionList(rule.DefaultActions, rule.Name+"default", kl, types),
		Task:           NewParsedTask(&(rule.Task), rule.Name+"task", kl, types),
	}
	copy(res.Events, rule.Events)
	return res
}

func NewParsedTask(t *Task, name string, kl *ast.KnowledgeLibrary, types map[string]string) ParsedTask {
	return ParsedTask{
		Mode:      t.Mode,
		Condition: NewParsedExpression(t.Condition, name+"cnd", kl),
		Actions:   NewParsedActionList(t.Actions, name+"actions", kl, types),
	}
}

func NewParsedActionList(acts []Action, name string, kl *ast.KnowledgeLibrary, types map[string]string) []ParsedAction {
	var res []ParsedAction
	for i, a := range acts {
		res = append(res, NewParsedAction(&a, name+strconv.Itoa(i), kl, types))
	}
	return res
}

func NewParsedAction(a *Action, name string, kl *ast.KnowledgeLibrary, types map[string]string) ParsedAction {
	rb := builder.NewRuleBuilder(kl)
	rule := "rule " + name + " { when true then this." + types[a.Resource] + "[\"" + a.Resource + "\"] = " + a.Expression + "; }"
	bs := pkg.NewBytesResource([]byte(rule))
	err := rb.BuildRuleFromResource("dummy", "0.0.0", bs)
	if err != nil {
		panic(err)
	}
	kb := kl.NewKnowledgeBaseInstance("dummy", "0.0.0")
	ruleEntry := kb.RuleEntries[name]
	return ParsedAction{
		Resource:   a.Resource,
		Expression: ruleEntry.ThenScope.ThenExpressionList.ThenExpressions[0].Assignment,
	}
}

func NewParsedExpression(str, name string, kl *ast.KnowledgeLibrary) *ast.Expression {
	rb := builder.NewRuleBuilder(kl)
	rule := "rule " + name + " { when " + str + " then Ok(); }"
	bs := pkg.NewBytesResource([]byte(rule))
	err := rb.BuildRuleFromResource("dummy", "0.0.0", bs)
	if err != nil {
		panic(err)
	}
	kb := kl.NewKnowledgeBaseInstance("dummy", "0.0.0")
	ruleEntry := kb.RuleEntries[name]
	return ruleEntry.WhenScope.Expression
}

// Precondition: rule.Task.Mode != "for"
func (r *ParsedRule) PreEvaluated(dataCtx ast.IDataContext, workMem *ast.WorkingMemory) ExternalAction {
	res := ExternalAction{
		CondWorkingSet: MakeStringSet(""),
		Constants:      make(map[string]interface{}),
		IntConstants:   make(map[string]int64),
		dataContext:    dataCtx,
		workingMemory:  workMem,
	}
	res.WorkingSets = make([]StringSet, 0, len(r.Task.Actions))
	for _, action := range r.Task.Actions {
		res.WorkingSets = append(res.WorkingSets, MakeStringSet(action.Resource))
	}
	res.Condition = res.preEvaluatedExpression(r.Task.Condition, res.CondWorkingSet)
	res.Actions = res.preEvaluatedActions(r.Task.Actions)
	return res
}
