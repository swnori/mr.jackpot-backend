package inventory

import (
	"mr.jackpot-backend/database/db"
	"mr.jackpot-backend/model"
)



type InventoryController interface {
	AddInventoryItem(name string) error
	UpdateInventoryItem(id, count int) error
	DeleteInventoryItem(id int) error
	GetAllInventoryList() []model.InventoryItem
}

type InventoryEntity struct {
	db db.InventoryLayer
	ItemList map[int]*InventoryItem
}

var H int

var Inventory = &InventoryEntity{
	db: db.NewInventoryDB(),
}

func NewInventory() *InventoryEntity{
	return Inventory
}

func Initialize() {
	itemList, err := Inventory.db.GetAllInventoryList()
	if err != nil {
		panic(err)
	}

	for _, item := range itemList {
		Inventory.ItemList[item.ID] = NewInventoryItem(item)
	}
}