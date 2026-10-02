-- name: CreateLoan :exec
CALL create_loan(
    sqlc.arg(loan_id)::UUID,
    sqlc.arg(customer_id)::UUID,
    sqlc.arg(principal_amount)::NUMERIC,
    sqlc.arg(interest_rate)::NUMERIC,
    sqlc.arg(tenor)::INT
);

-- name: GetLoan :one
SELECT *
FROM loans
WHERE id = $1;

-- name: ListLoans :many
SELECT *
FROM loans
ORDER BY created_at DESC
LIMIT $1
OFFSET $2;

-- name: CountLoans :one
SELECT COUNT(*)
FROM loans;