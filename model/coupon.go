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
	CreatedAt time.Time
	ExpiresAt time.Time
}

type CouponString struct {
	Title     string
	Amount    int
	Message   string
	CreatedAt string
	ExpiresAt string
}
