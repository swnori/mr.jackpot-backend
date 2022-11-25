package model

import "time"




type PaymentInfo struct {
	Price int
	CouponPrice int
	CouponName int
}

type DeliveryInfo struct {
	ID      int       `json:"-"`
	Name    string    `json:"name"`
	Time    time.Time `json:""`
	Phone   string    `json:"phone"`
	Address string    `json:"address"`
	Message string
}

type ClientInfo struct {
	
}