package fsm_test

import (
	"testing"

	fsm "github.com/hyeonjae/fsm"
)

func TestFSM(t *testing.T) {
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
	if statemachine.Current() != s1 {
		t.Error("Wrong result")
	}
	if statemachine.Accepted() != false {
		t.Error("Wrong result")
	}

	statemachine.Transit(a1)
	if statemachine.Current() != s2 {
		t.Error("Wrong result")
	}
	if statemachine.Accepted() != false {
		t.Error("Wrong result")
	}

	statemachine.Transit(a1)
	if statemachine.Current() != s1 {
		t.Error("Wrong result")
	}
	if statemachine.Accepted() != false {
		t.Error("Wrong result")
	}

	statemachine.Transit(a1)
	if statemachine.Current() != s2 {
		t.Error("Wrong result")
	}
	if statemachine.Accepted() != false {
		t.Error("Wrong result")
	}

	statemachine.Transit(a2)
	if statemachine.Current() != s3 {
		t.Error("Wrong result")
	}
	if statemachine.Accepted() != true {
		t.Error("Wrong result")
	}
}
