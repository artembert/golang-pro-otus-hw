package hw09structvalidator

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func ValidateInt(val int, rule string) (validationError []error, parsingError error) {
	var validationErrors []error
	for _, r := range strings.Split(rule, "|") {
		rulePair := strings.Split(r, ":")
		if rulePair[0] == "" || rulePair[1] == "" {
			return nil, ErrUnknownRule
		}
		fmt.Println(rulePair)
		switch rulePair[0] {
		case "min":
			minInt, err := strconv.Atoi(rulePair[1])
			if err != nil {
				return nil, ErrUnknownRule
			}
			if val < minInt {
				validationErrors = append(validationErrors, ErrMinConstraint)
			}
		case "max":
			maxInt, err := strconv.Atoi(rulePair[1])
			if err != nil {
				return nil, ErrUnknownRule
			}
			if val > maxInt {
				validationErrors = append(validationErrors, ErrMaxConstraint)
			}
		case "in":
			availableValues := strings.Split(rulePair[1], ",")
			if !slices.Contains(availableValues, strconv.FormatInt(int64(val), 10)) {
				validationErrors = append(validationErrors, ErrAvailableValues)
			}
		}
	}

	return validationErrors, nil
}
