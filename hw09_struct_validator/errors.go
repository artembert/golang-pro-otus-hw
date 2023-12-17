package hw09structvalidator

import "errors"

var (
	ErrUnknownRule        = errors.New("unknown validation rule")
	ErrEmptyValidationTag = errors.New("empty validation rule")
)

var (
	ErrMinConstraint    = errors.New("value is smaller then min")
	ErrMaxConstraint    = errors.New("value is bigger then max")
	ErrAvailableValues  = errors.New("out of available values")
	ErrLengthConstraint = errors.New("value is longer then allowed")
)
