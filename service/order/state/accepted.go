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
	o.Order.GetAllTaskList()	
	// o.Order.GetAllTaskList() // 어떤 방식으로 요청할지는 잘 모르겠다
	//for : worker.Cook.StartTaskProcess()
	return nil
}
func (o *AcceptedState) CeaseStep() error {
	return nil
}
func (o *AcceptedState) GetStateName() string {
	return model.StateAccepted
}