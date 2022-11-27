package state

import (
	"errors"

)

type DeliveredState struct {
	State
	ID       int
	StateID  int
	Order    OrderInterface
	NextStep *OrderState
}

func (o *DeliveredState) ProcessStep() error {
	return nil
}
func (o *DeliveredState) CeaseStep() error {
	return errors.New("")
}

func (o *DeliveredState) GetStateId() int {
	return o.StateID
}
