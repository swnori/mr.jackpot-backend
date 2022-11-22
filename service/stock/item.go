package stock

import (
	"mr.jackpot-backend/database/db"
	"mr.jackpot-backend/model"
)

type StockItem struct {
	model.StockItem
	db db.StockLayer
}

func (i *StockItem) UpdateStockItem(count int) error {
	if err := i.db.UpdateStockItem(i.ID, count); err != nil {
		return err
	}

	i.Count = count
	return nil
}

func (i *StockItem) GetStockItem() model.StockItem {
	return model.StockItem{
		ID:    i.ID,
		Name:  i.Name,
		Count: i.Count,
	}
}

func NewStockItem(item model.StockItem) *StockItem {
	Item := &StockItem{
		db: db.NewStockDB(),
	}

	Item.ID = item.ID
	Item.Name = item.Name
	Item.Count = item.Count

	return Item
}
