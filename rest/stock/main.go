package stock

import (
	"github.com/gin-gonic/gin"
	"mr.jackpot-backend/service/stock"
)

type StockService interface {
	AddStockItem(c *gin.Context)
	UpdateStockItem(c *gin.Context)
	DeleteStockItem(c *gin.Context)
	GetAllStockList(c *gin.Context)
}

type StockHandler struct {
	stock stock.StockController
}

func NewStockHandler() *StockHandler {
	return &StockHandler{
		stock: stock.NewStock(),
	}
}
