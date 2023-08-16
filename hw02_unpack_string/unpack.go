package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(line string) (string, error) {
	runes := []rune(line)
	unzipLetters := []rune{}

	for index := 0; index < len(runes); index++ {
		zipNumber, err := strconv.Atoi(string(runes[index]))
		if err != nil {
			unzipLetters = append(unzipLetters, runes[index])
			continue
		}
		if index == 0 {
			return "", ErrInvalidString
		}
		if index > 0 && unicode.IsDigit(runes[index-1]) {
			return "", ErrInvalidString
		}
		if zipNumber == 0 {
			unzipLetters = unzipLetters[:len(unzipLetters)-1]
		} else {
			for i := 0; i < zipNumber-1; i++ {
				unzipLetters = append(unzipLetters, runes[index-1])
			}
		}
	}

	var b strings.Builder
	for _, letter := range unzipLetters {
		b.WriteRune(letter)
	}

	return b.String(), nil
}
