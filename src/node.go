package main

import (
	"flag"
	"fmt"
	"steel-lang/communication"
	"steel-lang/datastructure"
	"steel-lang/semantics"
	"strings"
	"time"
)

var (
	port  = flag.Int("port", 0, "listening port")
	nodes = flag.String("nodes", "", "list of nodes to join (e.g. '127.0.0.1:8000,127:0.0.1:8001')")
)

func main() {
	flag.Parse()
	// Node main behaviour
	nodeBehaviour()
}

func nodeBehaviour() {
	// init rules
	//r1intp := ecaruleParser.NewpruleIntp("rule R1 on x;y; for x > 0 do y = x+3;")
	//r2intp := ecaruleParser.NewpruleIntp("rule R2 on x;y; for x < 0 do x = 0;")
	r1 := datastructure.Rule{
		Name:           "R1",
		Events:         []string{"x", "y"},
		DefaultActions: nil,
		Task: datastructure.Task{
			Mode:      "for",
			Condition: `this.Integer["x"] > 0`,
			Actions: []datastructure.Action{
				{Resource: "y",
					Expression: `"42"`,
				},
			},
		},
	}
	r2 := datastructure.Rule{
		Name:           "R2",
		Events:         []string{"x", "y"},
		DefaultActions: nil,
		Task: datastructure.Task{
			Mode:      "for",
			Condition: `this.Integer["x"] > 0`,
			Actions: []datastructure.Action{
				{Resource: "x",
					Expression: "0",
				},
			},
		},
	}
	r3 := datastructure.Rule{
		Name:           "R3",
		Events:         []string{"x"},
		DefaultActions: nil,
		Task: datastructure.Task{
			Mode:      "for all",
			Condition: `this.Integer["x"] > ext.Integer["x"] - 1`,
			Actions: []datastructure.Action{
				{Resource: "s",
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
	var initialNodes []string
	if len(*nodes) > 0 {
		initialNodes = strings.Split(*nodes, ",")
	}
	intp, err := semantics.NewMuSteelExecuter(memory, nil, communication.MakeMemberlistAgent(memory.ResourceNames(), *port, initialNodes))
	if err != nil {
		panic(err)
	}
	intp.AddRules(rules)
	intp.AddPool(pool)
	fmt.Print("Rules:\n\n")
	for _, rule := range rules {
		fmt.Println(rule)
	}
	fmt.Println(intp.PrintState())
	// intp.Input([]datastructure.Action{{Resource: "x", Expression: "4"}, {Resource: "y", Expression: `"f"`}})
	intp.Exec()
	fmt.Println()
	fmt.Println(intp.PrintState())
	fmt.Println()
	fmt.Println(intp.PrintState())
	intp.Exec()
	fmt.Println()
	fmt.Println(intp.PrintState())
	time.Sleep(60 * time.Second) // for memberlist testing
	err = intp.StopAgent()
	if err != nil {
		panic(err)
	}
	err = intp.StartAgent()
	if err != nil {
		panic(err)
	}
	time.Sleep(60 * time.Second) // for memberlist testing
}
