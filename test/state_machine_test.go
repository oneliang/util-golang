package test

import (
	"fmt"
	"github.com/oneliang/util-golang/state"
	"testing"
)

func TestStatusMachine(t *testing.T) {
	firstState := state.NewState[string]("one", "first")
	secondState := state.NewState[string]("two", "second")
	thirdState := state.NewState[string]("three", "third")
	firstState.AddNextState(secondState)
	secondState.AddNextState(thirdState)
	stateMachine := state.NewStateMachine(firstState)
	if err := stateMachine.NextState("tow"); err != nil {
		fmt.Println(fmt.Sprintf("%v", err))
	}
	if err := stateMachine.NextState("two"); err != nil {
		fmt.Println(fmt.Sprintf("%v", err))
	}
	fmt.Println(stateMachine.CurrentState.Key)
}
