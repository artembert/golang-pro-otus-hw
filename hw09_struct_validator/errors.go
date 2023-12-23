package hw09structvalidator

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrUnknownRule        = errors.New("unknown validation rule")
	ErrEmptyValidationTag = errors.New("empty validation rule")
	ErrNotAStruct         = errors.New("not a type of structure")
)

var (
	ErrMinConstraint    = errors.New("value is smaller then min")
	ErrMaxConstraint    = errors.New("value is bigger then max")
	ErrAvailableValues  = errors.New("out of available values")
	ErrLengthConstraint = errors.New("value is longer then allowed")
)

type ValidationError struct {
	Field string
	Err   error
}

type ValidationErrors []ValidationError

func (validationError ValidationError) Error() error {
	return fmt.Errorf("[%s]: %s", validationError.Field, validationError.Err)
}

func (v ValidationErrors) Error() string {
	if len(v) < 1 {
		return "No validation errors fount"
	}
	builder := strings.Builder{}
	builder.WriteString("Validation errors:")

	for _, err := range v {
		builder.WriteString(fmt.Sprintf("\n%s;", err.Error()))
	}

	return builder.String()
}
