package manager

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomerManagerService interface {
	GetCustomerList(c *gin.Context)
}

func (h *ManagerHandler) GetCustomerList(c *gin.Context) {
	customerList, err := h.cm.GetAllUserInfo()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, customerList)
}

