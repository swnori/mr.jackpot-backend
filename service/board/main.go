package board

import (
	"mr.jackpot-backend/database/db"
	"mr.jackpot-backend/model"
)

type BoardController interface {
	GetOrderBoard() ([]model.DinnerBoardItem, []model.MenuBoardItem, []model.StyleBoardItem)
	GetOrderStateList() ([]model.OrderState)
}

var Board = &OrderBoard{}

func NewBoard() *OrderBoard {
	return Board
}

func Initialize() error {
	Board.db = db.NewBoardDB()
	Board.GetAllEntity()

	return nil
}



