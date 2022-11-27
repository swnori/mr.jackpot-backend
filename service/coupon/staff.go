package coupon

import (
	"mr.jackpot-backend/model"
	"mr.jackpot-backend/utility/util"
)

type StaffCouponProvider interface {
	IssueCoupon(coupon model.CouponInfo) (model.CouponInfo, error)
	GetIssuedCouponList() ([]model.CouponInfo, error)
	DeleteCoupon(couponid int) error
}

func (c *CouponManager) IssueCoupon(coupon model.CouponInfo) (model.CouponInfo, error) {
	code := util.GetRandomString(16)

	return c.db.CreateCoupon(model.CouponInfo{
		Code:      code,
		Amount:    coupon.Amount,
		Title:     coupon.Title,
		Message:   coupon.Message,
		ExpiresAt: coupon.ExpiresAt,
	})
}

func (c *CouponManager) GetIssuedCouponList() ([]model.CouponInfo, error) {
	couponlist, err := c.db.GetIssuedCouponList()

	couponList := make([]model.CouponInfo, 0)

	for _, coupon := range couponlist {
		couponList = append(couponList, model.CouponInfo{
			ID: coupon.ID,
			Code: coupon.Code,
			Amount: coupon.Amount,
			Title: coupon.Title,
			Message: coupon.Message,
			CreatedAt: coupon.CreatedAt,
			ExpiresAt: coupon.ExpiresAt,
		})
	}
	
	return couponList, err
}

func (c *CouponManager) DeleteCoupon(couponid int) error {
	return c.db.DeleteCoupon(couponid)
}
