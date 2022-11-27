package task

import (
	"net/http"

	"github.com/gin-gonic/gin"
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


