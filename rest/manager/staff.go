package manager

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"mr.jackpot-backend/model"
)

type StaffManagerService interface {
	GetStaffList(c *gin.Context)
	RegisterStaff(c *gin.Context)
	UpdateStaffInfo(c *gin.Context)
}

func (h *ManagerHandler) GetStaffList(c *gin.Context) {
	stafflist, err := h.sm.GetAllUserInfo()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, stafflist)
}

func (h *ManagerHandler) RegisterStaff(c *gin.Context) {
	request := model.StaffInfo{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	if err := h.sm.CreateAccount(request); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "")
}

func (h *ManagerHandler) UpdateStaffInfo(c *gin.Context) {
	request := model.StaffInfo{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	if err := h.sm.UpdateUserInfo(); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "")
}
