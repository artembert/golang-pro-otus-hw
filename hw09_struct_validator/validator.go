package hw09structvalidator

import (
	"fmt"
	"reflect"
	"strconv"
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
			val := valueOfStruct.Field(i).String()
			validationResult, parsingError := ValidateString(val, rule)
			if parsingError != nil {
				return parsingError
			}
			fieldErrors = validationResult

			if fieldErrors != nil {
				validationErrors = append(validationErrors, ValidationError{
					Field: field.Name,
					Err:   joinErrors(fieldErrors...), // errors.Join(fieldErrors...) in Go v1.20
				})
			}
		case reflect.Int:
			val := int(valueOfStruct.Field(i).Int())
			validationResult, parsingError := ValidateInt(val, rule)
			if parsingError != nil {
				return parsingError
			}
			fieldErrors = validationResult

			if fieldErrors != nil {
				validationErrors = append(validationErrors, ValidationError{
					Field: field.Name,
					Err:   joinErrors(fieldErrors...), // errors.Join(fieldErrors...) in Go v1.20
				})
			}
		case reflect.Slice:
			sliceKind := field.Type.Elem().Kind()
			slice := valueOfStruct.Field(i)
			sliceValidationErrors, parsingError := validateSlice(slice, sliceKind, rule, field.Name)
			if parsingError != nil {
				return parsingError
			}
			if sliceValidationErrors != nil {
				validationErrors = append(validationErrors, sliceValidationErrors...)
			}
		default:
		}
	}

	if len(validationErrors) > 0 {
		return validationErrors
	}

	return nil
}

func validateSlice(
	slice reflect.Value,
	sliceKind reflect.Kind,
	rule string,
	fieldName string,
) (validationErrors []ValidationError, parsingError error) {
	switch sliceKind {
	case reflect.String:
		for sliceElIndex := 0; sliceElIndex < slice.Len(); sliceElIndex++ {
			val := slice.Index(sliceElIndex).String()
			validationResult, parsingError := ValidateString(val, rule)
			if parsingError != nil {
				return nil, parsingError
			}
			if validationResult != nil {
				validationErrors = append(validationErrors, ValidationError{
					Field: fmt.Sprintf("%s[%s]", fieldName, val),
					Err:   joinErrors(validationResult...), // errors.Join(fieldErrors...) in Go v1.20
				})
			}
		}

		return validationErrors, nil
	case reflect.Int:
		for sliceElIndex := 0; sliceElIndex < slice.Len(); sliceElIndex++ {
			val := int(slice.Index(sliceElIndex).Int())
			validationResult, parsingError := ValidateInt(val, rule)
			if parsingError != nil {
				return nil, parsingError
			}
			if validationResult != nil {
				validationErrors = append(validationErrors, ValidationError{
					Field: fmt.Sprintf("%s[%s]", fieldName, strconv.FormatInt(int64(val), 10)),
					Err:   joinErrors(validationResult...), // errors.Join(fieldErrors...) in Go v1.20
				})
			}
		}

		return validationErrors, nil
	default:
		return nil, nil
	}
}
