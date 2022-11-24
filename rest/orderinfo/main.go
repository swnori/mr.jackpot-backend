package orderinfo

import (
	"mr.jackpot-backend/service/board"
	"mr.jackpot-backend/service/vui"
)


type OrderInfoService interface {
	BoardService
	OrderHistoryService
	VUIService
}

type OrderInfoHandler struct {
	VUIHandler
	OrderHistoryHandler
	BoardHandler
}

func NewOrderInfoHandler() *OrderInfoHandler {
	handler := &OrderInfoHandler{}
	handler.vui = vui.NewVUIAccessor()
	handler.board = board.NewBoard()

	return handler
}