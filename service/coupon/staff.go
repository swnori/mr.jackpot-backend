package coupon

import (
	"time"

	"mr.jackpot-backend/model"
	"mr.jackpot-backend/utility/util"
)

type StaffCouponProvider interface {
	IssueCoupon(coupon model.CouponString) (model.CouponString, error)
	GetIssuedCouponList() ([]model.CouponString, error)
	DeleteCoupon(couponid int) error
}

func (c *CouponManager) IssueCoupon(coupon model.CouponString) (model.CouponString, error) {
	code := util.GetRandomString(16)

	expireTime, err := time.Parse(model.TimeSecondFormat, coupon.ExpiresAt)

	if err != nil {
		return model.CouponString{}, err
	}
	couponid, err := c.db.CreateCoupon(model.CouponInfo{
		Code:      code,
		Amount:    coupon.Amount,
		Title:     coupon.Title,
		Message:   coupon.Message,
		ExpiresAt: expireTime,
	})

	return model.CouponString{
		ID: couponid,
		Code:      code,
		Amount:    coupon.Amount,
		Title:     coupon.Title,
		Message:   coupon.Message,
		ExpiresAt: expireTime.Format(model.TimeDayFormat),
	}, err
}

func (c *CouponManager) GetIssuedCouponList() ([]model.CouponString, error) {
	couponlist, err := c.db.GetIssuedCouponList()

	couponList := make([]model.CouponString, 0)

	for _, coupon := range couponlist {
		couponList = append(couponList, model.CouponString{
			ID: coupon.ID,
			Amount: coupon.Amount,
			Title: coupon.Title,
			Message: coupon.Message,
			ExpiresAt: coupon.ExpiresAt.Format(model.TimeDayFormat),
		})
	}
	
	return couponList, err
}

func (c *CouponManager) DeleteCoupon(couponid int) error {
	return c.db.DeleteCoupon(couponid)
}
