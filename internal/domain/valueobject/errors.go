package valueobject

import "errors"

var (
	ErrInvalidEmail = errors.New("invalid email format")
	ErrInvalidDate  = errors.New("invalid date")
)
