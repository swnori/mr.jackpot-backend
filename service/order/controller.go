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
	SetMenuNextStep(id int) error
	SetDinnerNextStep(id int) error
}

func (o *OrderManager) SetMenuNextStep(id int) error {
	for i := range o.Menu {
		if o.Menu[i].ID == id {
			if o.Menu[i].StateID != 3 {
				o.Menu[i].StateID += 1
			}
		}
	}
	return nil
}

func (o *OrderManager) SetDinnerNextStep(id int) error {
	for i := range o.Dinner {
		if o.Dinner[i].ID == id {
			if o.Dinner[i].ID != 3 {
				o.Dinner[i].StateID += 1
			}
		}
	}
	//for _, order := range o.Orders {
	//	for did, dinner := range order.Order.DinnerList {
	//		if dinner.DinnerId != 3 {
	//			
	//		}
	//	}
	//}
	return nil
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


	menulist := make([]model.MenuFormed, 0)
	for _, dinner := range neworder.DinnerList {
		for _, menu := range dinner.MenuList {

			for i := 0; i < menu.Count; i++ {
				menulist = append(menulist, model.MenuFormed{
					StateID: 1,
					OrderedID: orderid,
					DinnerID: dinner.DinnerId,
					MenuID: menu.MenuId,
					ID: menu.OrderedMenuId,
					OptionList: menu.OptionId,
				})
			}
		}
	}
	o.Menu = menulist

	dinnerlist := make([]model.DinnerFormed, 0)
	for _, dinner := range neworder.DinnerList {
		menulist := make([]int, 0)
		for _, menu := range dinner.MenuList {
			menulist = append(menulist, menu.MenuId)
		}

		dinnerlist = append(dinnerlist, model.DinnerFormed{
			StateID: 1,
			OrderedID: orderid,
			ID: dinner.OrderedDinnerId,
			DinnerID: dinner.DinnerId,
			StyleID: dinner.StyleId,
			MenuList: menulist,
		})
	}
	o.Dinner = dinnerlist

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
	case 7:
		order.SetState(order.delivering)
		break
	case 8:
		order.SetState(order.delivered)
		break
	case 9:
		order.SetState(order.requested)
		break
	case 10:
		order.SetState(order.collected)
		break
	case 11:
		order.SetState(order.finished)
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
