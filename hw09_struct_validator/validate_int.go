package hw09structvalidator

import (
	"fmt"
	"strings"
)

func ValidateInt(val int, rule string) (validationError error, parsingError error) {
	for _, r := range strings.Split(rule, "|") {
		fmt.Println(r)
	}

	return nil, nil
}
