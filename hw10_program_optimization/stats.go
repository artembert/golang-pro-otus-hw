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

type Users [100_000]User

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	result := make(DomainStat)

	fileScanner := bufio.NewScanner(r)
	fileScanner.Split(bufio.ScanLines)
	query := "." + domain
	user := &User{}

	for fileScanner.Scan() {
		err := json.Unmarshal(fileScanner.Bytes(), user)
		if err != nil {
			return nil, fmt.Errorf("%w, %w", ErrUserParsing, err)
		}
		matched := strings.Contains(user.Email, query)
		if matched {
			key := strings.ToLower(strings.SplitN(user.Email, "@", 2)[1])
			result[key]++
		}
	}

	return result, nil
}
