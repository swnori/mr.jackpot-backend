package orderinfo

import (

	"github.com/gin-gonic/gin"
)


type OrderHistoryService interface {
	GetOrderHistory(c *gin.Context)
}

type OrderHistoryHandler struct {

}

func (h *OrderHistoryHandler) GetOrderHistory(c *gin.Context) {

}
