package board

import (
	"mr.jackpot-backend/model"
)

type BoardController interface {
	GetDinnerBoard() ([]model.DinnerBoardItem)
	GetMenuBoard() ([]model.MenuBoardItem)
}