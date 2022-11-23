package state

import (
	"errors"

	"mr.jackpot-backend/model"
)



type PreparedState struct {
	State
	ID int
	NextStep *OrderState
}

func (o *PreparedState) ProcessStep() error {
	return nil
}
func (o *PreparedState) CeaseStep() error {
	return errors.New("")
}

func (o *PreparedState) GetStateName() string {
	return model.StatePrepared
}