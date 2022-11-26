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
	
	reserveTime, err := time.Parse(time.RFC1123, info.ReserveAt)
	if err != nil {
		return err
	}

	orderinfo := model.AllOrderInfo{
		OwnerID: info.OwnerID,
		Name: info.Name,
		Address: info.Address,
		Phone: info.Phone,
		Message: info.Message,
		ReserveAt: reserveTime,
		CreatedAt: time.Now(),
		Price: info.Price,
	}

	//var couponid = order.CouponID
	//이후로 요청하는 로직이 있을거야

	orderid, err := o.db.CreateOrder(userid, order, orderinfo)
	if err != nil {
		return err
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
