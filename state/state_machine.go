package state

import (
	"github.com/oneliang/util-golang/common"
)

type StateMachine[K comparable] struct {
	CurrentState *State[K]
}

func NewStateMachine[K comparable](currentState *State[K]) *StateMachine[K] {
	if common.CheckNotNil(currentState) != nil {
		return nil
	}
	return &StateMachine[K]{
		CurrentState: currentState,
	}
}
func (this *StateMachine[K]) HasPreviousState() bool {
	return this.CurrentState.HasPrevious()
}

func (this *StateMachine[K]) HasNextState() bool {
	return this.CurrentState.HasNext()
}

func (this *StateMachine[K]) PreviousState(key K) error {
	state, err := this.CurrentState.Previous(key)
	if err != nil {
		return err
	}
	this.CurrentState = state
	return nil
}

func (this *StateMachine[K]) NextState(key K) error {
	state, err := this.CurrentState.Next(key)
	if err != nil {
		return err
	}
	this.CurrentState = state
	return nil
}
