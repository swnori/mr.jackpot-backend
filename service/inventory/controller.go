package inventory

import (
	"errors"

	"mr.jackpot-backend/model"
)


func (i *InventoryEntity) GetAllInventoryList() []model.InventoryItem {
	list := make([]model.InventoryItem, 0)

	for _, item := range i.ItemList {
		list = append(list, item.GetInventoryItem())
	}

	return list
}

func (i *InventoryEntity) AddInventoryItem(name string) error {
	item, err := i.db.AddInventoryItem(name)
	if err != nil {
		return err
	}
	id := item.ID

	_, exist := i.ItemList[id]
	if exist {
		return  errors.New("")
	}

	i.ItemList[id] = NewInventoryItem(item)
	return nil
}

func (i *InventoryEntity) UpdateInventoryItem(id, count int) error {
	_, exist := i.ItemList[id]
	if !exist {
		return  errors.New("")
	}

	err := i.ItemList[id].UpdateInventoryItem(count);
	return err
}

func (i *InventoryEntity) DeleteInventoryItem(id int) error {
	_, exist := i.ItemList[id]
	if !exist {
		return  errors.New("")
	}

	delete(i.ItemList, id)
	return nil
}