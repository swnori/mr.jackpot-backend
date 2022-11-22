-- name: GetAllStockList :many
SELECT
    stock_id,
    name,
    count
FROM
    stock;

-- name: UpdateStockItem :exec
UPDATE
    stock
SET
    count = (?)
WHERE
    stock_id = (?);

-- name: AddStockItem :execresult
INSERT INTO
    stock (name)
VALUES
    (?);

-- name: DeleteStockItem :exec
DELETE FROM
    stock
WHERE
    stock_id = (?);