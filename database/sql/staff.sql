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
SELECT status, role.tag, staff.name, score
FROM staff, role
WHERE staff.role_id = role.role_id
AND staff.status = TRUE
AND staff_id = (?);

-- name: GetAllStaffInfo :many
SELECT staff_id, status, role.tag, staff.name, score
FROM staff, role
WHERE staff.role_id = role.role_id
AND staff.status = TRUE;