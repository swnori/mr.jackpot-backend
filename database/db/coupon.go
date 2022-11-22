package db

import (
	"context"

	"mr.jackpot-backend/model"
)



type CouponLayer interface {
	CheckCouponCodeMatch(string) (int, error)
	OwnCoupon(userid, couponid int) error
	GetCouponInfo(couponid int) (model.CouponInfo, error)
	GetCouponListByID(userid int) ([]model.CouponInfo, error)

	CreateCoupon(model.CouponInfo) (model.CouponInfo, error)
	GetIssuedCouponList() ([]model.CouponInfo, error)
	DeleteCoupon(int) error
}

type CouponDB struct {
	DBAccessor
}

func NewCouponDB() *CouponDB {
	db := &CouponDB{}
	db.q = NewAccessor()

	return db
}



func (db *CouponDB) CheckCouponCodeMatch(code string) (int, error) {
	ctx := context.Background()

	couponid, err := db.q.GetCouponMatched(ctx, code)
	if err != nil {
		return 0, err
	}
	return int(couponid), nil
}

func (db *CouponDB) OwnCoupon(userid, couponid int) error {

}
func (db *CouponDB) GetCouponInfo(couponid int) (model.CouponInfo, error) {

}
func (db *CouponDB) GetCouponListByID(userid int) ([]model.CouponInfo, error) {

}

func (db *CouponDB) CreateCoupon(model.CouponInfo) (model.CouponInfo, error) {
	

}
func (db *CouponDB) GetIssuedCouponList() ([]model.CouponInfo, error) {

}

func (db *CouponDB) DeleteCoupon(int) error {
}