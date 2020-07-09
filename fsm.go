package fsm

import "errors"

type StateMachine interface {
	Start()
	Transit(Action) error

	Current() State
	Accepted() bool
}

type stateMachine struct {
	StateMachine stateMachineComprehension
	CurrentState State
}

func (machine *stateMachine) Start() {
	machine.CurrentState = machine.StateMachine.start
}

func (machine *stateMachine) Transit(a Action) error {
	s := machine.CurrentState
	next, exists := machine.StateMachine.transitions[StateActionTuple{s, a}]
	if !exists {
		return errors.New("No transition")
	}

	machine.CurrentState = next
	return nil
}

func (machine stateMachine) Current() State {
	return machine.CurrentState
}

func (machine stateMachine) Accepted() bool {
	for _, accept := range machine.StateMachine.accepts {
		if machine.CurrentState == accept {
			return true
		}
	}
	return false
}

func New(builder StateMachineBuilder) StateMachine {
	return &stateMachine{
		StateMachine: builder.Build(),
		CurrentState: State{},
	}
}
