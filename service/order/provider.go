package order

import (
	"errors"

	"mr.jackpot-backend/model"
)



type OrderProviderLayer interface {
	GetOrderInfo(int) model.Order
	GetDeliveryInfo(int) model.DeliveryInfo
	GetAllOrderInfo() []model.Order
	GetAllDeliveryInfo() []model.DeliveryInfo
	GetOrderState(int) (string, error)
	CheckOrderOwner(orderid, userid int) error
}



func (o *OrderManager) GetOrderInfo(orderid int) model.Order {
	return o.Orders[orderid].GetOrderInfo()
}

func (o *OrderManager) GetAllOrderInfo() []model.Order {
	orders := make([]model.Order, 0)
	for _, order := range o.Orders {
		orders = append(orders, order.GetOrderInfo())
	}
	return orders
}

func (o *OrderManager) GetDeliveryInfo(orderid int) model.DeliveryInfo {
	return o.Orders[orderid].GetDeliveryInfo()
}

func (o *OrderManager) GetAllDeliveryInfo() []model.DeliveryInfo {
	deliveries := make([]model.DeliveryInfo, 0)
	for _, order := range o.Orders {
		deliveries = append(deliveries, order.GetDeliveryInfo())
	}
	return deliveries
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
