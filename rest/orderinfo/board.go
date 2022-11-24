package orderinfo

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"mr.jackpot-backend/service/board"
)

type BoardService interface {
	GetOrderBoard(c *gin.Context)
	GetStateList(c *gin.Context)
}

type BoardHandler struct {
	board board.BoardController
}

func (h *BoardHandler) GetOrderBoard(c *gin.Context) {
	dinner, menu, style := h.board.GetOrderBoard()

	c.JSON(http.StatusOK, gin.H{
		"dinnerList": dinner,
		"menuList": menu,
		"styleList": style,
	})
}

func (h *BoardHandler) GetStateList(c *gin.Context) {
	statelist := h.board.GetOrderStateList()
	c.JSON(http.StatusOK, statelist)
}