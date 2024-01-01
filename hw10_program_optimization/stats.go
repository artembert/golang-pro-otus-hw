package hw10programoptimization

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"
)

var ErrUserParsing = errors.New("user parsing failed")

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
	reader := bufio.NewReader(r)
	decoder := json.NewDecoder(reader)
	fileScanner.Split(bufio.ScanLines)
	query := "." + domain
	user := &User{}

	for {
		err := decoder.Decode(user)
		if err == io.EOF {
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
