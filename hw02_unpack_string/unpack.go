package hw02unpackstring

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(line string) (string, error) {
	runes := []rune(line)
	unzipLetters := []rune{}

	for index := 0; index < len(runes); index++ {
		fmt.Printf("index: %d\n", index)
		if zipNumber, err := strconv.Atoi(string(runes[index])); err == nil {
			fmt.Printf("zipNumber: %d\n", zipNumber)
			if zipNumber == 0 {
				unzipLetters = unzipLetters[:len(unzipLetters)-1]
			} else {
				for i := 0; i < zipNumber-1; i++ {
					unzipLetters = append(unzipLetters, runes[index-1])
				}
			}
		} else {
			fmt.Printf("rune: %s\n", string(runes[index]))
			unzipLetters = append(unzipLetters, runes[index])
		}
		fmt.Printf("collection: [%s]\n", string(unzipLetters))
	}

	var b strings.Builder
	for _, letter := range unzipLetters {
		b.WriteRune(letter)
	}

	return b.String(), nil
}
