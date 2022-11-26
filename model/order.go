package model

type Order struct {
	DinnerList []DinnerOrder
}

type OrderRequest struct {
	Order
	Info OrderRequestInfo
}

type OrderResponse struct {
	Order
	AllOrderInfo
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
