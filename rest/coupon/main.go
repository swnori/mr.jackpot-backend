package coupon

import (
	"mr.jackpot-backend/service/coupon"
)

type CouponService interface {
	CustomerCouponService
	StaffCouponService
}

type CouponHandler struct {
	c coupon.CouponProvider
}

func NewCouponHandler() *CouponHandler {
	return &CouponHandler{
		c: coupon.NewCouponManager(),
	}
}