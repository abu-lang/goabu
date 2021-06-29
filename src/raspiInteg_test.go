// +build raspi

package main_test

import (
	"steel-lang/datastructure"
	"steel-lang/physical"
	"steel-lang/semantics"
	"testing"
	"time"

	"gobot.io/x/gobot/platforms/raspi"
)

func TestLed2Buttons(t *testing.T) {
	toggles := 6
	var a physical.IOAdaptor = raspi.NewAdaptor()
	memory := physical.MakeIOResources(a)
	memory.AddLed("led", "12")
	memory.AddButton("button1", "16")
	memory.AddButton("button2", "18")
	r1 := datastructure.Rule{
		Name:           "R1",
		Events:         []string{"button2"},
		DefaultActions: nil,
		Task: datastructure.Task{
			Mode:      "for all",
			Condition: `this.Bool["button1"] && this.Bool["button2"]`,
			Actions: []datastructure.Action{
				{Resource: "led",
					Expression: `!this.Bool["led"]`,
				},
			},
		},
	}
	e, err := semantics.NewMuSteelExecuter(memory, []datastructure.Rule{r1}, semantics.MakeMockAgent())
	if err != nil {
		t.Fatal(err)
	}
	ledStatus := e.GetState().Memory.Bool["led"]
	for toggles > 0 {
		time.Sleep(time.Millisecond)
		e.Exec()
		status := e.GetState().Memory.Bool["led"]
		if ledStatus != status {
			ledStatus = status
			toggles--
		}
	}
}
