package main

import (
	"flag"
	"fmt"
	"steel-lang/datastructure"
	"steel-lang/semantics"
)

var (
	nodeState semantics.State
	rules []datastructure.Rule
)

func init() {
	flag.Parse()
	nodeState.Memory = make(map[string]interface{})
}

func main() {
	// Node main behaviour
	nodeBehaviour()
}

func nodeBehaviour() {
	// init nodeState
	nodeState.Memory["x"] = 1
	nodeState.Memory["y"] = "3"
	nodeState.Pool = append(nodeState.Pool, []semantics.SemanticAction{{Resource: "x", Value: 4},{Resource: "y", Value: "s"}})
	nodeState.Pool = append(nodeState.Pool, []semantics.SemanticAction{{Resource: "z", Value: true}})
	// init rules
	rules = []datastructure.Rule{{Name: "R1", Event: []string{"x","y"}, DefaultActions: nil, Task: datastructure.Task{Mode: "for", Exp: "x > 0", Actions: []datastructure.Action{{Resource: "y", External: false, Expression: "x+3"}}}},{Name: "R2", Event: []string{"x","y"}, DefaultActions: nil, Task: datastructure.Task{Mode: "for", Exp: "x < 0", Actions: []datastructure.Action{{Resource: "x", External: false, Expression: "0"}}}}}
	// exec
	intp := semantics.NewmSteelExecuter(&nodeState,rules)
	fmt.Println("Rules:\n")
	for _, rule := range rules {
		fmt.Println(datastructure.PrintRule(rule))
	}
	fmt.Println(semantics.PrintState(intp.GetState()))
	intp.Exec()
	fmt.Println()
	fmt.Println(semantics.PrintState(intp.GetState()))
}

