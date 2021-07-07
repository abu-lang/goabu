// +build raspi

package main_test

import (
	"steel-lang/communication"
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
	memLed := physical.MakeIOResources(a)
	memButtons := physical.MakeIOResources(a)
	memLed.AddLed("led", "12")
	memButtons.AddButton("button1", "16")
	memButtons.AddButton("button2", "18")
	r1 := datastructure.Rule{
		Name:           "R1",
		Events:         []string{"button2"},
		DefaultActions: nil,
		Task: datastructure.Task{
			Mode:      "for all",
			Condition: `this.Bool["button1"] && this.Bool["button2"]`,
			Actions: []datastructure.Action{
				{Resource: "led",
					Expression: `!ext.Void["led"]`,
				},
			},
		},
	}
	eLed, err := semantics.NewMuSteelExecuter(memLed, nil, communication.MakeMemberlistAgent(memLed.ResourceNames(), 8100, nil))
	if err != nil {
		t.Fatal(err)
	}
	dummy, err := semantics.NewMuSteelExecuter(memButtons, []datastructure.Rule{r1}, communication.MakeMemberlistAgent(memButtons.ResourceNames(), 8101, []string{"127.0.0.1:8100"}))
	if err != nil {
		t.Fatal(err)
	}
	ledStatus := eLed.GetState().Memory.Bool["led"]
	for toggles > 0 {
		time.Sleep(time.Millisecond)
		eLed.Exec()
		status := eLed.GetState().Memory.Bool["led"]
		if ledStatus != status {
			ledStatus = status
			toggles--
		}
	}
	dummy.IsStable()
}
