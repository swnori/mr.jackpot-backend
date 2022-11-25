package order

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"mr.jackpot-backend/model"
)

type CustomerOrderService interface {
	CreateOrder(c *gin.Context)
	CancleOrder(c *gin.Context)
	RequestCollecting(c *gin.Context)
}


func (h *OrderHandler) CreateOrder(c *gin.Context) {

	order := model.OrderRequest{}
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	userid := c.Keys["userid"].(int)

	if err := h.order.CreateOrder(userid, order.Order, order.DeliveryInfo); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	
	c.JSON(http.StatusOK, order)
}



func (h *OrderHandler) CancleOrder(c *gin.Context) {

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


func (h *OrderHandler) RequestCollecting(c *gin.Context) {
	
	var orderid int
	if err := c.ShouldBindJSON(&orderid); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	orderState, err := h.order.GetOrderState(orderid);
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	if orderState != model.StateDelivered {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := h.order.FinishOrderStep(orderid); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	c.JSON(http.StatusOK, "")
}
