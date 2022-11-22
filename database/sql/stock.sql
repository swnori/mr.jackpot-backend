-- name: GetAllInventoryList :many
SELECT stock_id, name, count
FROM stock;

-- name: UpdateInventoryItem :exec
UPDATE stock
SET count = (?)
WHERE stock_id = (?);

-- name: AddInventoryItem :execresult
INSERT INTO stock (name)
VALUES (?);

-- name: DeleteInventoryItem :exec
DELETE FROM stock
WHERE stock_id = (?);
