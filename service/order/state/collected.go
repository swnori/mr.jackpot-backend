package state

import (
	"errors"

	"mr.jackpot-backend/model"
)




type CollectedState struct {
	State
	ID int
	Order OrderInterface
	NextStep *OrderState
}

func (o *CollectedState) ProcessStep() error {
	return nil
}
func (o *CollectedState) CeaseStep() error {
	return errors.New("")
}
func (o *CollectedState) GetStateName() string {
	return model.StateCollected
}

