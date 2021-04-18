package main

import (
	"fmt"
	"steel-lang/antlr/ecaruleParser"
	"steel-lang/datastructure"
	"steel-lang/semantics"
)

var (
	nodeState semantics.State
	rules []datastructure.Rule
)

func init() {
	nodeState.Memory = make(map[string]interface{})
}

/*
func main() {
	// Node main behaviour
	nodeBehaviour()
}
*/

func nodeBehaviour() {
	// init rules
	r1intp := ecaruleParser.NewpruleIntp("rule R1 on x;y; for x > 0 do y = x+3;")
	r2intp := ecaruleParser.NewpruleIntp("rule R2 on x;y; for x < 0 do x = 0;")
	rules = []datastructure.Rule{r1intp.RunpruleIntp(),r2intp.RunpruleIntp()}
	// init nodeState
	nodeState.Memory["x"] = 1
	nodeState.Memory["y"] = "3"
	nodeState.Pool = append(nodeState.Pool, []semantics.SemanticAction{{Resource: "x", Value: 4},{Resource: "y", Value: "s"}})
	nodeState.Pool = append(nodeState.Pool, []semantics.SemanticAction{{Resource: "z", Value: true}})
	// exec
	intp := semantics.NewmSteelExecuter(&nodeState,rules)
	fmt.Println("Rules:\n")
	for _, rule := range rules {
		fmt.Println(datastructure.PrintRule(rule))
	}
	fmt.Println(semantics.PrintState(intp.GetState()))
	// intp.Input([]semantics.SemanticAction{{Resource: "x", Value: 4},{Resource: "y", Value: "f"}})
	intp.Exec()
	fmt.Println()
	fmt.Println(semantics.PrintState(intp.GetState()))
}

