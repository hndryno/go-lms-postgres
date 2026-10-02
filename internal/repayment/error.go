package repayment

import "errors"

var (
	ErrLoanRequired = errors.New("loan id is required")

	ErrAmountInvalid = errors.New(
		"repayment amount must be greater than zero",
	)

	ErrLoanNotFound = errors.New(
		"loan not found",
	)

	ErrAmountExceedsRemaining = errors.New(
		"repayment amount exceeds remaining loan amount",
	)
)
