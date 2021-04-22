package main

import (
	"fmt"
	"steel-lang/antlr/ecaruleParser"
	"steel-lang/datastructure"
	"steel-lang/semantics"
)

func main() {
	// Node main behaviour
	nodeBehaviour()
}

func nodeBehaviour() {
	// init rules
	r1intp := ecaruleParser.NewpruleIntp("rule R1 on x;y; for x > 0 do y = x+3;")
	r2intp := ecaruleParser.NewpruleIntp("rule R2 on x;y; for x < 0 do x = 0;")
	rules := []datastructure.Rule{r1intp.RunpruleIntp(), r2intp.RunpruleIntp()}
	// init nodeState
	memory := make(map[string]interface{})
	memory["x"] = 1
	memory["y"] = "3"
	pool := make([][]semantics.SemanticAction, 0)
	pool = append(pool, []semantics.SemanticAction{{Resource: "x", Value: 4}, {Resource: "y", Value: "s"}})
	pool = append(pool, []semantics.SemanticAction{{Resource: "z", Value: true}})
	// exec
	intp := semantics.NewMuSteelExecuter(memory, pool, rules)
	fmt.Println("Rules:\n")
	for _, rule := range rules {
		fmt.Println(datastructure.PrintRule(rule))
	}
	fmt.Println(intp.PrintState())
	// intp.Input([]semantics.SemanticAction{{Resource: "x", Value: 4},{Resource: "y", Value: "f"}})
	intp.Exec()
	fmt.Println()
	fmt.Println(intp.PrintState())
}
