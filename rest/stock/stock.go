package stock

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"mr.jackpot-backend/model"
)

func (h *StockHandler) AddStockItem(c *gin.Context) {

	item := model.StockItem{}
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	if err := h.stock.AddStockItem(item.Name, item.Unit); err != nil {
		c.JSON(http.StatusBadRequest, "stock already exists")
		return
	}

	c.JSON(http.StatusOK, "")
}

func (h *StockHandler) UpdateStockItem(c *gin.Context) {

	item := model.StockItem{}
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	if err := h.stock.UpdateStockItem(item.ID, item.Count); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "")
}

func (h *StockHandler) DeleteStockItem(c *gin.Context) {

	item := model.StockItem{}
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	if err := h.stock.DeleteStockItem(item.ID); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "")
}

func (h *StockHandler) GetAllStockList(c *gin.Context) {
	stockList := h.stock.GetAllStockList()
	c.JSON(http.StatusOK, stockList)
}
