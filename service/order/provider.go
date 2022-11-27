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
			ID: info.ID,
			StateID: info.StateID,
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
