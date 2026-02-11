package validator

import (
	"errors"
	"regexp"
	"strings"
)

var (
	ErrRequired      = errors.New("field is required")
	ErrInvalidEmail  = errors.New("invalid email format")
	ErrInvalidLength = errors.New("invalid length")
	ErrInvalidFormat = errors.New("invalid format")
	ErrMinLength     = errors.New("minimum length requirement not met")
)

type Validator struct {
	errors map[string]string
}

func NewValidator() *Validator {
	return &Validator{
		errors: make(map[string]string),
	}
}

func (v *Validator) Required(field, value string) *Validator {
	if strings.TrimSpace(value) == "" {
		v.errors[field] = "field is required"
	}
	return v
}

func (v *Validator) Email(field, value string) *Validator {
	if !isValidEmail(value) {
		v.errors[field] = "invalid email format"
	}
	return v
}

func (v *Validator) MinLength(field string, value string, min int) *Validator {
	if len(value) < min {
		v.errors[field] = "minimum length is " + string(rune(min))
	}
	return v
}

func (v *Validator) MaxLength(field string, value string, max int) *Validator {
	if len(value) > max {
		v.errors[field] = "maximum length is " + string(rune(max))
	}
	return v
}

func (v *Validator) HasErrors() bool {
	return len(v.errors) > 0
}

func (v *Validator) GetErrors() map[string]string {
	return v.errors
}

func (v *Validator) GetFirstError() string {
	for _, err := range v.errors {
		return err
	}
	return ""
}

func isValidEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(email)
}
