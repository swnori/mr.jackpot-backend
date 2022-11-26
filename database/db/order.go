package db

import (
	"context"
	"time"

	"mr.jackpot-backend/database/orm"
	"mr.jackpot-backend/model"
)

type OrderLayer interface {
	CreateOrder(userid int, order model.Order, info model.AllOrderInfo) (id int, err error)
	UpdateOrderState(orderid int, orderstate string) error
}

type OrderDB struct {
	DBAccessor
}

func NewOrderDB() *OrderDB {
	db := &OrderDB{}
	db.q = NewAccessor()

	return db
}

func (db *OrderDB) CreateOrder(userid int, order model.Order, info model.AllOrderInfo) (id int, err error) {
	ctx := context.Background()

	result, err := db.q.CreateOrderInfo(ctx, orm.CreateOrderInfoParams{
		UserID:    int64(userid),
		Price:     int32(info.Price),
		Discount:  int32(info.CouponPrice),

		ReserveAt: info.ReserveAt,
		CreatedAt: time.Now(),

		Name:    info.Name,
		Address: info.Address,
		Phone:   info.Phone,
		Message: info.Message,
	})

	if err != nil {
		return
	}

	orderID, err := result.LastInsertId()
	id = int(orderID)
	if err != nil {
		return
	}

	if err := db.q.CreateOrderState(ctx, orderID); err != nil {
		return 0, err
	}

	for _, dinner := range order.DinnerList {
		result, err := db.q.CreateOrderedDinner(ctx, orm.CreateOrderedDinnerParams{
			OrderID: orderID,
			StyleID: int32(dinner.StyleId),
		})
		if err != nil {
			return id, err
		}

		dinnerID, err := result.LastInsertId()
		if err != nil {
			return id, err
		}

		for _, menu := range dinner.MenuList {

			menuStruct := orm.CreateOrderedMenuParams{
				OrderID:  orderID,
				DinnerID: dinnerID,
				MenuID:   int32(menu.MenuId),
				Count:    int32(menu.Count),
			}

			if len(menu.OptionId) >= 1 {
				menuStruct.Option1ID = int32(menu.OptionId[0])
			}
			if len(menu.OptionId) >= 2 {
				menuStruct.Option2ID = int32(menu.OptionId[1])
			}

			if err := db.q.CreateOrderedMenu(ctx, menuStruct); err != nil {
				return id, err
			}
		}
	}

	return id, nil
}

func (db *OrderDB) UpdateOrderState(orderid int, orderstate string) error {
	ctx := context.Background()

	return db.q.UpdateOrderState(ctx, orm.UpdateOrderStateParams{
		Name:    orderstate,
		OrderID: int64(orderid),
	})
}
