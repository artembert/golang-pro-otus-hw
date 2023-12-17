package hw09structvalidator

import (
	"errors"
	"testing"
)

func TestValidateInt(t *testing.T) {
	t.Run("should support single constraint", func(t *testing.T) {
		cases := []struct {
			name           string
			value          int
			validationRule string
			expectedErrors []error
		}{
			{
				name:           "✅Valid: min",
				value:          13,
				validationRule: "min:10",
				expectedErrors: []error{},
			},
			{
				name:           "⚠️Invalid: min",
				value:          9,
				validationRule: "min:17",
				expectedErrors: []error{ErrMinConstraint},
			},
			{
				name:           "✅Valid: max",
				value:          77,
				validationRule: "max:77",
				expectedErrors: []error{},
			},
			{
				name:           "⚠️Invalid: max",
				value:          14,
				validationRule: "max:13",
				expectedErrors: []error{ErrMaxConstraint},
			},
			{
				name:           "✅Valid: in",
				value:          13,
				validationRule: "in:13,17",
				expectedErrors: []error{},
			},
			{
				name:           "⚠️Invalid: in",
				value:          0,
				validationRule: "in:18,12",
				expectedErrors: []error{ErrAvailableValues},
			},
		}

		for _, testCase := range cases {
			t.Run(testCase.name, func(t *testing.T) {
				validationErrors, parsingErrors := ValidateInt(testCase.value, testCase.validationRule)
				if parsingErrors != nil {
					t.Errorf("unexpected parsing error, %v", parsingErrors)
				}
				for i, err := range validationErrors {
					if !errors.Is(testCase.expectedErrors[i], err) {
						t.Errorf("unexpected validation error, %v", parsingErrors)
					}
				}
			})
		}
	})

	t.Run("should support multiple constraints", func(t *testing.T) {
		cases := []struct {
			name           string
			value          int
			validationRule string
			expectedErrors []error
		}{
			{
				name:           "✅Valid: min|in",
				value:          13,
				validationRule: "min:10|in:13,17",
				expectedErrors: []error{},
			},
			{
				name:           "⚠️Invalid: min|in",
				value:          9,
				validationRule: "min:10|in:13,17",
				expectedErrors: []error{ErrMinConstraint, ErrAvailableValues},
			},
		}

		for _, testCase := range cases {
			t.Run(testCase.name, func(t *testing.T) {
				validationErrors, parsingErrors := ValidateInt(testCase.value, testCase.validationRule)
				if parsingErrors != nil {
					t.Errorf("unexpected parsing error, %v", parsingErrors)
				}
				for i, err := range validationErrors {
					if !errors.Is(testCase.expectedErrors[i], err) {
						t.Errorf("unexpected validation error, %v", parsingErrors)
					}
				}
			})
		}
	})

	t.Run("should throw parsing errors", func(t *testing.T) {
		_, parsingErrors := ValidateInt(76, "min:")
		if !errors.Is(parsingErrors, ErrUnknownRule) {
			t.Errorf("parcing error expected, %v", parsingErrors)
		}
	})
}
