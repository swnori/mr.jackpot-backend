package state

import "mr.jackpot-backend/model"

type CanceledState struct {
	State
	ID int
	Order OrderInterface
	NextStep *OrderState
}

func (o *CanceledState) ProcessStep() error {
	return nil
}
func (o *CanceledState) CeaseStep() error {
	return nil
}

func (o *CanceledState) GetStateName() string {
	return model.StateCreated
}
