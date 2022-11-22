package model

import "time"


type CouponInfo struct {
	Title   string
	Message string
	ExpiresAt time.Time
	AvailsAt  time.Time
}