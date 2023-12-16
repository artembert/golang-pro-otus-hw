package hw09structvalidator

import (
	"errors"
	"reflect"
	"slices"
	"testing"
)

func TestSelectValidateRule(t *testing.T) {
	t.Run("Should return content of 'validate' struct tag", func(t *testing.T) {
		expectedRules := []string{"len:36", "min:18|max:50", "regexp:^\\w+@\\w+\\.\\w+$", "in:admin,stuff", "len:11"}
		var expectedErrors []error
		user := User{}
		st := reflect.TypeOf(user)
		rules := make([]string, 0)
		validationErrors := make([]error, 0)

		for i := 0; i < st.NumField(); i++ {
			field := st.Field(i)
			rule, err := SelectValidateRule(field)
			if err != nil {
				validationErrors = append(validationErrors, err)
			} else {
				if rule != "" {
					rules = append(rules, rule)
				}
			}
		}

		if len(expectedRules) != len(rules) {
			t.Errorf("Expected rules length mismatch given.\nExpected: %v\nGiven: %v", len(expectedRules), len(rules))
			t.Errorf("Expected rules mismatch given.\nExpected: %v\nGiven: %v", expectedRules, rules)
		}

		areRulesEqual := slices.Equal(rules, expectedRules)

		if !areRulesEqual {
			t.Errorf("Expected rules mismatch given.\nExpected: %v\nGiven: %v", expectedRules, rules)
		}

		if len(expectedErrors) != len(validationErrors) {
			t.Errorf("Expected errors length mismatch given.\nExpected: %v\nGiven: %v", len(expectedErrors), len(validationErrors))
			t.Errorf("Expected errors mismatch given.\nExpected: %v\nGiven: %v", expectedErrors, validationErrors)
		}
		for i := 0; i < len(expectedErrors); i++ {
			if !errors.Is(expectedErrors[i], validationErrors[i]) {
				t.Errorf("Expected errors mismatch given.\nExpected: %v\nGiven: %v", expectedErrors, validationErrors)
			}
		}

	})

	t.Run("Should return error on empty validation tag", func(t *testing.T) {
		expectedErrors := []error{ErrEmptyValidationTag}
		config := Config{}
		st := reflect.TypeOf(config)
		validationErrors := make([]error, 0)

		for i := 0; i < st.NumField(); i++ {
			field := st.Field(i)
			_, err := SelectValidateRule(field)
			if err != nil {
				validationErrors = append(validationErrors, err)
			}
		}

		if len(expectedErrors) != len(validationErrors) {
			t.Errorf("Expected errors length mismatch given.\nExpected: %v\nGiven: %v", len(expectedErrors), len(validationErrors))
			t.Errorf("Expected errors mismatch given.\nExpected: %v\nGiven: %v", expectedErrors, validationErrors)
		}
		for i := 0; i < len(expectedErrors); i++ {
			if !errors.Is(expectedErrors[i], validationErrors[i]) {
				t.Errorf("Expected errors mismatch given.\nExpected: %v\nGiven: %v", expectedErrors, validationErrors)
			}
		}

	})
}
