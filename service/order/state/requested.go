package state

import (
	"errors"

	"mr.jackpot-backend/model"
)




type RequestedState struct {
	State
	ID int
	NextStep *OrderState
}

func (o *RequestedState) ProcessStep() error {
	return nil
}
func (o *RequestedState) CeaseStep() error {
	return errors.New("")
}

func (o *RequestedState) GetStateName() string {
	return model.StateRequested
}
