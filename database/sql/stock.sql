-- name: GetAllStockList :many
SELECT stock_id, name, count, unit
FROM stock;

-- name: UpdateStockItem :exec
UPDATE stock
SET count = (?)
WHERE stock_id = (?);

-- name: AddStockItem :execresult
INSERT INTO stock (name, unit)
VALUES (?, ?);

-- name: DeleteStockItem :exec
DELETE FROM stock
WHERE stock_id = (?);

