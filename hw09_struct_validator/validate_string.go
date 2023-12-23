package hw09structvalidator

import (
	"slices"
	"strconv"
	"strings"
)

func ValidateString(val string, rule string) (validationErr []error, parsingErr error) {
	var validationErrors []error
	for _, r := range strings.Split(rule, "|") {
		rulePair := strings.Split(r, ":")
		if rulePair[0] == "" || rulePair[1] == "" {
			return nil, ErrParsingRule{Rule: rule}
		}

		switch rulePair[0] {
		case "len":
			maxlength, err := strconv.Atoi(rulePair[1])
			if err != nil {
				return nil, ErrParsingRule{Rule: rule}
			}
			if len(val) != maxlength {
				validationErrors = append(validationErrors, ErrLengthConstraint{Constraint: maxlength, GivenValue: val})
			}
		case "in":
			availableValues := strings.Split(rulePair[1], ",")
			if !slices.Contains(availableValues, val) {
				validationErrors = append(validationErrors, ErrAvailableValues{Constraint: availableValues, GivenValue: val})
			}
			//case "regexp":
			//	maxInt, err := strconv.Atoi(rulePair[1])
			//	if err != nil {
			//		return nil, ErrParsingRule{Rule: rule}
			//	}
			//	if val > maxInt {
			//		validationErrors = append(validationErrors, ErrMaxConstraint{Constraint: maxInt, GivenValue: val})
			//	}
		}
	}

	return validationErrors, nil
}
