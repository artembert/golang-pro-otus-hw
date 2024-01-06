package hw10programoptimization

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/mailru/easyjson"
	"io"
	"strings"
)

var ErrUserParsing = errors.New("user parsing failed")

//easyjson:json -all
type User struct {
	ID       int
	Name     string
	Username string
	Email    string
	Phone    string
	Password string
	Address  string
}

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	result := make(DomainStat)

	fileScanner := bufio.NewScanner(r)
	fileScanner.Split(bufio.ScanLines)
	query := "." + domain
	user := &User{}

	for fileScanner.Scan() {
		err := easyjson.Unmarshal(fileScanner.Bytes(), user)
		if errors.Is(io.EOF, err) {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("%w, %w", ErrUserParsing, err)
		}
		matched := strings.HasSuffix(user.Email, query)
		if matched {
			key, err := extractDomain(user.Email)
			if err == nil {
				result[key]++
			}
		}
	}

	return result, nil
}

func extractDomain(email string) (string, error) {
	atIndex := strings.IndexRune(email, '@')
	if atIndex < 0 {
		return "", fmt.Errorf("unsupported email format: '@' was not found in %s", email)
	}
	return strings.ToLower(email[atIndex+1:]), nil
}
