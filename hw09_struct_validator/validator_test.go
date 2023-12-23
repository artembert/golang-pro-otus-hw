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
		Version string `validate:"len:5"`
	}

	Config struct {
		PreferredLanguage string `validate:""`
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
