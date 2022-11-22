package board

import "mr.jackpot-backend/model"

type MOCKDB interface {
	GetDinnerList() ([]model.DinnerBoardItem, error)
	GetMenuList() ([]model.MenuBoardItem, error)
	GetEntityList() ([]model.Action, error)
}



type OrderBoard struct {
	db MOCKDB

	DinnerList []model.DinnerBoardItem
	MenuList   []model.MenuBoardItem
	EntityList map[int]model.Action
}

func (b *OrderBoard) GetDBToEntity() (err error) {
	_, err = b.db.GetDinnerList()
	if err != nil {
		return
	}

	_, err = b.db.GetMenuList()
	if err != nil {
		return
	}
	
	_, err = b.db.GetEntityList()
	if err != nil {
		return
	}
	return
}

var Board *OrderBoard = &OrderBoard{}


func NewBoard() *OrderBoard {
	return Board
}

func Initialize() {
	if err := Board.GetDBToEntity(); err != nil {
		panic(err)
	}
}