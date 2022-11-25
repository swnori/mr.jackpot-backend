package db

import (
	"context"

	"mr.jackpot-backend/model"
)



type BoardLayer interface {
	GetDinnerList() ([]model.DinnerBoardItem, error)
	GetMenuList() ([]model.MenuBoardItem, error)
	GetStyleList() ([]model.StyleBoardItem, error)
	GetOrderStateList() ([]model.OrderState, error)
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
	if err != nil {
		return nil, err
	}

	dinnerBoardList := make([]model.DinnerBoardItem, 0)

	for _, DinnerBoard := range DinnerBoardList {

		menuInfoList, err := db.q.ReadDinnersMenu(ctx, DinnerBoard.DinnerID)
		dinnersMenu := make([]model.DinnersMenu, 0)

		for _, menuInfo := range menuInfoList {
			dinnersMenu = append(dinnersMenu, model.DinnersMenu{
				MenuId: int(menuInfo.MenuID),
				Count: int(menuInfo.DefaultCount),
			})
		}

		if err != nil {
			return nil, err
		}
		
		dinnerBoard := model.DinnerBoardItem{
			MenuList: dinnersMenu,
			Id: int(DinnerBoard.DinnerID),
			Name: DinnerBoard.Name,
			Price: int(DinnerBoard.Price),
		}

		dinnerBoardList = append(dinnerBoardList, dinnerBoard)
	}

	return dinnerBoardList, err
}



func (db *BoardDB) GetMenuList() ([]model.MenuBoardItem, error) {
	ctx := context.Background()

	MenuBoardList, err := db.q.ReadMenuEntity(ctx)
	if err != nil {
		return nil, err
	}

	menuBoardList := make([]model.MenuBoardItem, 0)

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
			Name: MenuBoard.Name,
			MenuType: MenuBoard.Typename,
		}

		menuBoardList = append(menuBoardList, menuBoard)
	}

	return menuBoardList, err
}



func (db *BoardDB) GetOption1Board(menuId int) ([]model.OptionBoardItem, error) {
	ctx := context.Background()

	Option1List, err := db.q.ReadOption1Entity(ctx, int32(menuId))

	optionList := make([]model.OptionBoardItem, 0)

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

	optionList := make([]model.OptionBoardItem, 0)

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



func (db *BoardDB) GetStyleList() ([]model.StyleBoardItem, error) {
	ctx := context.Background()
	stylelist, err := db.q.ReadStyleEntity(ctx)

	Stylelist := make([]model.StyleBoardItem, 0)
	for _, style := range stylelist {

		Stylelist = append(Stylelist, model.StyleBoardItem{
			Id: int(style.StyleID),
			Name: style.Name,
			Price: int(style.Price),
			Desc: style.Description,
		})
	}

	return Stylelist, err
}

func (db *BoardDB) GetOrderStateList() ([]model.OrderState, error) {
	ctx := context.Background()
	statelist, err := db.q.ReadOrderState(ctx)

	Statelist := make([]model.OrderState, 0)
	for _, state := range statelist {

		Statelist = append(Statelist, model.OrderState{
			Id: int(state.StateID),
			State: state.Name,
		})
	}


	return Statelist, err
}

