package loan

import (
	"github.com/google/uuid"
)

type CreateLoanRequest struct {
	CustomerID      uuid.UUID `json:"customer_id"`
	PrincipalAmount float64   `json:"principal_amount"`
	InterestRate    float64   `json:"interest_rate"`
	Tenor           int32     `json:"tenor"`
}

type ListLoanRequest struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}
