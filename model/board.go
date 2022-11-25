package model

type DinnerBoardItem struct {
	Id       int
	Name     string
	Price    int
	MenuList []DinnersMenu
}

type MenuBoardItem struct {
	Id         int
	Name       string
	Price      int
	MenuType   string
	OptionModelList []OptionModel
}

type DinnersMenu struct {
	MenuId int
	Count  int
}

type OptionModel struct {
	Name string
	OptionList []OptionBoardItem
}

type OptionBoardItem struct {
	Id    int
	Name  string
	Price int
}

type StyleBoardItem struct {
	Id    int
	Name  string
	Price int
	Desc  string
}

type OrderState struct {
	Id int
	State string
}



