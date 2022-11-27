package worker

import (
	"mr.jackpot-backend/database/db"
	"mr.jackpot-backend/service/order"
)




type StylerManager struct {
	o order.OrderLayer
	WorkerManager
}

func (w *StylerManager) FinishTaskProcess(orderid int) error {
	return w.o.FinishOrderStep(orderid);
}

var Styler = StylerManager{}


func init() {
	Styler.Workers = make(map[int]*Worker)
	Delivery.db = db.NewStaffDB()
}



