package domain

import "errors"

var (
	ErrInvalidCapacity             = errors.New("capacity must be positive")
	ErrInvalidTimeRange            = errors.New("start time must be before end time")
	ErrUserNotFound                = errors.New("user not found")
	ErrReservationNotFound         = errors.New("reservation not found")
	ErrDuplicateEmail              = errors.New("email already exists")
	ErrInvalidCredentials          = errors.New("invalid credentials")
	ErrUnauthorized                = errors.New("unauthorized")
	ErrInvalidUser                 = errors.New("invalid user")
	ErrInvalidTimeSlot             = errors.New("invalid time slot")
	ErrReservationNotPending       = errors.New("reservation is not pending")
	ErrReservationAlreadyCancelled = errors.New("reservation is already cancelled")
	ErrCapacityExceeded            = errors.New("capacity exceeded")
)
