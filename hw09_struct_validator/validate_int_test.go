package hw09structvalidator

import (
	"errors"
	"github.com/stretchr/testify/require"
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
				expectedErrors: []error{ErrMinConstraint{
					Constraint: 17,
					GivenValue: 9,
				}},
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
				expectedErrors: []error{ErrMaxConstraint{Constraint: 13, GivenValue: 14}},
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
				expectedErrors: []error{ErrAvailableValues{Constraint: []string{"18", "12"}, GivenValue: "0"}},
			},
		}

		for _, testCase := range cases {
			t.Run(testCase.name, func(t *testing.T) {
				validationErrors, parsingErrors := ValidateInt(testCase.value, testCase.validationRule)
				if parsingErrors != nil {
					t.Errorf("unexpected parsing error, %v", parsingErrors)
				}
				for i, err := range validationErrors {
					require.Equal(t, testCase.expectedErrors[i], err)
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
				expectedErrors: []error{ErrMinConstraint{Constraint: 10, GivenValue: 9}, ErrAvailableValues{Constraint: []string{"13", "17"}, GivenValue: "9"}},
			},
		}

		for _, testCase := range cases {
			t.Run(testCase.name, func(t *testing.T) {
				validationErrors, parsingErrors := ValidateInt(testCase.value, testCase.validationRule)
				if parsingErrors != nil {
					t.Errorf("unexpected parsing error, %v", parsingErrors)
				}
				for i, err := range validationErrors {
					require.Equal(t, testCase.expectedErrors[i], err)
				}
			})
		}
	})

	t.Run("should throw parsing errors", func(t *testing.T) {
		_, parsingErrors := ValidateInt(76, "min:")
		if !errors.Is(parsingErrors, ErrParsingRule{
			Rule: "min:",
		}) {
			t.Errorf("parcing error expected, %v", parsingErrors)
		}
	})
}
