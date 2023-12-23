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
		builder.WriteString(fmt.Sprintf("%s\n", err.Error()))
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
	return fmt.Sprintf("Failed to parse rule '%s", e.Rule)
}

type ErrMinConstraint struct {
	Constraint int
	GivenValue int
}

func (e ErrMinConstraint) Error() string {
	return fmt.Sprintf("Given value '%v' is smaller then min: '%v'", e.GivenValue, e.Constraint)
}

type ErrMaxConstraint struct {
	Constraint int
	GivenValue int
}

func (e ErrMaxConstraint) Error() string {
	return fmt.Sprintf("Given value '%v' is bigger then max: '%v'", e.GivenValue, e.Constraint)
}

type ErrAvailableValues struct {
	Constraint []string
	GivenValue string
}

func (e ErrAvailableValues) Error() string {
	return fmt.Sprintf("Value '%v' is out of available values: [%v]", e.GivenValue, strings.Join(e.Constraint, ", "))
}
