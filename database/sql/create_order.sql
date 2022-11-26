-- name: CreateOrderInfo :execresult
INSERT INTO `order` (user_id, price, discount, reserve_at, created_at, name, address, phone, message)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: CreateOrderState :exec
INSERT INTO order_state (order_id)
VALUES (?);

-- name: CreateOrderedDinner :execresult
INSERT INTO ordered_dinner (order_id, style_id, amount)
VALUES (?, ?, ?);

-- name: CreateOrderedMenu :exec
INSERT INTO ordered_menu (order_id, dinner_id, menutype_id, menu_id, option1_id, option2_id, count, price)
VALUES (?, ?, ?, ?, ?, ?, ?, ?);