package state

import (
	"errors"

	"mr.jackpot-backend/model"
)



type RejectedState struct {
	State
	ID int
	Order OrderInterface
	NextStep *OrderState
}

func (o *RejectedState) ProcessStep() error {
	return nil
}
func (o *RejectedState) CeaseStep() error {
	return errors.New("")
}

func (o *RejectedState) GetStateName() string {
	return model.StatePrepared
}