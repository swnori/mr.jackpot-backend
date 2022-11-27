package model



type DinnerFormed struct {
	Name string
	Desc string
	Price string
	MenuFormed []MenuFormed
}

type MenuFormed struct {
	Name string
	Price string	
	Optionlist []OptionFormed
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
	DinnerList []string `json:"dinnerList"`
}