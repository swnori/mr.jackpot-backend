-- name: UpdateOrderState :exec
UPDATE order_state
SET state_id = (
SELECT state_id
FROM state
WHERE name = (?))
WHERE order_id = (?);