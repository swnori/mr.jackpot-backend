// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: coupon.sql

package orm

import (
	"context"
	"database/sql"
	"time"
)

const createCoupon = `-- name: CreateCoupon :exec
INSERT INTO coupon_owned (coupon_id, owner_id)
VALUES (?, ?)
`

type CreateCouponParams struct {
	CouponID int64
	OwnerID  int64
}

func (q *Queries) CreateCoupon(ctx context.Context, arg CreateCouponParams) error {
	_, err := q.db.ExecContext(ctx, createCoupon, arg.CouponID, arg.OwnerID)
	return err
}

const getCouponAvailable = `-- name: GetCouponAvailable :many
SELECT (amount, title, description, created_at, expires_at)
FROM coupon_owned owned, coupon_issued issued
WHERE owned.owner_id = (?)
AND owned.valid IS TRUE
AND owned.coupon_id = issued.coupon_id
`

func (q *Queries) GetCouponAvailable(ctx context.Context, ownerID int64) ([]interface{}, error) {
	rows, err := q.db.QueryContext(ctx, getCouponAvailable, ownerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []interface{}
	for rows.Next() {
		var column_1 interface{}
		if err := rows.Scan(&column_1); err != nil {
			return nil, err
		}
		items = append(items, column_1)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCouponMatched = `-- name: GetCouponMatched :one
SELECT coupon_id
FROM coupon_issued
WHERE code = (?)
`

func (q *Queries) GetCouponMatched(ctx context.Context, code string) (int64, error) {
	row := q.db.QueryRowContext(ctx, getCouponMatched, code)
	var coupon_id int64
	err := row.Scan(&coupon_id)
	return coupon_id, err
}

const issueCoupon = `-- name: IssueCoupon :exec
INSERT INTO coupon_issued (code, amount, title, description, created_at, expires_at)
VALUES (?, ?, ?, ?, ?, ?)
`

type IssueCouponParams struct {
	Code        string
	Amount      int32
	Title       sql.NullString
	Description sql.NullString
	CreatedAt   time.Time
	ExpiresAt   time.Time
}

func (q *Queries) IssueCoupon(ctx context.Context, arg IssueCouponParams) error {
	_, err := q.db.ExecContext(ctx, issueCoupon,
		arg.Code,
		arg.Amount,
		arg.Title,
		arg.Description,
		arg.CreatedAt,
		arg.ExpiresAt,
	)
	return err
}

const useCoupon = `-- name: UseCoupon :exec
UPDATE coupon_owned
SET valid = TRUE
WHERE coupon_id = (?)
AND owner_id = (?)
`

type UseCouponParams struct {
	CouponID int64
	OwnerID  int64
}

func (q *Queries) UseCoupon(ctx context.Context, arg UseCouponParams) error {
	_, err := q.db.ExecContext(ctx, useCoupon, arg.CouponID, arg.OwnerID)
	return err
}
