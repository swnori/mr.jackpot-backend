package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"mr.jackpot-backend/database/orm"
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
	UseCoupon(userid, couponid int) error
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
	ctx := context.Background()

	return db.q.OwnCoupon(ctx, orm.OwnCouponParams{
		CouponID: int64(couponid),
		OwnerID:  int64(userid),
	})
}

func (db *CouponDB) GetCouponInfo(couponid int) (model.CouponInfo, error) {
	ctx := context.Background()

	coupon, err := db.q.GetCouponInfo(ctx, int64(couponid))
	return model.CouponInfo{
		ID:        int(coupon.CouponID),
		Title:     coupon.Title.String,
		Message:   coupon.Description.String,
		CreatedAt: coupon.CreatedAt,
		ExpiresAt: coupon.ExpiresAt,
	}, err
}

func (db *CouponDB) GetCouponListByID(userid int) ([]model.CouponInfo, error) {
	ctx := context.Background()

	couponList, err := db.q.GetCouponAvailable(ctx, int64(userid))
	fmt.Println(userid)
	fmt.Println(len(couponList))

	CouponList := make([]model.CouponInfo, 0);
	for _, coupon := range couponList {
		if time := time.Now().Sub(coupon.ExpiresAt).Seconds(); time > 0 {
			continue
		}

		CouponList = append(CouponList, model.CouponInfo{
			ID: int(coupon.CouponID),
			Code: coupon.Code,
			Amount: int(coupon.Amount),
			Title: coupon.Title.String,
			Message: coupon.Description.String,
			ExpiresAt: coupon.ExpiresAt,
		})
	}

	return CouponList, err
}

func (db *CouponDB) CreateCoupon(coupon model.CouponInfo) (model.CouponInfo, error) {
	ctx := context.Background()

	result, err := db.q.IssueCoupon(ctx, orm.IssueCouponParams{
		Code:   coupon.Code,
		Amount: int32(coupon.Amount),
		Title: sql.NullString{
			String: coupon.Title,
			Valid:  true,
		},
		Description: sql.NullString{
			String: coupon.Message,
			Valid:  true,
		},
		CreatedAt: coupon.CreatedAt,
		ExpiresAt: coupon.ExpiresAt,
	})

	if err != nil {
		return model.CouponInfo{}, err
	}

	couponID, err := result.LastInsertId()
	if err != nil {
		return model.CouponInfo{}, err
	}

	couponIssued, err := db.q.GetCouponInfo(ctx, couponID)

	return model.CouponInfo{
		ID: int(couponID),
		Amount: int(couponIssued.Amount),
		Code: couponIssued.Code,
		Title: couponIssued.Title.String,
		Message: couponIssued.Description.String,
		ExpiresAt: couponIssued.ExpiresAt,
	}, err
}

func (db *CouponDB) GetIssuedCouponList() ([]model.CouponInfo, error) {
	ctx := context.Background()

	couponList, err := db.q.GetCouponIssued(ctx)
	if err != nil {
		return nil, err
	}

	CouponList := make([]model.CouponInfo, 0)
	for _, coupon := range couponList {
		CouponList = append(CouponList, model.CouponInfo{
			ID:        int(coupon.CouponID),
			Amount:    int(coupon.Amount),
			Code:      coupon.Code,
			Title:     coupon.Title.String,
			Message:   coupon.Description.String,
			CreatedAt: coupon.CreatedAt,
			ExpiresAt: coupon.ExpiresAt,
		})
	}

	return CouponList, nil
}

func (db *CouponDB) DeleteCoupon(id int) error {
	ctx := context.Background()

	return db.q.DeleteCoupon(ctx, int64(id))
}


func (db *CouponDB) UseCoupon(userid, couponid int) error {
	ctx := context.Background()

	return db.q.UseCoupon(ctx, orm.UseCouponParams{
		CouponID: int64(couponid),
		OwnerID: int64(userid),
	})
}