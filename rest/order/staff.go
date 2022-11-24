package order

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


type StaffOrderService interface {
	GetAllOrderList(c *gin.Context)
	AcceptOrder(c *gin.Context)
	RejectOrder(c *gin.Context)
}



func (h *OrderHandler) AcceptOrder(c *gin.Context) {

	var order int

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	if err := h.order.FinishOrderStep(order); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	
	c.JSON(http.StatusOK, "")
}



func (h *OrderHandler) RejectOrder(c *gin.Context) {

	var orderid int
	if err := c.ShouldBindJSON(&orderid); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	if err := h.order.CeaseOrder(orderid); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	c.JSON(http.StatusOK, "")
}



func (h *OrderHandler) GetAllOrderList(c *gin.Context) {
	orderlist := h.order.GetAllOrderInfo()
	c.JSON(http.StatusOK, orderlist)
}
