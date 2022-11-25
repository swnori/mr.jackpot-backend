package coupon

import "mr.jackpot-backend/model"



type CustomerCouponProvider interface {
	GainCoupon(userid int, code string) (model.CouponInfo, error)
	GetCouponList(userid int) ([]model.CouponInfo, error)
}


func (c *CouponManager) GainCoupon(userid int, code string) (model.CouponInfo, error) {
	coupon := model.CouponInfo{}

	couponid, err := c.db.CheckCouponCodeMatch(code)
	if err != nil {
		return coupon, err
	}

	if err := c.db.OwnCoupon(userid, couponid); err != nil {
		return coupon, err
	}

	return c.db.GetCouponInfo(couponid)
}

func (c *CouponManager) GetCouponList(userid int) ([]model.CouponInfo, error) {
	return c.db.GetCouponListByID(userid)
}


