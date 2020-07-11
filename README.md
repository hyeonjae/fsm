# fsm

[![Go](https://github.com/hyeonjae/fsm/workflows/Go/badge.svg)](https://github.com/hyeonjae/fsm/workflows/Go/badge.svg) [![GoReport](https://goreportcard.com/badge/github.com/hyeonjae/fsm)](https://goreportcard.com/badge/github.com/hyeonjae/fsm) [![GoDoc](https://godoc.org/github.com/hyeonjae/fsm?status.png)](https://godoc.org/github.com/hyeonjae/fsm)



## Install

```
go get github.com/hyeonjae/fsm
```

## Usage

```go
package main

import "github.com/hyeonjae/fsm"

func main() {
	s1 := fsm.State{State: "s1"}
	s2 := fsm.State{State: "s2"}
	s3 := fsm.State{State: "s3"}

	a1 := fsm.Action{Action: "a"}
	a2 := fsm.Action{Action: "b"}

	t1 := fsm.Transition{
		Tuple: fsm.StateActionTuple{State: s1, Action: a1},
		State: s2,
	}
	t2 := fsm.Transition{
		Tuple: fsm.StateActionTuple{State: s1, Action: a2},
		State: s1,
	}
	t3 := fsm.Transition{
		Tuple: fsm.StateActionTuple{State: s2, Action: a1},
		State: s1,
	}
	t4 := fsm.Transition{
		Tuple: fsm.StateActionTuple{State: s2, Action: a2},
		State: s3,
	}

	builder := fsm.Builder().
		States([]fsm.State{s1, s2, s3}).
		Actions([]fsm.Action{a1, a2}).
		Start(s1).
		Accepts([]fsm.State{s3}).
		Transitions([]fsm.Transition{t1, t2, t3, t4})

	statemachine := fsm.New(builder)

	statemachine.Start()
  
	statemachine.Transit(a1)
	statemachine.Transit(a1)
	statemachine.Transit(a1)
	statemachine.Transit(a2)
  
	statemachine.Current()  // s3
	statemachine.Accepted() // true
}
```

## Test

```shell
$ go test ./...
ok  	github.com/hyeonjae/fsm	0.758s
```
