package coupon

import (
	"mr.jackpot-backend/model"
)



type CustomerCouponProvider interface {
	GainCoupon(userid int, code string) (model.CouponString, error)
	GetCouponList(userid int) ([]model.CouponInfo, error)
}


func (c *CouponManager) GainCoupon(userid int, code string) (model.CouponString, error) {

	couponid, err := c.db.CheckCouponCodeMatch(code)
	if err != nil {
		return model.CouponString{}, err
	}

	if err := c.db.OwnCoupon(userid, couponid); err != nil {
		return model.CouponString{}, err
	}

	coupon, err := c.db.GetCouponInfo(couponid)
	if err != nil {
		return model.CouponString{}, err
	}

	return model.CouponString{
		Title: coupon.Title,
		Amount: coupon.Amount,
		Message: coupon.Message,
		CreatedAt: coupon.CreatedAt.Format("2020-07-30"),
		ExpiresAt: coupon.ExpiresAt.Format("2020-07-30"),
	}, nil
}

func (c *CouponManager) GetCouponList(userid int) ([]model.CouponInfo, error) {
	coupon, err := c.db.GetCouponListByID(userid)
	return coupon, err
}


