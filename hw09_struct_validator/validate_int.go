package hw09structvalidator

import (
	"slices"
	"strconv"
	"strings"
)

func ValidateInt(val int, rule string) (validationErr []error, parsingErr error) {
	var validationErrors []error
	for _, r := range strings.Split(rule, "|") {
		rulePair := strings.Split(r, ":")
		if rulePair[0] == "" || rulePair[1] == "" {
			return nil, ErrParsingRule{Rule: rule}
		}

		switch rulePair[0] {
		case "min":
			minInt, err := strconv.Atoi(rulePair[1])
			if err != nil {
				return nil, ErrParsingRule{Rule: rule}
			}
			if val < minInt {
				validationErrors = append(validationErrors, ErrMinConstraint{Constraint: minInt, GivenValue: val})
			}
		case "max":
			maxInt, err := strconv.Atoi(rulePair[1])
			if err != nil {
				return nil, ErrParsingRule{Rule: rule}
			}
			if val > maxInt {
				validationErrors = append(validationErrors, ErrMaxConstraint{Constraint: maxInt, GivenValue: val})
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
