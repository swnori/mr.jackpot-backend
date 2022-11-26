package model

import "time"

// response
/*
type ClientInfo struct {
	ID        int       `json:"orderId"`
	StateID   int       `json:"stateId"`
	Member    bool      `json:"isMember"`
	Name      string    `json:"reserveName"`
	Address   string    `json:"address"`
	Phone     string    `json:"contact"`
	ReserveAt time.Time `json:"reserveDate"`
	OrderAt   time.Time `json:"orderDate"`
	Message   string    `json:"requestDetail"`
}

type PaymentInfo struct {
	Price       int    `json:"price"`
	CouponPrice int    `json:"couponPrice"`
	CouponName  string `json:"couponName"`
}

type OrderSummery struct {
	ID         int       `json:"id"`
	DinnerName string    `json:"dinnerName"`
	CreatedAt  time.Time `json:"createTime"`
	ReserveAt  time.Time `json:"reserveTime"`
	Price      int       `json:"price"`
}

type OrderInfo struct {
	ID        int      `json:"-"`
	Price     int      `json:"price"`
	CouponPrice int    `json:"couponPrice"`
	CouponName  string `json:"couponName"`
}

*/

type OrderRequest struct {
	Order
	Info OrderRequestInfo `json:"orderInfo"`
}

type OrderRequestInfo struct {
	OwnerID   int    `json:"-"`
	Name      string `json:"reserveName"`
	Phone     string `json:"call"`
	Address   string `json:"address"`
	Message   string `json:"requestDetail"`
	ReserveAt string `json:"reserveDate"`
	CouponID  int    `json:"couponId"`
	Price     int    `json:"price"`
}


type AllOrderInfo struct {
	ID        int    `json:"orderId"`
	OwnerID   int    `json:"-"`
	StateID   int    `json:"stateId"`

	Name      string `json:"reserveName"`
	Phone     string `json:"contact"`
	Address   string `json:"address"`
	Message   string `json:"requestDetail"`

	ReserveAt time.Time `json:"reserveDate"`
	CreatedAt time.Time `json:"createTime"`

	Price       int    `json:"price"`
	CouponPrice int    `json:"couponPrice"`
	CouponName  string `json:"couponName"`
}

type OrderResponse struct {
	Order
	AllOrderInfoResponse
}


type AllOrderInfoResponse struct {
	ID        int    `json:"orderId"`
	OwnerID   int    `json:"-"`
	StateID   int    `json:"stateId"`

	Name      string `json:"reserveName"`
	Phone     string `json:"contact"`
	Address   string `json:"address"`
	Message   string `json:"requestDetail"`

	ReserveAt string `json:"reserveDate"`
	CreatedAt string `json:"createTime"`

	Price       int    `json:"price"`
	CouponPrice int    `json:"couponPrice"`
	CouponName  string `json:"couponName"`	
}