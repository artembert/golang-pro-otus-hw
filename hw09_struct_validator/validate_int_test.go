package hw09structvalidator

import (
	"testing"
)

func TestValidateInt(t *testing.T) {
	t.Run("should split multiple rules", func(t *testing.T) {
		rulesSet := []string{"min:18|max:50", "in:admin,stuff", "len:11", "min:0|max:2040"}
		for _, rule := range rulesSet {
			_, parsingErrors := ValidateInt(1, rule)

			if parsingErrors != nil {
				t.Errorf("unexpected parsing error, %v", parsingErrors)
			}
		}

	})
}
