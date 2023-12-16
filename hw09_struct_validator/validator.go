package hw09structvalidator

import (
	"errors"
	"fmt"
	"reflect"
)

var (
	ErrEmptyValidationTag = errors.New("empty validation rule")
)

type ValidationError struct {
	Field string
	Err   error
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	panic("implement me")
}

func Validate(v interface{}) error {
	TagName := "validate"
	st := reflect.TypeOf(v)
	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)
		if rule, ok := field.Tag.Lookup(TagName); ok {
			if rule == "" {
				fmt.Println(ErrEmptyValidationTag)
			} else {
				fmt.Println(rule)
			}
		}
	}
	return nil
}
