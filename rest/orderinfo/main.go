package orderinfo

import "mr.jackpot-backend/service/vui"


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
	handler.vui = vui.VUI

	return handler
}