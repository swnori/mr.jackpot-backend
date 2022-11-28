package order

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"mr.jackpot-backend/model"
	"mr.jackpot-backend/service/order"
)


type StaffOrderService interface {
	GetAllOrderSummary(c *gin.Context)
	GetOrderInfo(c *gin.Context)
	AcceptOrder(c *gin.Context)
	RejectOrder(c *gin.Context)
	StartOrder(c *gin.Context)
}

func (h *OrderHandler) AcceptOrder(c *gin.Context) {

	var order model.OrderID

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	if err := h.order.FinishOrderStep(order.OrdrID); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	
	c.JSON(http.StatusOK, "")
}



func (h *OrderHandler) RejectOrder(c *gin.Context) {

	var order model.OrderID

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	if err := h.order.CeaseOrder(order.OrdrID); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	c.JSON(http.StatusOK, "")
}



func (h *OrderHandler) GetAllOrderSummary(c *gin.Context) {
	orderlist, err := h.order.GetAllOrderSummary()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, orderlist)
}

func (h *OrderHandler) StartOrder(c *gin.Context) {
	var iorder model.OrderID

	if err := c.ShouldBindJSON(&iorder); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	if err := h.order.FinishOrderStep(iorder.OrdrID); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	for _, menu := range order.OrderManagers.Menu {
		order.OrderManagers.SetMenuNextStep(menu.ID)
	}
	
	c.JSON(http.StatusOK, "")
}

