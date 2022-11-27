package state

import (
	"errors"

)

type RequestedState struct {
	State
	ID       int
	StateID  int
	Order    OrderInterface
	NextStep *OrderState
}

func (o *RequestedState) ProcessStep() error {
	return nil
}
func (o *RequestedState) CeaseStep() error {
	return errors.New("")
}

func (o *RequestedState) GetStateId() int {
	return o.StateID
}
