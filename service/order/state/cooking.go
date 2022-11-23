package state

import (
	"errors"

	"mr.jackpot-backend/model"
)



type CookingState struct {
	State
}

func (o *CookingState) ProcessStep() error {
	return nil
}
func (o *CookingState) CeaseStep() error {
	return errors.New("")
}

func (o *CookingState) GetStateName() string {
	return model.StateCooking
}