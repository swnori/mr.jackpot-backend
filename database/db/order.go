package db

import (
	"context"
	"database/sql"
	"time"

	"mr.jackpot-backend/database/orm"
	"mr.jackpot-backend/model"
)

type OrderLayer interface {
	CreateOrder(userid int, order model.Order, info model.AllOrderInfo) (id int, rorder model.Order, err error)
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

func (db *OrderDB) CreateOrder(userid int, order model.Order, info model.AllOrderInfo) (id int, null model.Order, err error) {
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
		return 0, null, err
	}


	for didx, dinner := range order.DinnerList {
		result, err := db.q.CreateOrderedDinner(ctx, orm.CreateOrderedDinnerParams{
			OrderID: orderID,
			StyleID: int32(dinner.StyleId),
			DinnerID: int32(dinner.DinnerId),
		})

		if err != nil {
			return id, null, err
		}

		dinnerID, err := result.LastInsertId()
		if err != nil {
			return id, null, err
		}

		order.DinnerList[didx].OrderedDinnerId = int(dinnerID)

		newMenuList := make([]model.MenuOrder, 0)

		for _, menu := range dinner.MenuList {

			menuStruct := orm.CreateOrderedMenuParams{
				OrderID:  orderID,
				DinnerID: dinnerID,
				MenuID:   int32(menu.MenuId),
				Count:    1,
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

			for i := 0; i < menu.Count; i++ {

				result, err := db.q.CreateOrderedMenu(ctx, menuStruct)
				if err != nil {
					return 0, order, err
				}
				menuid, err := result.LastInsertId()
				if err != nil {
					return 0, order, err
				}

				newMenuList = append(newMenuList, model.MenuOrder{
					OrderedMenuId: int(menuid),
					MenuId: menu.MenuId,
					Count: 1,
					OptionId: menu.OptionId,
					StateId: 0,
				})
			}
		}

		order.DinnerList[didx].MenuList = make([]model.MenuOrder, 0)
		order.DinnerList[didx].MenuList = append(order.DinnerList[didx].MenuList, newMenuList...)
	}

	return id, order, nil
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