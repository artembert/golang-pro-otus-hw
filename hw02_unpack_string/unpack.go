package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode/utf8"
)

var ErrInvalidString = errors.New("invalid string")

var slashRune, _ = utf8.DecodeRuneInString("\\")

func isSlash(letter rune) bool {
	slashRune, _ := utf8.DecodeRuneInString("\\")
	return slashRune == letter
}

func Unpack(line string) (string, error) {
	runes := []rune(line)
	var tokens []string
	//var unzipLetters []rune

	for index := 0; index < len(runes); index++ {
		if runes[index] == slashRune {
			var stringBuilder strings.Builder
			stringBuilder.WriteRune(slashRune)
			stringBuilder.WriteRune(runes[index+1])
			tokens = append(tokens, stringBuilder.String())
			index++
			continue
		}
		tokens = append(tokens, string(runes[index]))
	}

	var b strings.Builder
	for _, letter := range tokens {
		b.WriteString(letter)
		b.WriteString(", ")
	}

	print(b.String())
	return b.String(), nil

	//for index := 0; index < len(runes); index++ {
	//	if isSlash(runes[index]) {
	//		unzipLetters = append(unzipLetters, runes[index+1])
	//		index++
	//		continue
	//	}
	//	zipNumber, err := strconv.Atoi(string(runes[index]))
	//	if err != nil {
	//		unzipLetters = append(unzipLetters, runes[index])
	//		continue
	//	}
	//	if index == 0 {
	//		return "", ErrInvalidString
	//	}
	//	if zipNumber == 0 {
	//		unzipLetters = unzipLetters[:len(unzipLetters)-1]
	//	} else {
	//		for i := 0; i < zipNumber-1; i++ {
	//			unzipLetters = append(unzipLetters, runes[index-1])
	//		}
	//	}
	//}
	//
	//var b strings.Builder
	//for _, letter := range unzipLetters {
	//	b.WriteRune(letter)
	//}
	//
	//return b.String(), nil
}
