package inventory

import (
	"mr.jackpot-backend/database/db"
	"mr.jackpot-backend/model"
)



type InventoryItem struct {
	model.InventoryItem
	db db.InventoryLayer
}

func (i *InventoryItem) UpdateInventoryItem(count int) error {
	if err := i.db.UpdateInventoryItem(i.ID, count); err != nil {
		return err
	}

	i.Count = count
	return nil
}

func (i *InventoryItem) GetInventoryItem() model.InventoryItem {
	return model.InventoryItem{
		ID: i.ID,
		Name: i.Name,
		Count: i.Count,
	}
}

func NewInventoryItem(item model.InventoryItem) *InventoryItem {
	Item := &InventoryItem{
		db: db.NewInventoryDB(),
	}

	Item.ID = item.ID
	Item.Name = item.Name
	Item.Count = item.Count

	return Item
}