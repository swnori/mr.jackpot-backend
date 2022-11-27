package orderinfo

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"mr.jackpot-backend/database/db"
)


type OrderHistoryService interface {
	GetOrderHistory(c *gin.Context)
}

type OrderHistoryHandler struct {
	db db.OrderDB
}

func NewOrderHistoryHandler() *OrderHistoryHandler {
	return &OrderHistoryHandler{
		db: *db.NewOrderDB(),
	}
}

func (h *OrderHistoryHandler) GetOrderHistory(c *gin.Context) {
	userid := c.Keys["userid"].(int)

	orderlist, err := h.db.GetOrderHisory(userid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, orderlist)
}
