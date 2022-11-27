package order

import (
	"errors"
	"sort"
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

	orderid, neworder, err := o.db.CreateOrder(userid, order, orderinfo)
	if err != nil {
		return err
	}


	sort.Slice(neworder.DinnerList, func(i, j int) bool {
        return neworder.DinnerList[i].OrderedDinnerId < neworder.DinnerList[j].OrderedDinnerId
    })

	for _, dinner := range neworder.DinnerList {
		sort.Slice(dinner.MenuList, func(i, j int) bool {
			return dinner.MenuList[i].OrderedMenuId < dinner.MenuList[j].OrderedMenuId
		})	
	}

	o.Orders[orderid] = NewOrder(orderid)
	o.Orders[orderid].CreateOrder(neworder, orderinfo)

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
	state := order.GetOrderState()
	switch (state) {
	case 4:
		order.SetState(order.accepted)
		break
	case 5:
		order.SetState(order.started)
		break
	case 6:
		order.SetState(order.prepared)
		break
	default:
		return errors.New("unexpected order state")
	}

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
