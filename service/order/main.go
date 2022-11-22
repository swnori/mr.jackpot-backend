package order

import "mr.jackpot-backend/database/db"


type OrderLayer interface {
	OrderProviderLayer
	OrderControllerLayer
}

type OrderManager struct {
	Orders map[int]*Order
	db db.OrderLayer
}

var OrderManagers = &OrderManager{
	Orders: make(map[int]*Order),
	db: db.NewOrderDB(),
}
