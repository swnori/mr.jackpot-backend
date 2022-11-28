package task

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"mr.jackpot-backend/model"
	"mr.jackpot-backend/service/order"
)

func (h *TaskHandler) GetAllTaskList(c *gin.Context) {
	c.JSON(http.StatusOK, "preparing...")
}

func (h *TaskHandler) SetTaskNextStatus(c *gin.Context) {
	role := c.Keys["role"].(string)
	task := model.Task{}

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	switch (role) {
	case "cook":
		order.OrderManagers.SetMenuNextStep(task.ID)
		c.JSON(http.StatusOK, "")
		return

	case "styler":
		order.OrderManagers.SetDinnerNextStep(task.ID)
		c.JSON(http.StatusOK, "")
		return

	case "delivery":
		order.OrderManagers.FinishOrderStep(task.ID)
		c.JSON(http.StatusOK, "")
		return

	default:
		state := order.OrderManagers.Orders[task.ID].GetOrderState()
		if state == 4 || state == 5 {
			order.OrderManagers.FinishOrderStep(task.ID)
		} else {
			c.JSON(400, "")
		}
		break
	}
}


func (h *TaskHandler) SetTaskPreviousStatus(c *gin.Context) {
	c.JSON(http.StatusOK, "preparing...")
}

func (h *TaskHandler) GetTaskListByRole(c *gin.Context) {
	role := c.Keys["role"].(string)

	switch (role) {
	case "cook":
		c.JSON(http.StatusOK, order.OrderManagers.Menu)
		return
	case "styler":
		c.JSON(http.StatusOK, order.OrderManagers.Dinner)
		break
	case "delivery":
		c.JSON(http.StatusOK, "")
		break
	default:
		break
	}

	c.JSON(http.StatusUnprocessableEntity, "no role")
}
