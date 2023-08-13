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
	var b strings.Builder

	for index := 0; index < len(runes); index++ {
		fmt.Printf("index: %d\n", index)
		if zipNumber, err := strconv.Atoi(string(runes[index])); err == nil {
			fmt.Printf("zipNumber: %d\n", zipNumber)
			for i := 0; i < zipNumber-1; i++ {
				b.WriteRune(runes[index-1])
			}
		} else {
			fmt.Printf("rune: %s\n", string(runes[index]))
			b.WriteRune(runes[index])
		}
		fmt.Printf("collection: [%s]\n", b.String())
	}
	return b.String(), nil
}
