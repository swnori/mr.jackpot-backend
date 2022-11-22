package order

import "mr.jackpot-backend/service/order"

type OrderService interface {
	CustomerOrderService
	StaffOrderService
}

type OrderHandler struct {
	order order.OrderLayer
}

func NewOrderHandler() *OrderHandler {
	return &OrderHandler{
		order: order.OrderManagers,
	}
}