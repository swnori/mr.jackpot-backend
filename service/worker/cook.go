package worker

import (
	"mr.jackpot-backend/database/db"
	"mr.jackpot-backend/service/order"
)





type CookManager struct {
	o order.OrderLayer
	WorkerManager
}

func (w *CookManager) FinishTaskProcess(orderid int) error {
	return nil
	// 로직 보충 필요함
}

var Cook *CookManager


func init() {
	Cook.Workers = make(map[int]*Worker)
	Delivery.db = db.NewStaffDB()
}



