package state

import (
	"errors"

)

type RejectedState struct {
	State
	ID       int
	StateID  int
	Order    OrderInterface
	NextStep *OrderState
}

func (o *RejectedState) ProcessStep() error {
	return nil
}
func (o *RejectedState) CeaseStep() error {
	return errors.New("")
}

func (o *RejectedState) GetStateId() int {
	return o.StateID
}
