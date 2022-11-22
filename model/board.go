package model

type DinnerBoardItem struct {
	Id       int
	Name     string
	Price    int
	MenuList []int
}

type MenuBoardItem struct {
	Id         int
	Name       string
	Price      int
	OptionModelList []OptionModel
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
}




