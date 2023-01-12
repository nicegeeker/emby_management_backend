-- name: CreateOrderType :one
INSERT INTO order_types (
    days,
    price
) VALUES (
    $1, $2
) RETURNING *;


-- name: GetOrderType :one
SELECT * FROM order_types
WHERE id = $1 LIMIT 1;


-- name: ListOrderTypes :many
SELECT * FROM order_types
ORDER BY id;