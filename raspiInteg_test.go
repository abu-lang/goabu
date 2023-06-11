// Copyright 2021 Massimo Comuzzo, Michele Pasqua and Marino Miculan
// SPDX-License-Identifier: Apache-2.0

//go:build raspi
// +build raspi

package goabu_test

import (
	"runtime"
	"testing"
	"time"

	"github.com/abu-lang/goabu"
	"github.com/abu-lang/goabu/communication"
	"github.com/abu-lang/goabu/config"
	"github.com/abu-lang/goabu/physical"
	"github.com/abu-lang/goabu/physical/iodelegates"

	"gobot.io/x/gobot/v2/platforms/raspi"
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
	eLed, err := goabu.NewExecuter(memLed, nil, communication.NewMemberlistAgent("Led", 8100, config.TestsLogConfig), config.TestsLogConfig)
	if err != nil {
		t.Fatal(err)
	}
	eLed.SetOptimisticExec(*goabu.Optimistic)
	eLed.SetOptimisticInput(*goabu.Optimistic)
	dummy, err := goabu.NewExecuter(memButtons, []string{r1}, communication.NewMemberlistAgent("Buttons", 8101, config.TestsLogConfig, "127.0.0.1:8100"), config.TestsLogConfig)
	if err != nil {
		t.Fatal(err)
	}
	dummy.SetOptimisticExec(*goabu.Optimistic)
	dummy.SetOptimisticInput(*goabu.Optimistic)

	state, _ := eLed.TakeState()
	ledStatus, _ := state.Bool["led"]
	for toggles > 0 {
		time.Sleep(time.Millisecond)
		eLed.Exec()
		state, _ = eLed.TakeState()
		status := state.Bool["led"]
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
	r2 := "rule R2 on motor for this.motor >= 255 do motor = 0,"
	e, err := goabu.NewExecuter(mem, []string{r1, r2}, goabu.MakeMockAgent(), config.TestsLogConfig)
	if err != nil {
		t.Fatal(err)
	}
	e.SetOptimisticExec(*goabu.Optimistic)
	e.SetOptimisticInput(*goabu.Optimistic)
	e.Input("motor = -150")
	time.Sleep(8 * time.Second)
	e.Input("motor = 150,")
	for {
		time.Sleep(8 * time.Second)
		e.Exec()
		if state, _ := e.TakeState(); state.Integer["motor"] == 0 {
			break
		}
	}
}
