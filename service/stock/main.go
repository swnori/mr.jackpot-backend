package stock

import (
	"mr.jackpot-backend/database/db"
	"mr.jackpot-backend/model"
)

type StockController interface {
	AddStockItem(name string) error
	UpdateStockItem(id, count int) error
	DeleteStockItem(id int) error
	GetAllStockList() []model.StockItem
}

type StockEntity struct {
	db       db.StockLayer
	ItemList map[int]*StockItem
}

var H int

var Stock = &StockEntity{
	ItemList: make(map[int]*StockItem),
}

func NewStock() *StockEntity {
	return Stock
}

func Initialize() error {
	Stock.db = db.NewStockDB()

	itemList, err := Stock.db.GetAllStockList()

	for _, item := range itemList {
		Stock.ItemList[item.ID] = NewStockItem(item)
	}
	return err
}
