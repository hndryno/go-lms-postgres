package customer

import "errors"

var (
	ErrNameRequired     = errors.New("name is required")
	ErrEmailRequired    = errors.New("email is required")
	ErrCustomerNotFound = errors.New("customer not found")
	ErrEmailAlreadyUsed = errors.New("email already used")
)
