package model

type DinnerBoardItem struct {
	Id       int           `json:"id"`
	Name     string        `json:"name"`
	Price    int           `json:"price"`
	MenuList []DinnersMenu `json:"mainDish"`
}

type MenuBoardItem struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	MenuType   string `json:"type"`
	OptionModelList []OptionModel `json:"option"`
}

type DinnersMenu struct {
	MenuId int `json:"menuId"`
	Count  int `json:"count"`
}

type OptionModel struct {
	Name string `json:"name"`
	OptionList []OptionBoardItem
}

type OptionBoardItem struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type StyleBoardItem struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Desc  string `json:"desc"`
}