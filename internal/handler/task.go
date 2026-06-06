package handler

import (
	"learn-interfaces-go/internal/request"
	"learn-interfaces-go/internal/store"
	"net/http"

	"github.com/google/uuid"
)

func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	req, err := request.Validate[request.CreateTaskRequest](r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	payload := store.CreateTaskForm{
		Name:       req.Name,
		ExternalId: uuid.New(),
	}
	h.Store.CreateTask(r.Context(), payload)
	w.WriteHeader(http.StatusOK)
}
