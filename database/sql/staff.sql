
-- name: GetStaffRole :one
SELECT tag
FROM role, staff
WHERE staff.staff_id = (?)
AND role.role_id = staff.role_id;