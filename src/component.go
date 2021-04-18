package main

import (
	"fmt"
	"steel-lang/datastructure"

	"github.com/hyperjumptech/grule-rule-engine/ast"
)

type ruleData struct {
	Resources map[string]int64
}

type Component struct {
	localMem__ ruleData
	changed    datastructure.StringSet
	library    map[string]*datastructure.RuleDict

	knowledgeLibrary *ast.KnowledgeLibrary
}

func NewComponent(atts map[string]int64) *Component {
	return &Component{
		localMem__: ruleData{
			Resources: atts,
		},
		changed:          make(map[string]bool),
		library:          make(map[string]*datastructure.RuleDict),
		knowledgeLibrary: ast.NewKnowledgeLibrary(),
	}
}

func (cmp *Component) AddRule(rule *datastructure.Rule) {
	parsed := datastructure.NewParsedRule(rule, cmp.knowledgeLibrary)
	for _, evt := range parsed.Event {
		if cmp.library[evt] == nil {
			var dict datastructure.RuleDict = make(map[string]*datastructure.ParsedRule)
			cmp.library[evt] = &dict
		}
		cmp.library[evt].Insert(parsed)
	}
}

func (cmp *Component) fire() {
	dataContext := ast.NewDataContext()
	err := dataContext.Add("this", &(cmp.localMem__))
	if err != nil {
		panic(err)
	}
	knowledgeBase := cmp.knowledgeLibrary.NewKnowledgeBaseInstance("dummy", "0.0.0")
	defunc := &ast.BuiltInFunctions{
		Knowledge:     knowledgeBase,
		WorkingMemory: knowledgeBase.WorkingMemory,
		DataContext:   dataContext,
	}
	dataContext.Add("DEFUNC", defunc)
	knowledgeBase.WorkingMemory.ResetAll()
	knowledgeBase.InitializeContext(dataContext)

	for !cmp.changed.Empty() {
		rules := cmp.activeRules()
		var modified datastructure.StringSet = make(map[string]bool)
		for _, rule := range *rules {
			if rule.DefaultActions != nil {
				for _, action := range rule.DefaultActions {
					modified.Insert(action.Resource)
					assignment := action.Expression
					variable := assignment.Variable
					variable = knowledgeBase.WorkingMemory.AddVariable(variable)
					rexpr := assignment.Expression
					rexpr = knowledgeBase.WorkingMemory.AddExpression(rexpr)
					exprVal, err := rexpr.Evaluate(dataContext, knowledgeBase.WorkingMemory)
					if err != nil {
						panic(err)
					}
					variable.Assign(exprVal, dataContext, knowledgeBase.WorkingMemory)
				}
			}
			exp := rule.Task.Exp
			exp = knowledgeBase.WorkingMemory.AddExpression(exp)
			val, err := exp.Evaluate(dataContext, knowledgeBase.WorkingMemory)
			if err != nil {
				panic(err)
			}
			can := val.Bool()
			if can {
				if rule.Task.Mode == "for" {
					for _, action := range rule.Task.Actions {
						modified.Insert(action.Resource)
						assignment := action.Expression
						variable := assignment.Variable
						variable = knowledgeBase.WorkingMemory.AddVariable(variable)
						rexpr := assignment.Expression
						rexpr = knowledgeBase.WorkingMemory.AddExpression(rexpr)
						exprVal, err := rexpr.Evaluate(dataContext, knowledgeBase.WorkingMemory)
						if err != nil {
							panic(err)
						}
						variable.Assign(exprVal, dataContext, knowledgeBase.WorkingMemory)
					}
				}
			}
		}
		cmp.changed = modified
	}
}

func (cmp *Component) activeRules() *datastructure.RuleDict {
	var dict datastructure.RuleDict = make(map[string]*datastructure.ParsedRule)
	var res *datastructure.RuleDict = &dict
	for e := range cmp.changed {
		res.Add(cmp.library[e])
	}
	return res
}

func (cmp *Component) Get(att string) int64 {
	return cmp.localMem__.Resources[att]
}

func (cmp *Component) Set(att string, val int64) {
	cmp.localMem__.Resources[att] = val
	cmp.changed.Insert(att)
	cmp.fire()
}

// usage example
func main() {
	atts := make(map[string]int64)
	atts["x"] = 1
	atts["y"] = 3
	cmp := NewComponent(atts)
	r1 := datastructure.Rule{
		Name:           "R1",
		Event:          []string{"x", "y"},
		DefaultActions: nil,
		Task: datastructure.Task{
			Mode: "for",
			Exp:  `this.Resources["x"] > 0`,
			Actions: []datastructure.Action{
				{Resource: "y",
					External:   false,
					Expression: `this.Resources["x"] + 3`,
				},
			},
		},
	}
	r2 := datastructure.Rule{
		Name:           "R2",
		Event:          []string{"x", "y"},
		DefaultActions: nil,
		Task: datastructure.Task{
			Mode: "for",
			Exp:  `this.Resources["x"] > 0`,
			Actions: []datastructure.Action{
				{Resource: "x",
					External:   false,
					Expression: "0",
				},
			},
		},
	}
	cmp.AddRule(&r1)
	cmp.AddRule(&r2)
	cmp.Set("y", 3)
	fmt.Println(cmp.Get("x"))
	fmt.Println(cmp.Get("y"))
}
