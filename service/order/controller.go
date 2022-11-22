package order

import (
	"errors"

	"mr.jackpot-backend/model"
)



type OrderControllerLayer interface {
	CreateOrder(userid int, order model.Order, delivery model.DeliveryInfo) error
	FinishOrderStep(id int) error
	CeaseOrder(id int) error
}



func (o *OrderManager) CreateOrder(userid int, order model.Order, delivery model.DeliveryInfo) error {

	orderid, err := o.db.CreateOrder(userid, order, delivery)
	if err != nil {
		return err
	}

	o.Orders[orderid] = NewOrder(orderid)
	o.Orders[orderid].CreaetOrder(order, delivery)
	return nil
}

func (o *OrderManager) CeaseOrder(id int) error {

	order, exist := o.Orders[id]
	if exist == false {
		return errors.New("")
	}

	err := order.CeaseOrder()
	if err != nil {
		return err
	}
	return nil
}

func (o *OrderManager) FinishOrderStep(id int) error {
	order, exist := o.Orders[id]
	if exist == false {
		return errors.New("")
	}

	order.SetState(*order.GetNextStep())
	order.ProcessStep()
	return nil
}

func (o *OrderManager) DeleteOrder(id int) error {
	_, exist := o.Orders[id]
	if exist == false {
		return errors.New("")
	}

	delete(o.Orders, id)
	return nil
}