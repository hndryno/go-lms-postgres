package loan

import "errors"

var (
	ErrCustomerRequired = errors.New("customer id is required")
	ErrPrincipalInvalid = errors.New("principal amount must be greater than zero")
	ErrInterestInvalid  = errors.New("interest rate cannot be negative")
	ErrTenorInvalid     = errors.New("tenor must be greater than zero")
	ErrCustomerNotFound = errors.New("customer not found")
)
