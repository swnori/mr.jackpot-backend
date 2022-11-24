package db

import (
	"context"

	"mr.jackpot-backend/model"
	"mr.jackpot-backend/utility/util"
)



type BoardLayer interface {
	GetDinnerList() ([]model.DinnerBoardItem, error)
	GetMenuList() ([]model.MenuBoardItem, error)
}


type BoardDB struct {
	DBAccessor
}

func NewBoardDB() *BoardDB {
	db := &BoardDB{}
	db.q = NewAccessor()

	return db
}

func (db *BoardDB) GetDinnerList() ([]model.DinnerBoardItem, error) {
	ctx := context.Background()

	DinnerBoardList, err := db.q.ReadDinnerEntity(ctx)

	dinnerBoardList := make([]model.DinnerBoardItem, 0, len(DinnerBoardList))

	for _, DinnerBoard := range DinnerBoardList {

		menuList32, err := db.q.ReadDinnersMenu(ctx, DinnerBoard.DinnerID)
		menuList := util.IntAll(menuList32)
		if err != nil {
			return nil, err
		}
		
		dinnerBoard := model.DinnerBoardItem{
			MenuList: menuList,
			Id: int(DinnerBoard.DinnerID),
			Name: DinnerBoard.Name,
			Price: int(DinnerBoard.Price),
		}

		dinnerBoardList = append(dinnerBoardList, dinnerBoard)
	}

	return dinnerBoardList, err
}



func (db *BoardDB) GetMenuBoard() ([]model.MenuBoardItem, error) {
	ctx := context.Background()

	MenuBoardList, err := db.q.ReadMenuEntity(ctx)

	menuBoardList := make([]model.MenuBoardItem, len(MenuBoardList))

	for _, MenuBoard := range MenuBoardList {

		optionModelList := make([]model.OptionModel, 0)

		if MenuBoard.Option1Name.Valid {
			OptionList, err := db.GetOption1Board(int(MenuBoard.MenuID))

			if err != nil {
				return nil, err
			}

			optionModel := model.OptionModel{
				Name: MenuBoard.Option1Name.String,
				OptionList: OptionList,
			}

			optionModelList = append(optionModelList, optionModel)
		}

		if MenuBoard.Option2Name.Valid {
			OptionList, err := db.GetOption2Board(int(MenuBoard.MenuID))

			if err != nil {
				return nil, err
			}

			optionModel := model.OptionModel{
				Name: MenuBoard.Option1Name.String,
				OptionList: OptionList,
			}

			optionModelList = append(optionModelList, optionModel)
		}

		menuBoard  := model.MenuBoardItem{
			Id: int(MenuBoard.MenuID),
			Price: int(MenuBoard.Price),
			OptionModelList: optionModelList,
		}

		menuBoardList = append(menuBoardList, menuBoard)
	}

	return menuBoardList, err
}



func (db *BoardDB) GetOption1Board(menuId int) ([]model.OptionBoardItem, error) {
	ctx := context.Background()

	Option1List, err := db.q.ReadOption1Entity(ctx, int32(menuId))

	optionList := make([]model.OptionBoardItem, 0, len(Option1List))

	for _, Option1 := range Option1List {
		option := model.OptionBoardItem{
			Id: int(Option1.OptionID),
			Name: Option1.Name,
			Price: int(Option1.Price),
		}

		optionList = append(optionList, option)
	}

	return optionList, err
}



func (db *BoardDB) GetOption2Board(menuId int) ([]model.OptionBoardItem, error) {
	ctx := context.Background()

	Option2List, err := db.q.ReadOption2Entity(ctx, int32(menuId))

	optionList := make([]model.OptionBoardItem, 0, len(Option2List))

	for _, Option1 := range Option2List {
		option := model.OptionBoardItem{
			Id: int(Option1.OptionID),
			Name: Option1.Name,
			Price: int(Option1.Price),
		}

		optionList = append(optionList, option)
	}

	return optionList, err
}