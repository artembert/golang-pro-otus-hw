package hw09structvalidator

import (
	"fmt"
	"reflect"
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
	st := reflect.TypeOf(v)
	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)
		rule, err := SelectValidateRule(field)
		if err != nil {
			return err
		} else {
			if rule != "" {
				fmt.Println(rule)
			}
		}
	}
	return nil
}
