package state

import (
	"mr.jackpot-backend/model"
)




type AcceptedState struct {
	State
	ID int
	Order OrderInterface
	NextStep *OrderState
	CeasedStep *OrderState
}

func (o *AcceptedState) ProcessStep() error {

	return nil
}
func (o *AcceptedState) CeaseStep() error {
	return nil
}
func (o *AcceptedState) GetStateName() string {
	return model.StateAccepted
}