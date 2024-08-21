package state

import (
	"errors"
	"fmt"
)

type State[K comparable] struct {
	Key  K
	Name string
	//protected
	previousStateMap map[K]*State[K]
	//protected
	nextStateMap map[K]*State[K]
}

func NewState[K comparable](key K, name string) *State[K] {
	return &State[K]{
		Key:              key,
		Name:             name,
		previousStateMap: make(map[K]*State[K]),
		nextStateMap:     make(map[K]*State[K]),
	}
}
func (this *State[K]) HasPrevious() bool {
	return len(this.previousStateMap) > 0
}

func (this *State[K]) HasNext() bool {
	return len(this.nextStateMap) > 0
}

func (this *State[K]) addPreviousState(state *State[K]) {
	this.previousStateMap[state.Key] = state
}

func (this *State[K]) AddNextState(state *State[K]) {
	this.nextStateMap[state.Key] = state
	state.addPreviousState(this)
}

func (this *State[K]) Previous(key K) (*State[K], error) {
	data, ok := this.previousStateMap[key]
	if ok {
		return data, nil
	} else {
		return nil, errors.New(fmt.Sprintf("previous state[%v] not found", key))
	}
}

func (this *State[K]) Next(key K) (*State[K], error) {
	data, ok := this.nextStateMap[key]
	if ok {
		return data, nil
	} else {
		return nil, errors.New(fmt.Sprintf("next state[%v] not found", key))
	}
}

func (this *State[K]) CheckPrevious(key K) bool {
	if _, err := this.Previous(key); err != nil {
		return false
	}
	return true
}

func (this *State[K]) CheckNext(key K) bool {
	if _, err := this.Next(key); err != nil {
		return false
	}
	return true
}
