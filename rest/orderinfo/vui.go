package orderinfo

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"mr.jackpot-backend/model"
	"mr.jackpot-backend/service/vui"
)

type VUIService interface {
	HandleVUIStep(c *gin.Context)
}

type VUIHandler struct {
	vui vui.VUIController
}

func (h *VUIHandler) HandleVUIStep(c *gin.Context) {
	request := model.OrderChoiceRequest{}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, request)
	}

	response, err := h.vui.HandleOrderChoice(request);
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return		
	}
	c.JSON(http.StatusOK, response)
}