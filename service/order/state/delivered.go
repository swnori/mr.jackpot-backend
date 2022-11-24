package state

import (
	"errors"

	"mr.jackpot-backend/model"
)



type DeliveredState struct {
	State
	ID int
	Order OrderInterface
	NextStep *OrderState
}

func (o *DeliveredState) ProcessStep() error {
	return nil
}
func (o *DeliveredState) CeaseStep() error {
	return errors.New("")
}

func (o *DeliveredState) GetStateName() string {
	return model.StateDelivered
}
