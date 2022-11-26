package model

type Order struct {
	DinnerList []DinnerOrder `json:"dinnerList"`
}

type OrderResponse struct {
	Order
	AllOrderInfo 
}

type DinnerOrder struct {
	DinnerId int         `json:"dinnerId"`
	StyleId  int         `json:"styleId"`
	MenuList []MenuOrder `json:"menuList"`
}

type MenuOrder struct {
	MenuId   int   `json:"menuId"`
	Count    int   `json:"count"`
	OptionId []int `json:"optionId"`
}
