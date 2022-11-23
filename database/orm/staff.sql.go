// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: staff.sql

package orm

import (
	"context"
	"database/sql"
)

const createStaffAccount = `-- name: CreateStaffAccount :execresult
INSERT INTO staff (role_id, name)
VALUES (?, ?)
`

type CreateStaffAccountParams struct {
	RoleID int32
	Name   string
}

func (q *Queries) CreateStaffAccount(ctx context.Context, arg CreateStaffAccountParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createStaffAccount, arg.RoleID, arg.Name)
}

const createStaffAuth = `-- name: CreateStaffAuth :exec
INSERT INTO staff_auth (staff_id, code)
VALUES (?, ?)
`

type CreateStaffAuthParams struct {
	StaffID int64
	Code    string
}

func (q *Queries) CreateStaffAuth(ctx context.Context, arg CreateStaffAuthParams) error {
	_, err := q.db.ExecContext(ctx, createStaffAuth, arg.StaffID, arg.Code)
	return err
}

const getAllStaffInfo = `-- name: GetAllStaffInfo :many
SELECT staff_id, status, role.tag, staff.name, score
FROM staff, role
WHERE staff.role_id = role.role_id
AND staff.status = TRUE
`

type GetAllStaffInfoRow struct {
	StaffID int64
	Status  bool
	Tag     string
	Name    string
	Score   int32
}

func (q *Queries) GetAllStaffInfo(ctx context.Context) ([]GetAllStaffInfoRow, error) {
	rows, err := q.db.QueryContext(ctx, getAllStaffInfo)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAllStaffInfoRow
	for rows.Next() {
		var i GetAllStaffInfoRow
		if err := rows.Scan(
			&i.StaffID,
			&i.Status,
			&i.Tag,
			&i.Name,
			&i.Score,
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

const getStaffID = `-- name: GetStaffID :one
SELECT staff_id
FROM staff_auth
WHERE code = (?)
`

func (q *Queries) GetStaffID(ctx context.Context, code string) (int64, error) {
	row := q.db.QueryRowContext(ctx, getStaffID, code)
	var staff_id int64
	err := row.Scan(&staff_id)
	return staff_id, err
}

const getStaffInfo = `-- name: GetStaffInfo :one
SELECT status, role.tag, staff.name, score
FROM staff, role
WHERE staff.role_id = role.role_id
AND staff.status = TRUE
AND staff_id = (?)
`

type GetStaffInfoRow struct {
	Status bool
	Tag    string
	Name   string
	Score  int32
}

func (q *Queries) GetStaffInfo(ctx context.Context, staffID int64) (GetStaffInfoRow, error) {
	row := q.db.QueryRowContext(ctx, getStaffInfo, staffID)
	var i GetStaffInfoRow
	err := row.Scan(
		&i.Status,
		&i.Tag,
		&i.Name,
		&i.Score,
	)
	return i, err
}

const getStaffRole = `-- name: GetStaffRole :one
SELECT tag
FROM role, staff
WHERE staff.staff_id = (?)
AND role.role_id = staff.role_id
`

func (q *Queries) GetStaffRole(ctx context.Context, staffID int64) (string, error) {
	row := q.db.QueryRowContext(ctx, getStaffRole, staffID)
	var tag string
	err := row.Scan(&tag)
	return tag, err
}

const setStaffQuit = `-- name: SetStaffQuit :exec
UPDATE staff
SET status = FALSE
WHERE staff_id = (?)
`

func (q *Queries) SetStaffQuit(ctx context.Context, staffID int64) error {
	_, err := q.db.ExecContext(ctx, setStaffQuit, staffID)
	return err
}
