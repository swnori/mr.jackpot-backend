package coupon

import "mr.jackpot-backend/model"


type StaffCouponService interface {
	IssueCoupon(coupon model.CouponInfo) (model.CouponInfo, error)
	GetIssuedCouponList() ([]model.CouponInfo, error)
	DeleteCoupon(couponid int) error
}

func (c *CouponManager) IssueCoupon(coupon model.CouponInfo) (model.CouponInfo, error) {
	return c.db.CreateCoupon(coupon)
}

func (c *CouponManager) GetIssuedCouponList() ([]model.CouponInfo, error) {
	return c.db.GetIssuedCouponList()
}

func (c *CouponManager) DeleteCoupon(couponid int) error {
	return c.db.DeleteCoupon(couponid)
}