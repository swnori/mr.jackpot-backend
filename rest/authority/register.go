package authority

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"mr.jackpot-backend/model"
	"mr.jackpot-backend/utility/util"
)


type RegisterService interface {
	Register(c *gin.Context)
	Unregister(c *gin.Context)
}



func (h *CustomerAuthHandler) Register(c *gin.Context) {
	var request model.CustomerRegister

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	if err := h.m.CreateAccound(request); err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	c.JSON(http.StatusOK, "")
}

func (h *CustomerAuthHandler) Unregister(c *gin.Context) {
	userid := c.Keys["userid"].(int)

	if err := h.m.RemoveAccount(userid); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "")
}


func (h *VisitorAuthHandler) Register() (string, error) {
	identifier := util.GetRandomString(16)

	_, err := h.m.CreateAccount(identifier)
	return identifier, err
}


func (h *VisitorAuthHandler) Unregister(c *gin.Context) {

}