package hw09structvalidator

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type UserRole string

// Test the function on different structures and other types.
type (
	User struct {
		ID        string `json:"id" validate:"len:36"`
		Name      string
		Age       int             `validate:"min:18|max:50"`
		Email     string          `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		Role      UserRole        `validate:"in:admin,stuff"`
		Phones    []string        `validate:"len:11"`
		BirthDate []int           `validate:"min:0|max:2040"`
		meta      json.RawMessage //nolint:unused
	}

	App struct {
		Version     string `validate:"len:5"`
		ReleaseDate string `validate:"len:10"`
		Env         string `validate:"in:dev,prod"`
		Pod         string `validate:"in:us1,us2,ca1"`
	}

	Config struct {
		PreferredLanguage string `validate:""`
	}

	Contacts struct {
		Email        string `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		WorkingEmail string `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
	}

	DemographicsAndContacts struct {
		Phones        []string `validate:"len:11"`
		WorkingPhones []string `validate:"len:11"`
		BirthDate     []int    `validate:"min:0|max:2040"`
	}

	Token struct {
		Header    []byte
		Payload   []byte
		Signature []byte
	}

	Response struct {
		Code int    `validate:"in:200,404,500"`
		Body string `json:"omitempty"`
	}

	SizesCorrupted struct {
		Height int `validate:"min:"`
		Width  int `validate:"min:7"`
		Length int `validate:"max:23"`
	}

	Sizes struct {
		Height int `validate:"min:5"`
		Width  int `validate:"min:7"`
		Length int `validate:"max:23"`
	}
)

func TestValidate(t *testing.T) {
	tests := []struct {
		name        string
		examine     interface{}
		expectedErr string
	}{
		{
			name: "Parsing rule error",
			examine: SizesCorrupted{
				Height: 21,
				Width:  6,
				Length: 24,
			},
			expectedErr: ErrParsingRule{
				Rule: "min:",
			}.Error(),
		},
		{
			name: "Int validation",
			examine: Sizes{
				Height: 21,
				Width:  6,
				Length: 24,
			},
			expectedErr: ValidationErrors{
				ValidationError{Field: "Width", Err: ErrMinConstraint{Constraint: 7, GivenValue: 6}},
				ValidationError{Field: "Length", Err: ErrMaxConstraint{Constraint: 23, GivenValue: 24}},
			}.Error(),
		},
		{
			name: "String validation",
			examine: App{
				Version:     "1.1",
				ReleaseDate: "2023-07-13",
				Env:         "production",
				Pod:         "ca1",
			},
			expectedErr: ValidationErrors{
				ValidationError{Field: "Version", Err: ErrLengthConstraint{Constraint: 5, GivenValue: "1.1"}},
				ValidationError{
					Field: "Env", Err: ErrAvailableValues{
						Constraint: []string{"dev", "prod"}, GivenValue: "production",
					},
				},
			}.Error(),
		},
		{
			name: "RegExp validation",
			examine: Contacts{
				Email:        "testadmin@corp.com",
				WorkingEmail: "test-admin@corp",
			},
			expectedErr: ValidationErrors{
				ValidationError{
					Field: "WorkingEmail", Err: ErrRegexp{
						Constraint: "regexp:^\\w+@\\w+\\.\\w+$", GivenValue: "test-admin@corp",
					},
				},
			}.Error(),
		},
		{
			name: "Slice validation",
			examine: DemographicsAndContacts{
				Phones:        []string{"54-555-123", "554555-123"},
				WorkingPhones: []string{"554-555-123", "554-555-122", "554-555-123"},
				BirthDate:     []int{-30, 10, 20, 40, 2050},
			},
			expectedErr: ValidationErrors{
				ValidationError{
					Field: "Phones[54-555-123]", Err: ErrLengthConstraint{
						Constraint: 11, GivenValue: "54-555-123",
					},
				},
				ValidationError{
					Field: "Phones[554555-123]", Err: ErrLengthConstraint{
						Constraint: 11, GivenValue: "554555-123",
					},
				},
				ValidationError{Field: "BirthDate[-30]", Err: ErrMinConstraint{Constraint: 0, GivenValue: -30}},
				ValidationError{Field: "BirthDate[2050]", Err: ErrMaxConstraint{Constraint: 2040, GivenValue: 2050}},
			}.Error(),
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf(tt.name, i), func(t *testing.T) {
			tt := tt
			t.Parallel()

			err := Validate(tt.examine)

			require.Equal(t, tt.expectedErr, err.Error())
			_ = tt
		})
	}
}
