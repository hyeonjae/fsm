package fsm

type StateMachineBuilder interface {
	States([]State) StateMachineBuilder
	Actions([]Action) StateMachineBuilder
	Start(State) StateMachineBuilder
	Accepts([]State) StateMachineBuilder
	Transitions([]Transition) StateMachineBuilder

	Build() stateMachineComprehension
}

type stateMachineBuilder struct {
	states      []State
	actions     []Action
	start       State
	accepts     []State
	transitions []Transition
}

func Builder() *stateMachineBuilder {
	return &stateMachineBuilder{}
}

type stateMachineComprehension struct {
	states      []State
	actions     []Action
	start       State
	accepts     []State
	transitions map[StateActionTuple]State
}

func (builder *stateMachineBuilder) States(states []State) StateMachineBuilder {
	builder.states = states
	return builder
}

func (builder *stateMachineBuilder) Actions(actions []Action) StateMachineBuilder {
	builder.actions = actions
	return builder
}

func (builder *stateMachineBuilder) Start(start State) StateMachineBuilder {
	builder.start = start
	return builder
}

func (builder *stateMachineBuilder) Accepts(accepts []State) StateMachineBuilder {
	builder.accepts = accepts
	return builder
}

func (builder *stateMachineBuilder) Transitions(transitions []Transition) StateMachineBuilder {
	builder.transitions = transitions
	return builder
}

func (builder *stateMachineBuilder) Build() stateMachineComprehension {
	transitions := map[StateActionTuple]State{}
	for _, transition := range builder.transitions {
		transitions[transition.Tuple] = transition.State
	}

	return stateMachineComprehension{
		states:      builder.states,
		actions:     builder.actions,
		start:       builder.start,
		accepts:     builder.accepts,
		transitions: transitions,
	}
}
