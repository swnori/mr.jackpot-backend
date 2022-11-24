package task

import "github.com/gin-gonic/gin"

type TaskService interface {
	GetAllTaskList(c *gin.Context)
	SetTaskNextStatus(c *gin.Context)
	SetTaskPreviousStatus(c *gin.Context)
}

type TaskHandler struct {
	
}

func NewTaskHandler() *TaskHandler {
	return &TaskHandler{}
}