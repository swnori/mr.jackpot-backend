package state

import (
	"errors"

	"mr.jackpot-backend/model"
)



type FinishedState struct {
	State
	ID int
	Order OrderInterface
	NextStep *OrderState
}

func (o *FinishedState) ProcessStep() error {
	return nil
}
func (o *FinishedState) CeaseStep() error {
	return errors.New("")
}

func (o *FinishedState) GetStateName() string {
	return model.StateFinished
}