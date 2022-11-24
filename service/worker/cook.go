package worker

import (
	"errors"

	"mr.jackpot-backend/database/db"
	"mr.jackpot-backend/service/order"
)





type CookManager struct {
	o order.OrderLayer
	WorkerManager
}

func (w *CookManager) FinishTaskProcess(orderid int) error {
	return nil
}

func (w *CookManager) StartSubTask(workerid int, taskid int) error {
	// 부모 메서드를 호출한 이후에 동작하는 메서드
	// 해당 orderid 를 찾음. 만약 state를 order.SetNextState() 메서드 호출

	return nil
}


var Cook = &CookManager{}


func init() {
	Cook.Workers = make(map[int]*Worker)
	Delivery.db = db.NewStaffDB()
}



