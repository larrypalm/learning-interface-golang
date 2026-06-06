package request

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Validator interface {
	ValidateRequest() error
}

func Validate[T Validator](r *http.Request) (T, error) {
	var v T
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		return v, fmt.Errorf("invalid json: %w", err)
	}
	return v, v.ValidateRequest()
}
