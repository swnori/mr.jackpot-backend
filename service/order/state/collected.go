package state

import (
	"errors"

)

type CollectedState struct {
	State
	ID       int
	StateID  int
	Order    OrderInterface
	NextStep *OrderState
}

func (o *CollectedState) ProcessStep() error {
	return nil
}
func (o *CollectedState) CeaseStep() error {
	return errors.New("")
}
func (o *CollectedState) GetStateId() int {
	return o.StateID
}
