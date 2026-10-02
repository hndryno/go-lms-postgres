-- name: CreateRepayment :exec
CALL create_repayment(
    sqlc.arg(repayment_id)::UUID,
    sqlc.arg(loan_id)::UUID,
    sqlc.arg(amount)::NUMERIC
);

-- name: GetRepayment :one
SELECT *
FROM repayments
WHERE id = $1;

-- name: ListRepayments :many
SELECT *
FROM repayments
WHERE loan_id = $1
ORDER BY payment_date DESC, created_at DESC
LIMIT $2
OFFSET $3;

-- name: CountRepayments :one
SELECT COUNT(*)
FROM repayments
WHERE loan_id = $1;