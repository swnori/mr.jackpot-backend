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
	/*
	taskList := o.Order.GetAllTaskList()

	orderid := o.ID
	styleSubTask := make([]int, 0)

	for dinnerid, menulist := range taskList {
		styleSubTask = append(styleSubTask, dinnerid)

		worker.Cook.StartTaskProcess(orderid, dinnerid, menulist)
	}
	worker.Styler.StartTaskProcess(orderid, orderid, styleSubTask)
*/
	/*
	이걸 받은 입장에서 초기화를 따로 하는걸로

	제약 조건 :
	모든 order은 디너를 반드시 하나 포함한다
	모든 dinner은 order 을 반드시 하나 포함한다
	*/

	return nil
}

func (o *StartedState) CeaseStep() error {
	return errors.New("")
}
func (o *StartedState) GetStateName() string {
	return model.StateStarted
}