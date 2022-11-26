package stock

import (
	"errors"

	"mr.jackpot-backend/model"
)

func (i *StockEntity) GetAllStockList() []model.StockItem {
	list := make([]model.StockItem, 0)

	for _, item := range i.ItemList {
		list = append(list, item.GetStockItem())
	}

	return list
}

func (i *StockEntity) AddStockItem(name, unit string) error {
	item, err := i.db.AddStockItem(name, unit)
	if err != nil {
		return err
	}
	id := item.ID

	_, exist := i.ItemList[id]
	if exist {
		return errors.New("")
	}

	i.ItemList[id] = NewStockItem(item)
	return nil
}

func (i *StockEntity) UpdateStockItem(id, count int) error {
	_, exist := i.ItemList[id]
	if !exist {
		return errors.New("")
	}

	err := i.ItemList[id].UpdateStockItem(count)
	return err
}

func (i *StockEntity) DeleteStockItem(id int) error {
	_, exist := i.ItemList[id]
	if !exist {
		return errors.New("")
	}

	delete(i.ItemList, id)
	
	return nil
}
