package hw10programoptimization

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

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
	u, err := GetUsers(r)
	if err != nil {
		return nil, fmt.Errorf("get users error: %w", err)
	}
	return CountDomains(u, domain)
}

type Users [100_000]User

func GetUsers(r io.Reader) (result Users, err error) {
	content, err := io.ReadAll(r)
	if err != nil {
		return
	}

	lines := strings.Split(string(content), "\n")
	user := &User{}
	for i, line := range lines {
		err = json.Unmarshal([]byte(line), user)
		if err != nil {
			return
		}
		result[i] = *user
	}
	return
}

func CountDomains(u Users, domain string) (DomainStat, error) {
	result := make(DomainStat)
	query := "." + domain

	for _, user := range u {
		matched := strings.Contains(user.Email, query)
		if matched {
			key := strings.ToLower(strings.SplitN(user.Email, "@", 2)[1])
			result[key]++
		}
	}
	return result, nil
}
