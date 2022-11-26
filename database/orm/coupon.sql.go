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

const deleteCoupon = `-- name: DeleteCoupon :exec

DELETE FROM coupon_issued
WHERE coupon_id = (?)
`

// AND issued.expires_at <= NOW();
func (q *Queries) DeleteCoupon(ctx context.Context, couponID int64) error {
	_, err := q.db.ExecContext(ctx, deleteCoupon, couponID)
	return err
}

const getCouponAvailable = `-- name: GetCouponAvailable :many
SELECT issued.coupon_id, code, amount, title, description, expires_at
FROM coupon_owned owned, coupon_issued issued
WHERE owned.owner_id = (?)
AND owned.coupon_id = issued.coupon_id
AND owned.valid IS TRUE
`

func (q *Queries) GetCouponAvailable(ctx context.Context, ownerID int64) ([]CouponIssued, error) {
	rows, err := q.db.QueryContext(ctx, getCouponAvailable, ownerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []CouponIssued
	for rows.Next() {
		var i CouponIssued
		if err := rows.Scan(
			&i.CouponID,
			&i.Code,
			&i.Amount,
			&i.Title,
			&i.Description,
			&i.ExpiresAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCouponInfo = `-- name: GetCouponInfo :one
SELECT coupon_id, code, amount, title, description, expires_at
FROM coupon_issued
WHERE coupon_id = (?)
`

func (q *Queries) GetCouponInfo(ctx context.Context, couponID int64) (CouponIssued, error) {
	row := q.db.QueryRowContext(ctx, getCouponInfo, couponID)
	var i CouponIssued
	err := row.Scan(
		&i.CouponID,
		&i.Code,
		&i.Amount,
		&i.Title,
		&i.Description,
		&i.ExpiresAt,
	)
	return i, err
}

const getCouponIssued = `-- name: GetCouponIssued :many
SELECT coupon_id, code, amount, title, description, expires_at
FROM coupon_issued
`

func (q *Queries) GetCouponIssued(ctx context.Context) ([]CouponIssued, error) {
	rows, err := q.db.QueryContext(ctx, getCouponIssued)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []CouponIssued
	for rows.Next() {
		var i CouponIssued
		if err := rows.Scan(
			&i.CouponID,
			&i.Code,
			&i.Amount,
			&i.Title,
			&i.Description,
			&i.ExpiresAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
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

const issueCoupon = `-- name: IssueCoupon :execresult
INSERT INTO coupon_issued (code, amount, title, description, expires_at)
VALUES (?, ?, ?, ?, ?)
`

type IssueCouponParams struct {
	Code        string
	Amount      int32
	Title       sql.NullString
	Description sql.NullString
	ExpiresAt   time.Time
}

func (q *Queries) IssueCoupon(ctx context.Context, arg IssueCouponParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, issueCoupon,
		arg.Code,
		arg.Amount,
		arg.Title,
		arg.Description,
		arg.ExpiresAt,
	)
}

const ownCoupon = `-- name: OwnCoupon :exec
INSERT INTO coupon_owned (coupon_id, owner_id)
VALUES (?, ?)
`

type OwnCouponParams struct {
	CouponID int64
	OwnerID  int64
}

func (q *Queries) OwnCoupon(ctx context.Context, arg OwnCouponParams) error {
	_, err := q.db.ExecContext(ctx, ownCoupon, arg.CouponID, arg.OwnerID)
	return err
}

const useCoupon = `-- name: UseCoupon :exec
UPDATE coupon_owned
SET valid = FALSE
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
