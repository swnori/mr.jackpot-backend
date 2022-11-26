package coupon

import (
	"mr.jackpot-backend/model"
)



type CustomerCouponProvider interface {
	GainCoupon(userid int, code string) (model.CouponString, error)
	GetCouponList(userid int) ([]model.CouponString, error)
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
		ExpiresAt: coupon.ExpiresAt.Format(model.TimeDayFormat),
	}, nil
}

func (c *CouponManager) GetCouponList(userid int) ([]model.CouponString, error) {
	couponlist, err := c.db.GetCouponListByID(userid)
	Couponlist := make([]model.CouponString, 0)
	for _, coupon := range couponlist {
		Couponlist = append(Couponlist, model.CouponString{
			ID: coupon.ID,
			Title: coupon.Title,
			Amount: coupon.Amount,
			Message: coupon.Message,
			ExpiresAt: coupon.ExpiresAt.Format(model.TimeDayFormat),
		})
	}
	return Couponlist, err
}


