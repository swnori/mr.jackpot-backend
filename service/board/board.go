package board

import (
	"mr.jackpot-backend/database/db"
	"mr.jackpot-backend/model"
)



type OrderBoard struct {
	db db.BoardLayer

	DinnerList []model.DinnerBoardItem
	MenuList   []model.MenuBoardItem
	StyleList  []model.StyleBoardItem
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
	
	//_, err = b.db.GetEntityList()
	//if err != nil {
	//	return
	//}
	return
}

var Board *OrderBoard = &OrderBoard{}


func NewBoard() *OrderBoard {
	return Board
}

func Initialize() {
	//Board.db = db.NewBoardDB()

	if err := Board.GetDBToEntity(); err != nil {
		panic(err)
	}
}