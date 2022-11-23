-- name: IssueCoupon :execresult
INSERT INTO coupon_issued (code, amount, title, description, created_at, expires_at)
VALUES (?, ?, ?, ?, ?, ?);

-- name: GetCouponIssued :many
SELECT *
FROM coupon_issued;

-- name: GetCouponMatched :one
SELECT coupon_id
FROM coupon_issued
WHERE code = (?);

-- name: GetCouponInfo :one
SELECT *
FROM coupon_issued
WHERE coupon_id = (?);

-- name: OwnCoupon :exec
INSERT INTO coupon_owned (coupon_id, owner_id)
VALUES (?, ?);

-- name: UseCoupon :exec
UPDATE coupon_owned
SET valid = TRUE
WHERE coupon_id = (?)
AND owner_id = (?);

-- name: GetCouponAvailable :many
SELECT issued.coupon_id, code, amount, title, description, created_at, expires_at
FROM coupon_owned owned, coupon_issued issued
WHERE owned.owner_id = (?)
AND owned.coupon_id = issued.coupon_id
AND owned.valid IS TRUE
AND coupon_issued.expires_at <= GETDATE;

-- name: DeleteCoupon :exec
DELETE FROM coupon_issued
WHERE coupon_id = (?);
