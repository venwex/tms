package auth

import "errors"

var (
	ErrRequiredEmail     = errors.New("email is required")
	ErrRequiredPassword  = errors.New("password is required")
	ErrRequiredFirstName = errors.New("first name is required")
	ErrRequiredLastName  = errors.New("last name is required")

	ErrInvalidPassword = errors.New("password is invalid")
)
