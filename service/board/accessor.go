package board

import "mr.jackpot-backend/model"


type BoardProvider struct {
	OrderBoard
}

func (b *BoardProvider) GetBoard() () {
}



func (b *BoardProvider) GetStateList() []model.MenuBoardItem {
	return b.MenuList
}


func (b *BoardProvider) TransferToEntity(entityId int) (model.Action) {
	return b.EntityList[entityId]
}


