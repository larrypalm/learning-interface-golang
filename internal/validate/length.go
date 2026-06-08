package validate

import "fmt"

type LengthValidator struct {
	Min int
}

func NewLengthValidator(min int) LengthValidator {
	return LengthValidator{Min: min}
}

func (l LengthValidator) Validate(input string) error {
	fmt.Println(input, len(input), l.Min)
	if len(input) < l.Min {
		return fmt.Errorf("input: \"%s\", to short. Minimum %d characters", input, l.Min)
	}
	return nil
}
