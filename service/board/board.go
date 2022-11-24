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
	OrderStateList []model.OrderState
}



func (b *OrderBoard) GetAllEntity() (err error) {
	if b.DinnerList, err = b.db.GetDinnerList(); err != nil {
		return
	}

	if b.MenuList, err = b.db.GetMenuList(); err != nil {
		return
	}

	if b.StyleList, err = b.db.GetStyleList(); err != nil {
		return
	}

	if b.OrderStateList, err = b.db.GetOrderStateList(); err != nil {
		return
	}

	return
}



func (b *OrderBoard) GetOrderBoard() ([]model.DinnerBoardItem, []model.MenuBoardItem, []model.StyleBoardItem) {
	return b.DinnerList, b.MenuList, b.StyleList
}

func (b *OrderBoard) GetOrderStateList() ([]model.OrderState) {
	return b.OrderStateList
}
