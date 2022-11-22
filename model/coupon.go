package model

import "time"


type CouponInfo struct {
	ID      int
	Code    string
	Amount  int
	Title   string
	Message string
	CreatedAt  time.Time
	ExpiresAt time.Time
}

type CouponCreateRequest struct {
	Title   string
	Amount  int
	Message string
	CreatedAt  time.Time
	ExpiresAt time.Time

}