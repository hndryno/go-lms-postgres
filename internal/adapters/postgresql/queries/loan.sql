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