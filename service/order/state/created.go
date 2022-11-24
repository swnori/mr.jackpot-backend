package state

import "mr.jackpot-backend/model"


type CreatedState struct {
	State
	ID int
	Order OrderInterface
	NextStep *OrderState
	CeasedStep *OrderState
}

func (o *CreatedState) ProcessStep() error {
	return nil
}
func (o *CreatedState) CeaseStep() error {
	return nil
}

func (o *CreatedState) GetStateName() string {
	return model.StateCreated
}
