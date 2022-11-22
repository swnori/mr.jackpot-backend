package coupon

import "mr.jackpot-backend/database/db"


type CouponService interface {
	StaffCouponService
	CustomerCouponService
}


type CouponManager struct {
	db db.CouponLayer
}

func NewCouponManager() *CouponManager {
	manager := CouponManager{}
	manager.db = db.NewCouponDB()

	return &manager
}