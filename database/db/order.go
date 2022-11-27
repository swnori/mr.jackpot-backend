package db

import (
	"context"
	"database/sql"
	"time"

	"mr.jackpot-backend/database/orm"
	"mr.jackpot-backend/model"
)

type OrderLayer interface {
	CreateOrder(userid int, order model.Order, info model.AllOrderInfo) (id int, err error)
	UpdateOrderState(orderid int, orderstate string) error
	GetOrderHisory(userid int) ([]model.OrderFormed, error)
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
			DinnerID: int32(dinner.DinnerId),
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
				menuStruct.Option1ID = sql.NullInt32{
					Int32: int32(menu.OptionId[0]),
					Valid: true,
				}
			}
			if len(menu.OptionId) >= 2 {
				menuStruct.Option2ID = sql.NullInt32{
					Int32: int32(menu.OptionId[1]),
					Valid: true,
				}
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


func (db *OrderDB) GetOrderHisory(userid int) ([]model.OrderFormed, error) {
	ctx := context.Background()

	orderlist, err := db.q.GetOrderHistory(ctx, int64(userid))

	orderList := make([]model.OrderFormed, 0)
	for _, order := range orderlist {
		orderitem := model.OrderFormed{
			OrderID: int(order.OrderID),
			Price: int(order.Price) - int(order.Discount),
			CreatedAt: order.CreatedAt.Format(model.TimeSecondFormat),
			ReserveAt: order.ReserveAt.Format(model.TimeSecondFormat),
		}
		dinnerlist, err := db.q.GetDinnerListHistory(ctx, order.OrderID)
		orderitem.DinnerList = dinnerlist

		if err != nil {
			return nil, err
		}
		orderList = append(orderList, orderitem)
	}

	return orderList, err
}