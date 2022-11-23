package worker

import (
	"mr.jackpot-backend/database/db"
	"mr.jackpot-backend/service/order"
)




type DeliveryManager struct {
	o order.OrderLayer
	WorkerManager
}

func (w *DeliveryManager) FinishTaskProcess(orderid int) error {
	return w.o.FinishOrderStep(orderid);
}

var Delivery = &DeliveryManager{}


func init() {
	Delivery.Workers = make(map[int]*Worker)
	Delivery.db = db.NewStaffDB()
}