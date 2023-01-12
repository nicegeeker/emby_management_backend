
-- name: CreateOrder :one
INSERT INTO orders (
    user_name,
    order_type_id,
    discount
) VALUES (
    $1, $2, $3
) RETURNING *;


-- name: GetOrder :one
SELECT * FROM orders
WHERE id = $1 LIMIT 1;


-- name: ListOrders :many
SELECT * FROM orders
WHERE user_name = $1
ORDER BY id
LIMIT $2
OFFSET $3;