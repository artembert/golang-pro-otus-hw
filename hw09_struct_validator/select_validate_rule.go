package hw09structvalidator

import (
	"reflect"
)

func SelectValidateRule(field reflect.StructField) (rule string, err error) {
	TagName := "validate"
	rule, ok := field.Tag.Lookup(TagName)

	if ok {
		if rule == "" {
			return "", ErrEmptyValidationTag
		} else {
			return rule, nil
		}
	} else {
		return "", nil
	}
}
