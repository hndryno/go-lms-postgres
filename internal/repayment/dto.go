package repayment

import "github.com/google/uuid"

type CreateRepaymentRequest struct {
	LoanID uuid.UUID `json:"loan_id"`
	Amount float64   `json:"amount"`
}

type ListRepaymentRequest struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}
