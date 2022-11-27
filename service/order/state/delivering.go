package state

import (
	"errors"

)

type DeliveringState struct {
	State
	ID       int
	StateID  int
	Order    OrderInterface
	NextStep *OrderState
}

func (o *DeliveringState) ProcessStep() error {
	return nil
}
func (o *DeliveringState) CeaseStep() error {
	return errors.New("")
}

func (o *DeliveringState) GetStateId() int {
	return o.StateID
}
