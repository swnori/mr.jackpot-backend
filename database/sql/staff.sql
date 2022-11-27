-- name: GetStaffRole :one
SELECT tag
FROM role, staff
WHERE staff.staff_id = (?)
AND role.role_id = staff.role_id;

-- name: GetStaffID :one
SELECT staff_id
FROM staff_auth
WHERE code = (?);

-- name: CreateStaffAccount :execresult
INSERT INTO staff (role_id, name)
VALUES (?, ?);

-- name: CreateStaffAuth :exec
INSERT INTO staff_auth (staff_id, code)
VALUES (?, ?);

-- name: SetStaffQuit :exec
UPDATE staff
SET status = FALSE
WHERE staff_id = (?);

-- name: GetStaffInfo :one
SELECT status, role.tag, staff.name, score, created_at, code, staff.staff_id
FROM staff, role, staff_auth
WHERE staff.staff_id = (?)
AND staff.role_id = role.role_id
AND staff.status = TRUE
AND staff.staff_id = staff_auth.staff_id;

-- name: GetAllStaffInfo :many
SELECT staff.staff_id, status, role.tag, staff.name, score, created_at, code
FROM staff, role, staff_auth
WHERE staff.role_id = role.role_id
AND staff.status = TRUE
AND staff.staff_id = staff_auth.staff_id;