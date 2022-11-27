-- name: CreateCustomerAuth :exec
INSERT INTO customer_auth (id, password, customer_id)
VALUES (?, ?, ?);

-- name: CreateCustomer :exec
INSERT INTO customer (customer_id, name, address, phone)
VALUES (?, ?, ?, ?);

-- name: GetCustomerPassword :one
SELECT password, customer_id
FROM customer_auth
WHERE id = (?);

-- name: CheckCustomerID :one
SELECT customer_id
FROM customer_auth
WHERE id = (?);

-- name: GetPersonalInfo :one
SELECT customer_id, name, address, phone
FROM customer
WHERE customer_id = (?);

-- name: UpdatePersonalInfo :exec
UPDATE customer
SET name = (?),
    phone = (?),
    address = (?)
WHERE customer_id = (?);

-- name: GetAllCustomerInfo :many
SELECT customer_id, name, address, phone, orders, rating, paid, created_at
FROM customer
WHERE status = TRUE;

-- name: SetCustomerQuit :exec
UPDATE customer
SET status = FALSE
WHERE customer_id = (?);

