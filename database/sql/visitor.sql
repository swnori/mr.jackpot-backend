-- name: CreateUser :execresult
INSERT INTO user
VALUES ();

-- name: CreateVisitor :exec
INSERT INTO visitor (visitor_id, identifier)
VALUES (?, ?);

-- name: GetVisitorID :one
SELECT visitor_id
FROM visitor
WHERE identifier = (?);

