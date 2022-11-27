package model



type DinnerFormed struct {
	ID       int
	DinnerID int
	MenuList []int
}

type MenuFormed struct {
	ID         int
	MenuID     int
	OptionList []int
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