package manager

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"mr.jackpot-backend/model"
)

type CustomerManagerService interface {
	GetCustomerList(c *gin.Context)
	GetPersonalInfo(c *gin.Context)
	UpdatePersonalInfo(c *gin.Context)
}

func (h *ManagerHandler) GetCustomerList(c *gin.Context) {
	customerList, err := h.cm.GetAllUserInfo()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, customerList)
}

func (h *ManagerHandler) GetPersonalInfo(c *gin.Context) {
	userid := c.Keys["userid"].(int)

	personal, err := h.cm.GetPersonalnfo(userid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, personal)
}

func (h *ManagerHandler) UpdatePersonalInfo(c *gin.Context) {
	userid := c.Keys["userid"].(int)

	var personal model.PersonalInfo

	if err := c.ShouldBindJSON(&personal); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	if err := h.cm.UpdateUserInfo(userid, personal); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, personal)
}
