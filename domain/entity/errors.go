package entity

import "errors"

const (
	REJECTED = "rejected"
	APPROVED = "approved"
)

var (
	ErrInvalidCreditCard       = errors.New("invalid credit card")
	ErrInvalidCreditCardNumber = errors.New("invalid credit card number")
	ErrInvalidCreditCardMonth  = errors.New("invalid credit card Monh")
	ErrInvalidCreditCardYear   = errors.New("invalid expiration year")

	ErrLimitForTransaction = errors.New("value out of transaction limit")
)
