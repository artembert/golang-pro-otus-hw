package hw09structvalidator

import (
	"errors"
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
					Err:   errors.Join(fieldErrors...),
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
					Err:   errors.Join(fieldErrors...),
				})
			}
		case reflect.Slice:
			sliceKind := field.Type.Elem().Kind()
			slice := valueOfStruct.Field(i)

			switch sliceKind {
			case reflect.String:
				for sliceElIndex := 0; sliceElIndex < valueOfStruct.Field(i).Len(); sliceElIndex++ {
					val := slice.Index(sliceElIndex).String()
					validationResult, parsingError := ValidateString(val, rule)
					if parsingError != nil {
						return parsingError
					}
					if validationResult != nil {
						validationErrors = append(validationErrors, ValidationError{
							Field: fmt.Sprintf("%s[%s]", field.Name, val),
							Err:   errors.Join(validationResult...),
						})
					}
				}
			case reflect.Int:
				for sliceElIndex := 0; sliceElIndex < valueOfStruct.Field(i).Len(); sliceElIndex++ {
					val := int(slice.Index(sliceElIndex).Int())
					validationResult, parsingError := ValidateInt(val, rule)
					if parsingError != nil {
						return parsingError
					}
					if validationResult != nil {
						validationErrors = append(validationErrors, ValidationError{
							Field: fmt.Sprintf("%s[%s]", field.Name, strconv.FormatInt(int64(val), 10)),
							Err:   errors.Join(validationResult...),
						})
					}
				}
			}
		}
	}

	if len(validationErrors) > 0 {
		return validationErrors
	}

	return nil
}
