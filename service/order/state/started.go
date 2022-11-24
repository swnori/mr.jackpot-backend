package state

import (
	"errors"

	"mr.jackpot-backend/model"
)



type StartedState struct {
	State
	ID int
	Order OrderInterface
	NextStep *OrderState
}

func (o *StartedState) ProcessStep() error {
	return nil
}
func (o *StartedState) CeaseStep() error {
	return errors.New("")
}
func (o *StartedState) GetStateName() string {
	return model.StateStarted
}