package hw09structvalidator

import (
	"errors"
	"reflect"
)

func Validate(v interface{}) error {
	valueOfStruct := reflect.ValueOf(v)
	typeOfStruct := reflect.TypeOf(v)
	validationErrors := make(ValidationErrors, 0, valueOfStruct.NumField())

	if valueOfStruct.Kind() != reflect.Struct {
		return ErrNotAStruct
	}

	for i := 0; i < typeOfStruct.NumField(); i++ {
		field := typeOfStruct.Field(i)
		rule, err := SelectValidateRule(field)
		var fieldErrors []error

		if err != nil {
			return err
		}
		if rule == "" {
			continue
		}

		valueType := valueOfStruct.Field(i).Kind()
		switch valueType {
		case reflect.String:
		case reflect.Int:
			val := int(valueOfStruct.Field(i).Int())
			validationResult, parsingError := ValidateInt(field.Name, val, rule)
			if parsingError != nil {
				return parsingError
			}
			fieldErrors = validationResult
		case reflect.Slice:
		}

		if fieldErrors != nil {
			validationErrors = append(validationErrors, ValidationError{
				Field: field.Name,
				Err:   errors.Join(fieldErrors...),
			})
		}
	}

	if len(validationErrors) > 0 {
		return validationErrors
	}

	return nil
}
