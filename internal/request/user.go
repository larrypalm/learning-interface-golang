package request

import "fmt"

type CreateUserRequest struct {
	Email string `json:"email"`
}

func (r CreateUserRequest) ValidateRequest() error {
	if r.Email == "" {
		return fmt.Errorf("email is required")
	}
	return nil
}
