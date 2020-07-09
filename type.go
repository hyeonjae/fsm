package fsm

type State struct {
	State string
}

type Action struct {
	Action string
}

type StateActionTuple struct {
	State  State
	Action Action
}

type Transition struct {
	Tuple StateActionTuple
	State State
}
