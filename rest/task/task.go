package task

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"mr.jackpot-backend/service/order"
)

func (h *TaskHandler) GetAllTaskList(c *gin.Context) {
	c.JSON(http.StatusOK, "preparing...")
}

func (h *TaskHandler) SetTaskNextStatus(c *gin.Context) {
	c.JSON(http.StatusOK, "preparing...")
}

func (h *TaskHandler) SetTaskPreviousStatus(c *gin.Context) {
	c.JSON(http.StatusOK, "preparing...")
}

func (h *TaskHandler) GetTaskListByRole(c *gin.Context) {
	role := c.Keys["role"].(string)
	
	switch (role) {
	case "cook":
		idlist, err :=  order.OrderManagers.GetAllMenuInfo()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		}
		c.JSON(http.StatusOK, idlist)
		break

	case "styler":
		idlist, err :=  order.OrderManagers.GetAllDinnerInfo()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		}
		c.JSON(http.StatusOK, idlist)
		break

	case "delivery":
		break
	default:
		break
	}

	c.JSON(http.StatusUnprocessableEntity, "no role")
}
