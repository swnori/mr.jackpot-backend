
-- name: GetOrderHistory :many
SELECT price, discount, created_at, reserve_at, order_id
FROM `order`
WHERE user_id = (?);

-- name: GetDinnerListHistory :many
SELECT name
FROM ordered_dinner, dinner, board_entity
WHERE order_id = (?)
AND ordered_dinner.dinner_id = dinner.dinner_id
AND dinner.entity_id = board_entity.entity_id;
