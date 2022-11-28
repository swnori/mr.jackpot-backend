package model



type DinnerFormed struct {
	StateID   int   `json:"stateId"`
	OrderedID int   `json:"orderId"`
	ID        int   `json:"id"`
	DinnerID  int   `json:"dinnerId"`
	StyleID   int   `json:"styleId"`
	MenuList  []int `json:"menuList"`
}

type MenuFormed struct {
	StateID    int   `json:"stateId"`
	OrderedID  int   `json:"orderId"`
	DinnerID   int   `json:"dinnerId"`
	ID         int   `json:"id"`
	MenuID     int   `json:"menuId"`
	TypeID     int   `json:"typeId"`
	OptionList []int `json:"optionList"`
}

type OptionFormed struct {
	Name string
}


type OrderFormed struct {
	OrderID   int       `json:"orderId"`
	CreatedAt string    `json:"createdAt"`
	ReserveAt string    `json:"reserveAt"`
	Price     int       `json:"price"`
	DinnerList []string `json:"dinnerList"`
}

type OrderSummary struct {
	OrderID   int       `json:"orderId"`
	StateID   int       `json:"stateId"`
	ReserveAt string    `json:"reserveAt"`
	Price     int       `json:"price"`
	DinnerList []int `json:"dinnerList"`
}