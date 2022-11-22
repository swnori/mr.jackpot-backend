package model



type DeliveryInfo struct {
	PersonalInfo
	Message string
}

type Order struct {
	Dinner []DinnerOrder
}

type OrderRequest struct {
	Order
	DeliveryInfo 
}

type DinnerOrder struct {
	DinnerId int
	StyleId  int
	MenuList []MenuOrder
}

type MenuOrder struct {
	MenuId   int
	Count    int
	OptionId []int
}