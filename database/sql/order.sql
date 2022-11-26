-- name: UpdateOrderState :exec
UPDATE order_state
SET state_id = (
SELECT state_id
FROM state
WHERE name = (?))
WHERE order_id = (?);

-- name: GetOrderInfo :exec
SELECT order.order_id, order.price, order.discount, order.reserve_at
FROM `order`, order_state
WHERE order.user_id = (?)
AND order.order_id = order_state.order_id
AND order_state.state_id = (
    SELECT state_id 
    FROM state
    WHERE name = "Finished"
);

-- name: GetOrderHistory :many
SELECT order_id, price, discount, reserve_at
FROM `order`
WHERE user_id = (?);