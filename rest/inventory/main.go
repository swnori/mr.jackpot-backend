package inventory

import (
	"github.com/gin-gonic/gin"
	"mr.jackpot-backend/service/inventory"
)





type InventoryService interface {
	AddInventoryItem(c *gin.Context)
	UpdateInventoryItem(c *gin.Context)
	DeleteInventoryItem(c *gin.Context)
	GetAllInventoryList(c *gin.Context)
}

type InventoryHandler struct {
	inventory inventory.InventoryController
}


func NewInventoryHandler() *InventoryHandler {
	return &InventoryHandler{
		inventory: inventory.NewInventory(),
	}
}
