package model




type DinnerStaff struct {
	DinnerId        int  `json:"dinnerId"`
	OrderedDinnerId int  `json:"id"`
	StateId         int  `json:"stateId"`

	StyleId  int         `json:"styleId"`
	MenuList []MenuOrder `json:"menuList"`
}

type MenuStaff struct {
	MenuId   int   `json:"menuId"`
	Count    int   `json:"count"`

	OrderedMenuId int  `json:"id"`
	StateId       int  `json:"stateId"`
	OptionId []int `json:"option"`
}

