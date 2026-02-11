package valueobject

import (
	"regexp"
	"strings"
)

type Email struct {
	value string
}

func NewEmail(email string) (*Email, error) {
	email = strings.TrimSpace(strings.ToLower(email))

	if !isValidEmail(email) {
		return nil, ErrInvalidEmail
	}

	return &Email{value: email}, nil
}

func (e *Email) Value() string {
	return e.value
}

func (e *Email) Equals(other *Email) bool {
	if other == nil {
		return false
	}
	return e.value == other.value
}

func isValidEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(email)
}
