package domain

import "errors"

var (
	ErrAccountNotFound = errors.New("account not found")
	ErrDuplicatedAPIKey = errors.New("duplicated api key")
	ErrInvoiceNotFound = errors.New("invoice not found")
	ErrUnauthorized = errors.New("unauthorized access")
	ErrAccountAlreadyExists = errors.New("account already exists")
	ErrInvalidAmount = errors.New("invalid amount")
	ErrInvalidStatus = errors.New("invalid status")
	ErrInvalidAccountID = errors.New("invalid account id")
	ErrInvalidPaymentType = errors.New("invalid payment type")
)
