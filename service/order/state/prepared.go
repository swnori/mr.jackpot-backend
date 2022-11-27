package state

import (
	"errors"
)

type PreparedState struct {
	State
	ID       int
	StateID  int
	Order    OrderInterface
	NextStep *OrderState
}

func (o *PreparedState) ProcessStep() error {
	return nil
}
func (o *PreparedState) CeaseStep() error {
	return errors.New("")
}

func (o *PreparedState) GetStateId() int {
	return o.StateID
}
