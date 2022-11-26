package model

import "time"

type CouponCode struct {
	Code string `json:"code"`
}

type CouponInfo struct {
	ID        int
	Code      string
	Amount    int
	Title     string
	Message   string
	ExpiresAt time.Time
}

type CouponString struct {
	ID        int
	Code      string
	Title     string
	Amount    int
	Message   string
	ExpiresAt string
}
