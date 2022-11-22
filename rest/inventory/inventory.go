package inventory

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"mr.jackpot-backend/model"
)



func (h *InventoryHandler) AddInventoryItem(c *gin.Context) {

	item := model.InventoryItem{}
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	
	if err := h.inventory.AddInventoryItem(item.Name); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "")
}

func (h *InventoryHandler) UpdateInventoryItem(c *gin.Context) {

	item := model.InventoryItem{}
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	if err := h.inventory.UpdateInventoryItem(item.ID, item.Count); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "")
}

func (h *InventoryHandler) DeleteInventoryItem(c *gin.Context) {

	item := model.InventoryItem{}
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	if err := h.inventory.DeleteInventoryItem(item.ID); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "")
}



func (h *InventoryHandler) GetAllInventoryList(c *gin.Context) {

	inventoryList, err := h.inventory.GetAllInventoryList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, inventoryList)
}