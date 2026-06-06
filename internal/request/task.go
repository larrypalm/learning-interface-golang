package request

import (
	"fmt"

	"github.com/google/uuid"
)

type CreateTaskRequest struct {
	Name string    `json:"name"`
	Uuid uuid.UUID `json:"uuid"`
}

func (r CreateTaskRequest) ValidateRequest() error {
	if r.Name == "" {
		return fmt.Errorf("name is requuuuuidred")
	}

	return nil
}
