// +build raspi

package main_test

import (
	"runtime"
	"steel-lang/communication"
	"steel-lang/config"
	"steel-lang/physical"
	"steel-lang/physical/iodelegates"
	"steel-lang/semantics"
	"testing"
	"time"

	"gobot.io/x/gobot/platforms/raspi"
)

func TestLed2Buttons(t *testing.T) {
	toggles := 6
	var a physical.IOadaptor = raspi.NewAdaptor()
	memLed := iodelegates.MakeIOresources(a)
	memButtons := iodelegates.MakeIOresources(a)
	memLed.Add("DigitalPin", "led", "36")
	memButtons.Add("Button", "button1", "38")
	memButtons.Add("Button", "button2", "40")
	r1 := "rule R1 on button2 for all this.button1 && this.button2 do ext.led = !ext.led"
	eLed, err := semantics.NewExecuter(memLed, nil, communication.NewMemberlistAgent(8100, config.TestsLogConfig), config.TestsLogConfig)
	if err != nil {
		t.Fatal(err)
	}
	eLed.SetOptimisticExec(*optimistic)
	eLed.SetOptimisticInput(*optimistic)
	dummy, err := semantics.NewExecuter(memButtons, []string{r1}, communication.NewMemberlistAgent(8101, config.TestsLogConfig, "127.0.0.1:8100"), config.TestsLogConfig)
	if err != nil {
		t.Fatal(err)
	}
	dummy.SetOptimisticExec(*optimistic)
	dummy.SetOptimisticInput(*optimistic)
	ledStatus := eLed.TakeState().Memory.Bool["led"]
	for toggles > 0 {
		time.Sleep(time.Millisecond)
		eLed.Exec()
		status := eLed.TakeState().Memory.Bool["led"]
		if ledStatus != status {
			ledStatus = status
			toggles--
		}
	}
	runtime.KeepAlive(dummy)
}

func TestMotor(t *testing.T) {
	var a physical.IOadaptor = raspi.NewAdaptor()
	mem := iodelegates.MakeIOresources(a)
	mem.Add("Motor", "motor", "13", "11")
	r1 := "rule R1 on motor for this.motor > 0 && this.motor < 255 do motor = this.motor + 60"
	r2 := "rule R2 on motor for this.motor >= 255 do motor = 0;"
	e, err := semantics.NewExecuter(mem, []string{r1, r2}, semantics.MakeMockAgent(), config.TestsLogConfig)
	if err != nil {
		t.Fatal(err)
	}
	e.SetOptimisticExec(*optimistic)
	e.SetOptimisticInput(*optimistic)
	e.Input("motor = -150")
	time.Sleep(8 * time.Second)
	e.Input("motor = 150;")
	for {
		time.Sleep(8 * time.Second)
		e.Exec()
		if e.TakeState().Memory.Integer["motor"] == 0 {
			break
		}
	}
}
