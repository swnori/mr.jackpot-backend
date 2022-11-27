package order

import (
	"errors"
	"time"

	"mr.jackpot-backend/model"
)

type OrderController interface {
	CreateOrder(userid int, info model.OrderRequestInfo, order model.Order) error
	FinishOrderStep(id int) error
	CeaseOrder(id int) error
}

func (o *OrderManager) CreateOrder(userid int, info model.OrderRequestInfo, order model.Order) error {
	
	reserveTime, err := time.Parse(model.TimeSecondFormat, info.ReserveAt)
	if err != nil {
		return err
	}

	orderinfo := model.AllOrderInfo{
		OwnerID: info.OwnerID,
		Name: info.Name,
		Address: info.Address,
		Phone: info.Phone,
		Message: info.Message,
		CreatedAt: time.Now(),
		ReserveAt: reserveTime,
		Price: info.Price,
	}

	orderid, err := o.db.CreateOrder(userid, order, orderinfo)
	if err != nil {
		return err
	}
	
	for i := range order.DinnerList {
		dinner := order.DinnerList[i]
		dinner.DinnerId = i*2;
		dinner.OrderedDinnerId = 1
	}

	o.Orders[orderid] = NewOrder(orderid)
	o.Orders[orderid].CreateOrder(order, orderinfo)

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
		return errors.New("no id")
	}
	if order.GetOrderState() == 4 {
		order.SetState(order.started)
	}
	//fmt.Println(*order.GetNextStep())
	//order.SetState()
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
