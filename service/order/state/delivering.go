package state

import (
	"errors"

	"mr.jackpot-backend/model"
)



type DeliveringState struct {
	State
	ID int
	Order OrderInterface
	NextStep *OrderState
}

func (o *DeliveringState) ProcessStep() error {
	return nil
}
func (o *DeliveringState) CeaseStep() error {
	return errors.New("")
}

func (o *DeliveringState) GetStateName() string {
	return model.StateDelivering
}