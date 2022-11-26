package db

import (
	"context"

	"mr.jackpot-backend/database/orm"
	"mr.jackpot-backend/model"
)

type StockLayer interface {
	GetAllStockList() ([]model.StockItem, error)
	UpdateStockItem(id int, count int) error
	AddStockItem(name, unit string) (model.StockItem, error)
	DeleteStockItem(id int) error
}

type StockDB struct {
	DBAccessor
}

func NewStockDB() *StockDB {
	db := &StockDB{}
	db.q = NewAccessor()

	return db
}

func (db *StockDB) GetAllStockList() (stockList []model.StockItem, err error) {
	ctx := context.Background()

	itemList, err := db.q.GetAllStockList(ctx)

	stockList = make([]model.StockItem, len(itemList))
	for _, item := range itemList {
		stockList = append(stockList, model.StockItem{
			ID:    int(item.StockID),
			Name:  item.Name,
			Count: int(item.Count),
			Unit: item.Unit,
		})
	}

	return
}

func (db *StockDB) UpdateStockItem(id int, count int) error {
	ctx := context.Background()

	return db.q.UpdateStockItem(ctx, orm.UpdateStockItemParams{
		Count:   int32(count),
		StockID: int64(id),
	})
}

func (db *StockDB) AddStockItem(name, unit string) (item model.StockItem, err error) {
	ctx := context.Background()

	result, err := db.q.AddStockItem(ctx, orm.AddStockItemParams{
		Unit: unit,
		Name: name,
	})
	if err != nil {
		return
	}

	ID, err := result.LastInsertId()
	if err != nil {
		return
	}

	return model.StockItem{
		ID:    int(ID),
		Name:  name,
		Count: 0,
		Unit: item.Unit,
	}, nil
}

func (db *StockDB) DeleteStockItem(id int) error {
	ctx := context.Background()

	return db.q.DeleteStockItem(ctx, int64(id))
}
