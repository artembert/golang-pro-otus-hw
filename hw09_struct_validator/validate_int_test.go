package hw09structvalidator

import (
	"errors"
	"testing"
)

func TestValidateInt(t *testing.T) {
	t.Run("should split multiple rules", func(t *testing.T) {
		rulesSet := []string{"min:18|max:50", "in:13,17", "min:18|max:50|in:13,17"}
		for _, rule := range rulesSet {
			_, parsingErrors := ValidateInt(1, rule)

			if parsingErrors != nil {
				t.Errorf("unexpected parsing error, %v", parsingErrors)
			}
		}

	})

	t.Run("should support single condition", func(t *testing.T) {
		t.Run("valid", func(t *testing.T) {
			validationErrors, parsingErrors := ValidateInt(13, "min:10")
			if validationErrors != nil {
				t.Errorf("unexpected validation error, %v", parsingErrors)
			}
			if parsingErrors != nil {
				t.Errorf("unexpected parsing error, %v", parsingErrors)
			}
		})

		t.Run("minConstraint", func(t *testing.T) {
			validationErrors, parsingErrors := ValidateInt(9, "min:10")
			if len(validationErrors) != 1 {
				t.Errorf("unexpected number of validation errors, %v", parsingErrors)
			}
			if !errors.Is(validationErrors[0], ErrMinConstraint) {
				t.Errorf("unexpected validation error, %v", parsingErrors)
			}
			if parsingErrors != nil {
				t.Errorf("unexpected parsing error, %v", parsingErrors)
			}
		})

		t.Run("maxConstraint", func(t *testing.T) {
			validationErrors, parsingErrors := ValidateInt(31, "max:30")
			if len(validationErrors) != 1 {
				t.Errorf("unexpected number of validation errors, %v", parsingErrors)
			}
			if !errors.Is(validationErrors[0], ErrMaxConstraint) {
				t.Errorf("unexpected validation error, %v", parsingErrors)
			}
			if parsingErrors != nil {
				t.Errorf("unexpected parsing error, %v", parsingErrors)
			}
		})

		t.Run("availableValues", func(t *testing.T) {
			validationErrors, parsingErrors := ValidateInt(14, "in:13,17")
			if len(validationErrors) != 1 {
				t.Errorf("unexpected number of validation errors, %v", parsingErrors)
			}
			if !errors.Is(validationErrors[0], ErrAvailableValues) {
				t.Errorf("unexpected validation error, %v", parsingErrors)
			}
			if parsingErrors != nil {
				t.Errorf("unexpected parsing error, %v", parsingErrors)
			}
		})
	})
}
