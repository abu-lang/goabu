# GoAbU

[![Go Reference](https://pkg.go.dev/badge/github.com/abu-lang/goabu.svg)](https://pkg.go.dev/github.com/abu-lang/goabu)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/abu-lang/goabu/blob/main/LICENSE)

Golang implementation of the AbU calculus. 

The theoretical foundation of GoAbU has been presented in the peer-reviewed publication:
> Marino Miculan and Michele Pasqua. "A Calculus for Attribute-Based Memory Updates". In Antonio Cerone and Peter Ölveczky, editors, *Proceedings of the 18th international colloquium on theoretical aspects of computing, ICTAC 2021*, volume 12819 of *Lecture Notes in Computer Science*. Springer, 2021. 

You can access the pubblication on the Publisher website (here is the [DOI](http://dx.doi.org/10.1007/978-3-030-85315-0_21)).

This project makes use of:
- some packages, the ANTLR v4 grammar and the parser from the Grule rule engine, released on [github](https://github.com/hyperjumptech/grule-rule-engine) and [licensed](https://raw.githubusercontent.com/hyperjumptech/grule-rule-engine/master/LICENSE.txt) by hyperjump.tech under the [Apache 2.0 License](https://www.apache.org/licenses/LICENSE-2.0)
- the ANTLR v4 parser generator, released on [github](https://github.com/antlr/antlr4) and [licensed](https://raw.githubusercontent.com/antlr/antlr4/master/LICENSE.txt) by The ANTLR Project under the BSD-3-Clause License
- the memberlist library, released on [github](https://github.com/hashicorp/memberlist) and [licensed](https://raw.githubusercontent.com/hashicorp/memberlist/master/LICENSE) by HashiCorp under the [Mozilla Public License 2.0](https://www.mozilla.org/en-US/MPL/2.0/)
- the uuid package, released on [github](https://github.com/google/uuid) and [licensed](https://raw.githubusercontent.com/google/uuid/master/LICENSE) by Google under the BSD-3-Clause License
- the Gobot framework, released on [github](https://github.com/hybridgroup/gobot/) and [licensed](https://raw.githubusercontent.com/hybridgroup/gobot/release/LICENSE.txt) by The Hybrid Group under the [Apache 2.0 License](https://www.apache.org/licenses/LICENSE-2.0)
- the zap logger, released on [github](https://github.com/uber-go/zap) and [licensed](https://raw.githubusercontent.com/uber-go/zap/master/LICENSE.txt) by Uber Technologies, Inc. under the MIT License

# Installation

GoAbU can be retrieved with go get:

```Shell
$ go get github.com/abu-lang/goabu
```

## Simulator

Try GoAbU on our Docker-based [simulator](https://github.com/abu-lang/abusim).

# Quick Start

## Import GoAbU

```go
import (
	"github.com/abu-lang/goabu"
	"github.com/abu-lang/goabu/communication"
	"github.com/abu-lang/goabu/config"
	"github.com/abu-lang/goabu/memory"
)
```

## Creating a Resources struct

```go
package memory

type Resources struct {
	Bool    map[string]bool
	Integer map[string]int64
	Float   map[string]float64
	Text    map[string]string
	Time    map[string]time.Time
	Other   map[string]interface{}
}
```

memory.Resources is a struct constituted by maps that will contain the resources used by the node.

The function memory.MakeResources() can be used to initialize all the fields with empty maps.
Then the needed resources can be initializated as needed:

```go
mem := memory.MakeResources()
mem.Integer["foo"] = 1
mem.Text["bar"] = "octocat"
```

**NOTE** that the names of the resources (aka the map keys) should adhere to the standard syntax for identifiers and also that the subsequent case insensitive keywords are reserved: this, ext, rule, when, then, true, false, nil, salience, on, default, for, all, do.

## GoAbU Rules

In the ECA rules of GoAbU, the Condition and the Actions are encoded in a task part starting with the "for" keyword.
A task can be local (encoding local actions) or remote (encoding global actions):
- rules with local tasks are like standard ECA rules and can influence only the the current node
- remote tasks specify global actions which are performed on all the other nodes apart from the current one

### Local Tasks

```go
localRule := `rule MyLocalRule on foo bar for "octocat" == bar do foo = foo * 2, bar = "gopher"`
```
This rule specifies that whenever the values of foo or bar change then if bar == "octocat" foo shuld be doubled and bar should take the value "gopher".

### Global Tasks

```go
globalRule := `rule MyGlobalRule on foo for all this.foo >= ext.foo do ext.foo = ext.foo + this.foo`
```

**NOTE** that the keyword **all** is used to distinguish remote tasks from local ones.

This rule specifies that when the value of the (local) resource foo changes then some update should be performed on all the other nodes that have foo which is less or equal than the value of foo on the current node.
In particular these nodes should change their foo with the sum of their value of foo with the value of foo from the node that fired the rule.

**NOTE** that to distinguish between local and remote resources the prefixes "this." and "ext." are used.
This can be a little verbose but on every assignment LHS the resource type can be inferred and if no prefix is specified then it is assumed that "this." was the intended one.

So we can simplify a little bit the previous rule:

```go
globalRule = `rule MyGlobalRule on foo for all foo >= ext.foo do foo = ext.foo + foo`
```

This also explain why rules with local tasks, as the one seen before, do not require prefixes.

## Creating an Agent

To perform the communication required by the remote task we have to create an Agent which is an interface that abstracts the communication between the various nodes.

Currently the package communication has an implementation called MemberlistAgent based on [memberlist](https://github.com/hashicorp/memberlist).

A MemberlistAgent can be created by the function NewMemberlistAgent which takes an identifier for the Agent, an int that specifies the listening port and optionally a variadic list of strings of the type "host:port" that indicate the other MemberlistAgents to join:

```go
agent := communication.NewMemberlistAgent("Agent", 5000, config.LogConfig{})
```

## Creating the Executer

Finally we are ready to start our node.
A node is represented by an Executer that will contain the Resources struct, a knowledge base of GoAbU rules and an Agent.
The Executer specifies the ECA rule execution model. It uses the knowledge base to apply the required updates to the resources and to send updates to the other nodes by relying on the Agent for the communication.

The Executer can be constructed using the NewExecuter function:

**NOTE** that for simplicity in the tutorial we will not check for returned errors, when using GoAbU errors should be checked.

```go
executer, _ := goabu.NewExecuter(mem, []string{localRule}, agent, config.LogConfig{})
```

The function NewExecuter also starts the Agent and performs the join operation.

**NOTE** that the resources of mem are copied inside the Executer by means of the method mem.Copy() but for the elements of mem.Other only a shallow copy is performed.
So an external synchronization may be required.

## Another Local Node

Let's make another Executer to make thing livelier, for simplicity we will create it locally.

We simply repeat the previous steps with some modifications:

```go
mem2 := memory.MakeResources()
mem2.Integer["foo"] = 1
mem2.Float["baz"] = 3.14

agent2 := communication.NewMemberlistAgent("Agent-2", 5001, config.LogConfig{}, "localhost:5000")

executer2, _ := goabu.NewExecuter(mem2, []string{globalRule}, agent2, config.LogConfig{})
```

## Input

Now we have our local cluster with two nodes but the situation is still the same as no resource changed and consequently no rules were fired.

We can change the resource values using the Input method as follow:

```go
executer2.Input("foo = 3, baz = 2.72")
```

Now we changed the resources of executer2 but actually no modification happened on the other Executer.
The fact is that when a rule is fired its changes are evaluated but aren't applied immediately.
The changes are grouped in an atomic Update (goabu.Update) and appended to a pool of the relative Executer.

So the changes implied by MyGLobalRule are currently in the pool owned by executer.

We can take an Update from the pull and perform its changes by means of the method Exec:

```go
executer.Exec()
executer.Exec()
```

We call Exec two times to also apply the changes deriving from MyLocalRule.

## Inspecting the State

To access the values of the resources we can use the method TakeState().
TakeState() returns a copy of the Executer's Resources struct and a copy of its Update pool.

```go
state, _ := executer.TakeState()
fmt.Println("foo =", state.Integer["foo"])
fmt.Println("bar =", state.Text["bar"])
state2, _ := executer2.TakeState()
fmt.Println("foo =", state2.Integer["foo"])
fmt.Println("baz =", state2.Float["baz"])
```

# Input/Output Resources

Apart from normal resources GoAbU also has Input/Output resources that can map and reflect the state of GPIO sensors and actuators.

The struct IOresources defined in the package physical generalizes the Resources struct and also permits the use of Input/Output resources by relying on the [Gobot](https://gobot.io/) framework.

To initialize the struct it is sufficient to call the MakeIOresources constructor providing a gobot.Adaptor that implements the physical.IOadaptor interface.

For example on a Raspberry Pi:

```go
import (
	"github.com/abu-lang/goabu"
	"github.com/abu-lang/goabu/communication"
	"github.com/abu-lang/goabu/config"
	"github.com/abu-lang/goabu/physical"
	"github.com/abu-lang/goabu/physical/iodelegates"

	"gobot.io/x/gobot/platforms/raspi"
)

mem := iodelegates.MakeIOresources(raspi.NewAdaptor())
```

Then as IOresources embeds a Resources struct we can add normal resources as before but also add Input/Output resources by specifying the GPIO pins as aguments to the Add method:

```go
mem.Integer["myint"] = 0

mem.Add("DigitalPin", "led", "36")
mem.Add("Button", "button1", "38")
mem.Add("Button", "button2", "40")
mem.Add("Motor", "motor", "13", "11")
```

mem can then be used as the first argument to the NewExecuter constructor.

The currently supported sensors/actuators are digital output pins, motors and buttons.
But to add and use other devices it is sufficient to implement the physical.IOdelegate interface.

# Appendix

## Default Actions

The rules of GoAbu can also have some default actions that are performed when the rule is activated regardless of the rule's condition.

```go
r := `rule R on foo default baz = 0.0, bar = "octocat" for all ext.foo < 0 do foo = ext.foo * -1`
```

**NOTE** that default actions are **always** performed on the current node and can access only local resources.

## Rules with multiple tasks

A rule is not limited to have a single task but can have multiple tasks.

For example, the rule R of the previous section can also be encoded with an equivalent rule using two tasks:
```go
r := `rule R on foo for all ext.foo < 0 do foo = ext.foo * -1 for true do baz = 0.0, bar = "octocat"`
```

## Invariants

An Executer can have some invariants that indicate the correct states of its resources.
In particular if a call to Exec selects an update (discovered locally or received from another node) that would violate the invariants then that update is removed from the pool but no resource is modified.

These invariants can be specified as optional arguments upon the Executer's construction:

```go
executer, err := goabu.NewExecuter(mem, []string{localRule}, agent, config.LogConfig{},
	"foo > -273", "bar == \"octocat\" || bar == \"gopher\"")
```

## Full Example

```go
package main

import (
	"fmt"

	"github.com/abu-lang/goabu"
	"github.com/abu-lang/goabu/communication"
	"github.com/abu-lang/goabu/config"
	"github.com/abu-lang/goabu/memory"
)

func main() {
	mem := memory.MakeResources()
	mem.Integer["foo"] = 1
	mem.Text["bar"] = "octocat"

	localRule := `rule MyLocalRule on foo bar for "octocat" == bar do foo = foo * 2, bar = "gopher"`

	globalRule := `rule MyGlobalRule on foo for all this.foo >= ext.foo do ext.foo = ext.foo + this.foo`
	globalRule = `rule MyGlobalRule on foo for all foo >= ext.foo do foo = ext.foo + foo`

	agent := communication.NewMemberlistAgent("Agent", 5000, config.LogConfig{})

	executer, _ := goabu.NewExecuter(mem, []string{localRule}, agent, config.LogConfig{})

	mem2 := memory.MakeResources()
	mem2.Integer["foo"] = 1
	mem2.Float["baz"] = 3.14

	agent2 := communication.NewMemberlistAgent("Agent-2", 5001, config.LogConfig{}, "localhost:5000")

	executer2, _ := goabu.NewExecuter(mem2, []string{globalRule}, agent2, config.LogConfig{})

	executer2.Input("foo = 3, baz = 2.72")

	executer.Exec()
	executer.Exec()

	state, _ := executer.TakeState()
	fmt.Println("foo =", state.Integer["foo"])
	fmt.Println("bar =", state.Text["bar"])
	state2, _ := executer2.TakeState()
	fmt.Println("foo =", state2.Integer["foo"])
	fmt.Println("baz =", state2.Float["baz"])
}
```
