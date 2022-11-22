package db

import (
	"context"

	"mr.jackpot-backend/database/orm"
	"mr.jackpot-backend/model"
)



type InventoryLayer interface {
	GetAllInventoryList() ([]model.InventoryItem, error)
	UpdateInventoryItem(id int, count int) error
	AddInventoryItem(name string) (model.InventoryItem, error)
	DeleteInventoryItem(id int) error
}

type InventoryDB struct {
	DBAccessor
}

func NewInventoryDB() *InventoryDB {
	db := &InventoryDB{}
	db.q = NewAccessor()

	return db
}



func (db *InventoryDB) GetAllInventoryList() (inventoryList []model.InventoryItem, err error) {
	ctx := context.Background()

	itemList, err := db.q.GetAllInventoryList(ctx)

	inventoryList = make([]model.InventoryItem, len(itemList))
	for _, item := range itemList {
		inventoryList = append(inventoryList, model.InventoryItem{
			ID: int(item.StockID),
			Name: item.Name,
			Count: int(item.Count),
		})
	}

	return
}

func (db *InventoryDB) UpdateInventoryItem(id int, count int) error {
	ctx := context.Background()

	return db.q.UpdateInventoryItem(ctx, orm.UpdateInventoryItemParams{
		Count: int32(count),
		StockID: int64(id),
	})
}

func (db *InventoryDB) AddInventoryItem(name string) (item model.InventoryItem, err error) {
	ctx := context.Background()

	result, err := db.q.AddInventoryItem(ctx, name);
	if err != nil {
		return
	}

	ID, err := result.LastInsertId()
	if err != nil {
		return 
	}
	
	return model.InventoryItem{
		ID: int(ID),
		Name: name,
		Count: 0,
	}, nil
}

func (db *InventoryDB) DeleteInventoryItem(id int) error {
	ctx := context.Background()

	return db.q.DeleteInventoryItem(ctx, int64(id));
}