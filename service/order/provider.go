package order

import (
	"errors"

	"mr.jackpot-backend/model"
)



type OrderProvider interface {
	GetOrderInfo(orderid int) (model.OrderResponse, error)
	GetOrderState(int) (int, error)
	GetAllOrderInfo() []model.OrderResponse
	GetOrderIdByUserId(userid int) (int, error)
	GetAllOrderSummary() ([]model.OrderSummary, error)
	GetMenuInfo(menuid int) (model.MenuFormed, error)
	GetDinnerInfo(dinnerid int) (model.DinnerFormed, error)
	GetAllMenuInfo() ([]model.MenuFormed, error)
	GetAllDinnerInfo() ([]model.DinnerFormed, error)
}

func (o *OrderManager) GetAllMenuInfo() ([]model.MenuFormed, error) {
	menulist := make([]model.MenuFormed, 0)

	for _, order := range o.Orders {
		for _, dinner := range order.Order.DinnerList {
			for _, menu := range dinner.MenuList {
				menulist = append(menulist, model.MenuFormed{
					MenuID: menu.MenuId,
					ID: menu.OrderedMenuId,
					OptionList: menu.OptionId,
				})
			}
		}
	}

	return menulist, nil
}

func (o *OrderManager) GetMenuInfo(menuid int) (model.MenuFormed, error) {

	for _, order := range o.Orders {
		for _, dinner := range order.Order.DinnerList {
			for _, menu := range dinner.MenuList {
				if menu.OrderedMenuId == menuid {
					return model.MenuFormed{
						MenuID: menu.MenuId,
						ID: menuid,
						OptionList: menu.OptionId,
					}, nil
				}
			}
		}
	}

	return model.MenuFormed{}, errors.New("no menu matched")
}

func (o *OrderManager) GetAllDinnerInfo() ([]model.DinnerFormed, error) {
	dinnerlist := make([]model.DinnerFormed, 0)

	for _, order := range o.Orders {
		for _, dinner := range order.Order.DinnerList {
			menulist := make([]int, 0)
			for _, menu := range dinner.MenuList {
				menulist = append(menulist, menu.MenuId)
			}

			dinnerlist = append(dinnerlist, model.DinnerFormed{
				ID: dinner.OrderedDinnerId,
				DinnerID: dinner.DinnerId,
				MenuList: menulist,
			})
		}
	}

	return dinnerlist, nil
}

func (o *OrderManager) GetDinnerInfo(dinnerid int) (model.DinnerFormed, error) {
	for _, order := range o.Orders {
		for _, dinner := range order.Order.DinnerList {
			if dinner.OrderedDinnerId == dinnerid {
				menulist := make([]int, 0)
				for _, menu := range dinner.MenuList {
					menulist = append(menulist, menu.MenuId)
				}

				return model.DinnerFormed{
					ID: dinnerid,
					DinnerID: dinner.DinnerId,
					MenuList: menulist,
				}, nil
			}
		}
	}

	return model.DinnerFormed{}, errors.New("no menu matched")
}


func (o *OrderManager) GetAllOrderSummary() ([]model.OrderSummary, error) {
	orderlist := make([]model.OrderSummary, 0)
	
	for id, order := range o.Orders {
		orderinfo := order.GetOrderInfo()

		dinnerlist := make([]int, 0)
		for _, dinner := range order.Order.DinnerList {
			dinnerlist = append(dinnerlist, dinner.DinnerId)
		}

		orderlist = append(orderlist, model.OrderSummary{
			OrderID: id,
			StateID: order.GetOrderState(),
			ReserveAt: orderinfo.CreatedAt.Format(model.TimeSecondFormat),
			Price: orderinfo.Price,
			DinnerList: dinnerlist,
		})
	}

	return orderlist, nil
}



func (o *OrderManager) GetOrderIdByUserId(userid int) (int, error) {
	for orderid, order := range o.Orders {
		if order.GetOrderInfo().OwnerID == userid {
			return orderid, nil
		}
	}

	return 0, errors.New("no order match")
}


func (o *OrderManager) GetOrderInfo(orderid int) (model.OrderResponse, error) {

	order, exist := o.Orders[orderid]
	if !exist {
		return model.OrderResponse{}, errors.New("orderid not match with order")
	}
	info := order.GetOrderInfo()

	return model.OrderResponse{
		Order: order.GetOrder(),
		AllOrderInfoResponse: model.AllOrderInfoResponse{
			ID: orderid,
			StateID: order.GetOrderState(),
			Name: info.Name,
			Address: info.Address,
			Phone: info.Phone,
			Message: info.Message,
			Price: info.Price,
			CouponPrice: info.CouponPrice,
			CouponName: info.CouponName,

			ReserveAt: info.ReserveAt.Format(model.TimeSecondFormat),
			CreatedAt: info.CreatedAt.Format(model.TimeSecondFormat),
		},
	}, nil
}

func (o *OrderManager) GetAllOrderInfo() []model.OrderResponse {
	orders := make([]model.OrderResponse, 0)
	for _, order := range o.Orders {
		orders = append(orders, model.OrderResponse{
			Order: order.GetOrder(),
			//AllOrderInfoResponse: order.GetOrderInfo(),
		})
	}
	return orders
}

func (o *OrderManager) GetOrderState(orderid int) (int, error) {
	order, exist := o.Orders[orderid]
	if !exist {
		return 0, errors.New("")
	}

	return order.GetOrderState(), nil
}

func (o *OrderManager) CheckOrderOwner(orderid int, userid int) error {
	return nil
}
