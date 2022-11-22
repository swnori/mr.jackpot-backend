-- name: GetNonmemberToken: one
SELECT token
FROM nonmember;

-- name: GetMemberHashedPW :one
SELECT password
FROM member;

-- name: CreateUser :execresult
INSERT INTO user
VALUES ();

-- name: CreateMember :exec
INSERT INTO member (id, userid, password)
VALUES (?, ?, ?);

-- name: CreateNonmember :exec
INSERT INTO nonmember (id, token)
VALUES (?, ?);

-- name: CreatePersonalInfo :exec
INSERT INTO member_info (userid, name, address, phone)
VALUES (?, ?, ?, ?);

-- name: GetUserInfo :one
SELECT address, phone
FROM member_info
WHERE user_id = (?);