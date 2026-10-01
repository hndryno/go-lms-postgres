-- name: CreateCustomer :one
INSERT INTO customers (
    name,
    email,
    phone,
    address
)
VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING *;

-- name: ListCustomers :many
SELECT *
FROM customers
ORDER BY created_at DESC
LIMIT $1
OFFSET $2;

-- name: CountCustomers :one
SELECT COUNT(*)
FROM customers;

-- name: GetCustomer :one
SELECT *
FROM customers
WHERE id = $1;