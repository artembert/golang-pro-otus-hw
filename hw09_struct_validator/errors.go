package hw09structvalidator

import (
	"errors"
	"fmt"
	"strings"
)

var (
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

func (validationError ValidationError) Error() error {
	return fmt.Errorf("[%s]: %s", validationError.Field, validationError.Err)
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	builder := strings.Builder{}

	for _, err := range v {
		builder.WriteString(fmt.Sprintf("%s", err.Error()))
	}

	return builder.String()
}

type ErrUnknownRule struct {
	Tag string
}

func (e ErrUnknownRule) Error() string {
	return fmt.Sprintf("Unknown validator tag '%s'", e.Tag)
}

type ErrParsingRule struct {
	Rule string
}

func (e ErrParsingRule) Error() string {
	return fmt.Sprintf("failed to parse rule '%s", e.Rule)
}
