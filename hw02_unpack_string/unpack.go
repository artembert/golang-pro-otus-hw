package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode/utf8"
)

var ErrInvalidString = errors.New("invalid string")

func isSlash(letter rune) bool {
	slashRune, _ := utf8.DecodeRuneInString("\\")
	return slashRune == letter
}

func Unpack(line string) (string, error) {
	runes := []rune(line)
	var unzipLetters []rune

	for index := 0; index < len(runes); index++ {
		if isSlash(runes[index]) {
			unzipLetters = append(unzipLetters, runes[index+1])
			index++
			continue
		}
		if zipNumber, err := strconv.Atoi(string(runes[index])); err == nil {
			if zipNumber == 0 {
				if len(unzipLetters) > 0 {
					unzipLetters = unzipLetters[:len(unzipLetters)-1]
				} else {
					return "", errors.New("invalid format of string")
				}
			} else {
				for i := 0; i < zipNumber-1; i++ {
					unzipLetters = append(unzipLetters, runes[index-1])
				}
			}
		} else {
			unzipLetters = append(unzipLetters, runes[index])
		}
	}

	var b strings.Builder
	for _, letter := range unzipLetters {
		b.WriteRune(letter)
	}

	return b.String(), nil
}
