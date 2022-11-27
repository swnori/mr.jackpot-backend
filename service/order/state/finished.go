package state

import (
	"errors"

)

type FinishedState struct {
	State
	ID       int
	StateID  int
	Order    OrderInterface
	NextStep *OrderState
}

func (o *FinishedState) ProcessStep() error {
	return nil
}
func (o *FinishedState) CeaseStep() error {
	return errors.New("")
}

func (o *FinishedState) GetStateId() int {
	return o.StateID
}
