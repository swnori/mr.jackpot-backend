package board

import "mr.jackpot-backend/model"


type BoardProvider struct {
	OrderBoard
}


func (b *BoardProvider) GetDinnerBoard() []model.DinnerBoardItem {
	return b.DinnerList
}

func (b *BoardProvider) GetMenuBoard() []model.MenuBoardItem {
	return b.MenuList
}

func (b *BoardProvider) TransferToEntity(entityId int) (model.Action) {
	return b.EntityList[entityId]
}
