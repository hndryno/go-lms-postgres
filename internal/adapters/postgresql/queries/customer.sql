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

-- name: UpdateCustomer :one
UPDATE customers
SET
    name = $2,
    email = $3,
    phone = $4,
    address = $5,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteCustomer :exec
DELETE FROM customers
WHERE id = $1;