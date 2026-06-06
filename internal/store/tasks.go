package store

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type CreateTaskForm struct {
	Name       string
	ExternalId uuid.UUID
}

func (s *Store) CreateTask(ctx context.Context, payload CreateTaskForm) error {
	_, err := s.Pool.Exec(ctx, `
		INSERT INTO tasks (external_id, name)
		VALUES ($1, $2)
	`, payload.ExternalId, payload.Name)

	if err != nil {
		return fmt.Errorf("error: ", err)
	}

	return nil
}
