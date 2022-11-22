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
	db: db.NewStockDB(),
}

func NewStock() *StockEntity {
	return Stock
}

func Initialize() {
	itemList, err := Stock.db.GetAllStockList()
	if err != nil {
		panic(err)
	}

	for _, item := range itemList {
		Stock.ItemList[item.ID] = NewStockItem(item)
	}
}
