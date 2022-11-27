package order

import (
	"errors"

	"mr.jackpot-backend/model"
)



type OrderProvider interface {
	GetOrderInfo(orderid int) (model.OrderResponse, error)
	GetOrderState(int) (string, error)
	GetAllOrderInfo() []model.OrderResponse
	CheckOrderOwner(orderid, userid int) error
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
