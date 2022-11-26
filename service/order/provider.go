package order

import (
	"errors"
	"fmt"

	"mr.jackpot-backend/model"
)



type OrderProvider interface {
	GetOrderInfo(int) (model.OrderResponse, error)
	GetOrderState(int) (string, error)
	GetAllOrderInfo() []model.OrderResponse
	CheckOrderOwner(orderid, userid int) error
}



func (o *OrderManager) GetOrderInfo(userid int) (model.OrderResponse, error) {
	var (
		orderid int
		exist bool
	)

	for id, order := range o.Orders {
		fmt.Println(order.GetOrderInfo().OwnerID)
		if order.GetOrderInfo().OwnerID == userid {
			orderid = id
			exist = true
			break			
		}
	}

	if !exist {
		return model.OrderResponse{}, errors.New("NNO ORDER")
	}

	order := o.Orders[orderid]

	return model.OrderResponse{
		Order: order.GetOrder(),
		AllOrderInfo: order.GetOrderInfo(),
	}, nil
}

func (o *OrderManager) GetAllOrderInfo() []model.OrderResponse {
	orders := make([]model.OrderResponse, 0)
	for _, order := range o.Orders {
		orders = append(orders, model.OrderResponse{
			Order: order.GetOrder(),
			AllOrderInfo: order.GetOrderInfo(),
		})
	}
	return orders
}

func (o *OrderManager) GetOrderState(orderid int) (string, error) {
	order, exist := o.Orders[orderid]
	if !exist {
		return "", errors.New("")
	}

	return order.GetOrderState(), nil
}

func (o *OrderManager) CheckOrderOwner(orderid int, userid int) error {
	return nil
}
