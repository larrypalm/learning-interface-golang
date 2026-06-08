package validate

import (
	"fmt"
	"strings"
)

type EmailValidator struct{}

func NewEmailValidator(email string) EmailValidator {
	return EmailValidator{}
}

func (e EmailValidator) Validate(input string) error {
	if !strings.Contains(input, "@") {
		return fmt.Errorf("%s is not a valid email", input)
	}
	return nil
}
