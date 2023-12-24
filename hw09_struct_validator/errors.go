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

type ValidationError struct {
	Field string
	Err   error
}

func (validationError ValidationError) Error() error {
	return fmt.Errorf("[%s]: %w", validationError.Field, validationError.Err)
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

type ErrLengthConstraint struct {
	Constraint int
	GivenValue string
}

func (e ErrLengthConstraint) Error() string {
	return fmt.Sprintf(
		"Value '%v' of length '%v' not equal to: %v which is required",
		e.GivenValue,
		len(e.GivenValue),
		e.Constraint,
	)
}

type ErrRegexp struct {
	Constraint string
	GivenValue string
}

func (e ErrRegexp) Error() string {
	return fmt.Sprintf("Value '%v' is not satisfied regexp rules '%s'", e.GivenValue, e.Constraint)
}

func joinErrors(errs ...error) error {
	s := make([]string, 0, len(errs))
	nonNilErrs := make([]any, 0, len(errs))
	for _, err := range errs {
		if err == nil {
			continue
		}
		s = append(s, "%w")
		nonNilErrs = append(nonNilErrs, err)
	}
	// If all the errors were nil, return nil.
	if len(nonNilErrs) == 0 {
		return nil
	}
	allErrs := strings.Join(s, "\n")
	return fmt.Errorf(allErrs, nonNilErrs...)
}
