package coupon

import (
	"time"

	"mr.jackpot-backend/model"
	"mr.jackpot-backend/utility/util"
)

type StaffCouponProvider interface {
	IssueCoupon(coupon model.CouponString) (model.CouponInfo, error)
	GetIssuedCouponList() ([]model.CouponInfo, error)
	DeleteCoupon(couponid int) error
}

func (c *CouponManager) IssueCoupon(coupon model.CouponString) (model.CouponInfo, error) {
	code := util.GetRandomString(16)

	expireTime, err := util.TimeParsor(coupon.ExpiresAt)
	if err != nil {
		return model.CouponInfo{}, err
	}
	return c.db.CreateCoupon(model.CouponInfo{
		Code:      code,
		Amount:    coupon.Amount,
		Title:     coupon.Title,
		Message:   coupon.Message,
		CreatedAt: time.Now(),
		ExpiresAt: expireTime,
	})
}

func (c *CouponManager) GetIssuedCouponList() ([]model.CouponInfo, error) {
	return c.db.GetIssuedCouponList()
}

func (c *CouponManager) DeleteCoupon(couponid int) error {
	return c.db.DeleteCoupon(couponid)
}
