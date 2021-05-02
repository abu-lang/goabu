package main

import (
	"fmt"
	"steel-lang/datastructure"
	"steel-lang/semantics"
)

func main() {
	// Node main behaviour
	nodeBehaviour()
}

func nodeBehaviour() {
	// init rules
	//r1intp := ecaruleParser.NewpruleIntp("rule R1 on x;y; for x > 0 do y = x+3;")
	//r2intp := ecaruleParser.NewpruleIntp("rule R2 on x;y; for x < 0 do x = 0;")
	r1 := datastructure.Rule{
		Name:           "R1",
		Event:          []string{"x", "y"},
		DefaultActions: nil,
		Task: datastructure.Task{
			Mode: "for",
			Exp:  `this.Integer["x"] > 0`,
			Actions: []datastructure.Action{
				{Resource: "y",
					External:   false,
					Expression: `"42"`,
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
			Exp:  `this.Integer["x"] > 0`,
			Actions: []datastructure.Action{
				{Resource: "x",
					External:   false,
					Expression: "0",
				},
			},
		},
	}
	r3 := datastructure.Rule{
		Name:           "R3",
		Event:          []string{"x"},
		DefaultActions: nil,
		Task: datastructure.Task{
			Mode: "for all",
			Exp:  `this.Integer["x"] > ext.Integer["x"] - 1`,
			Actions: []datastructure.Action{
				{Resource: "s",
					External:   false,
					Expression: `ext.Other[this.Text["y"].ToUpper()]`,
				},
			},
		},
	}
	rules := []datastructure.Rule{r1, r2, r3}
	// init nodeState
	memory := datastructure.MakeResources()
	memory.Integer["x"] = 1
	memory.Text["y"] = "3"
	memory.Bool["z"] = false
	memory.Other["S"] = r3
	memory.Other["s"] = false
	pool := make([][]datastructure.Action, 0)
	pool = append(pool, []datastructure.Action{{Resource: "x", Expression: "4"}, {Resource: "y", Expression: `"s"`}})
	pool = append(pool, []datastructure.Action{{Resource: "z", Expression: "true"}})
	// exec
	intp, err := semantics.NewMuSteelExecuter(memory)
	intp.AddRules(rules)
	intp.AddPool(pool)
	if err != nil {
		panic(err)
	}
	fmt.Println("Rules:\n")
	for _, rule := range rules {
		fmt.Println(datastructure.PrintRule(rule))
	}
	fmt.Println(intp.PrintState())
	// intp.Input([]datastructure.Action{{Resource: "x", Expression: "4"}, {Resource: "y", Expression: `"f"`}})
	intp.Exec()
	fmt.Println()
	fmt.Println(intp.PrintState())
	intp.TestExtPool()
	fmt.Println()
	fmt.Println(intp.PrintState())
}
