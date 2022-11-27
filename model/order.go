package model

type Order struct {
	DinnerList []DinnerOrder `json:"dinnerList"`
}

type OrderID struct {
	OrdrID int `json:"orderId"`
}


type DinnerOrder struct {
	DinnerId int         `json:"dinnerId"`
	OrderedDinnerId int  `json:"id"`
	StateId         int  `json:"stateId"`

	StyleId  int         `json:"styleId"`
	MenuList []MenuOrder `json:"menuList"`
}

type MenuOrder struct {
	MenuId   int   `json:"menuId"`
	Count    int   `json:"count"`
	OptionId []int `json:"option"`
}