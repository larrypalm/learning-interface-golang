package validate

type Validator interface {
	Validate(input string) error
}

func ValidateAll[T Validator](validators []T, input string) []error {
	var errors []error
	for _, validate := range validators {
		if err := validate.Validate(input); err != nil {
			errors = append(errors, err)
		}
	}
	return errors
}
