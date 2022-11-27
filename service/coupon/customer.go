package coupon

import (
	"mr.jackpot-backend/model"
)



type CustomerCouponProvider interface {
	GainCoupon(userid int, code string) (model.CouponInfo, error)
	GetCouponList(userid int) ([]model.CouponInfo, error)
}


func (c *CouponManager) GainCoupon(userid int, code string) (model.CouponInfo, error) {

	couponid, err := c.db.CheckCouponCodeMatch(code)
	if err != nil {
		return model.CouponInfo{}, err
	}

	if err := c.db.OwnCoupon(userid, couponid); err != nil {
		return model.CouponInfo{}, err
	}

	coupon, err := c.db.GetCouponInfo(couponid)
	if err != nil {
		return model.CouponInfo{}, err
	}

	return model.CouponInfo{
		Title: coupon.Title,
		Amount: coupon.Amount,
		Message: coupon.Message,
		ExpiresAt: coupon.ExpiresAt,
	}, nil
}

func (c *CouponManager) GetCouponList(userid int) ([]model.CouponInfo, error) {
	couponlist, err := c.db.GetCouponListByID(userid)
	Couponlist := make([]model.CouponInfo, 0)
	for _, coupon := range couponlist {
		Couponlist = append(Couponlist, model.CouponInfo{
			ID: coupon.ID,
			Title: coupon.Title,
			Amount: coupon.Amount,
			Message: coupon.Message,
			ExpiresAt: coupon.ExpiresAt,
		})
	}
	return Couponlist, err
}


