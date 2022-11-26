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
	GetOrderInfo(c *gin.Context)
}

func (h *OrderHandler) GetOrderInfo(c *gin.Context) {
	userid := c.Keys["userid"].(int)

	order, err := h.order.GetOrderInfo(userid);
	if err != nil {
		c.JSON(http.StatusOK, "no order")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"orderinfo": order.AllOrderInfo,
		"order": order.Order,
	})

}

func (h *OrderHandler) CreateOrder(c *gin.Context) {

	var request model.OrderRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	userid := c.Keys["userid"].(int)

	if err := h.order.CreateOrder(userid, request.Info, request.Order); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "")
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

	orderState, err := h.order.GetOrderState(orderid)
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
