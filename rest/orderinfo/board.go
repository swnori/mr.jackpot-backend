package orderinfo

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"mr.jackpot-backend/service/board"
)

type BoardService interface {
	GetDinnerBoard(c *gin.Context)
	GetMenuBoard(c *gin.Context)
	GetStyleBoard(c *gin.Context)
	GetStateList(c *gin.Context)
}

type BoardHandler struct {
	board board.BoardController
}

func (h *BoardHandler) GetDinnerBoard(c *gin.Context) {
	dinnerBoard := h.board.GetDinnerBoard()
	c.JSON(http.StatusOK, dinnerBoard)
}

func (h *BoardHandler) GetMenuBoard(c *gin.Context) {
	menuBoard := h.board.GetMenuBoard()
	c.JSON(http.StatusOK, menuBoard)
}

func (h *BoardHandler) GetStyleBoard(c *gin.Context) {
	//dinnerBoard := h.board.GetStyleBoard()
	//c.JSON(http.StatusOK, dinnerBoard)
}

func (h *BoardHandler) GetStateList(c *gin.Context) {
	//dinnerBoard := h.board.GetStateList()
	//c.JSON(http.StatusOK, dinnerBoard)
}
