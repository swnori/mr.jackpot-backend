package order

import (
	"mr.jackpot-backend/database/db"
	"mr.jackpot-backend/model"
)


type OrderLayer interface {
	OrderProvider
	OrderController
}

type OrderManager struct {
	Orders map[int]*Order
	db db.OrderLayer

	Dinner []model.DinnerFormed
	Menu   []model.MenuFormed
}

var OrderManagers = &OrderManager{}

func NewOrderManager() *OrderManager {
	OrderManagers.Orders = make(map[int]*Order)
	OrderManagers.db = db.NewOrderDB()
	return OrderManagers
}
