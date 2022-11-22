package coupon

import (
	"time"

	"mr.jackpot-backend/model"
	"mr.jackpot-backend/utility/util"
)

type StaffCouponService interface {
	IssueCoupon(coupon model.CouponCreateRequest) (model.CouponInfo, error)
	GetIssuedCouponList() ([]model.CouponInfo, error)
	DeleteCoupon(couponid int) error
}

func (c *CouponManager) IssueCoupon(coupon model.CouponCreateRequest) (model.CouponInfo, error) {
	code := util.GetRandomString(16)

	return c.db.CreateCoupon(model.CouponInfo{
		Code:      code,
		Title:     coupon.Title,
		Message:   coupon.Message,
		CreatedAt: time.Now(),
		ExpiresAt: coupon.ExpiresAt,
	})
}

func (c *CouponManager) GetIssuedCouponList() ([]model.CouponInfo, error) {
	return c.db.GetIssuedCouponList()
}

func (c *CouponManager) DeleteCoupon(couponid int) error {
	return c.db.DeleteCoupon(couponid)
}
