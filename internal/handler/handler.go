package handler

import (
	"fmt"
	"learn-interfaces-go/internal/store"
	"net/http"

	goauth "github.com/larrypalm/go-auth"
)

type Handler struct {
	Mux   *http.ServeMux
	Store *store.Store
}

func New(store *store.Store) *http.ServeMux {
	h := &Handler{
		Mux:   http.NewServeMux(),
		Store: store,
	}

	authConfig := goauth.Config{
		UserStore:  store,
		TokenStore: store,
	}
	goauth := goauth.New(authConfig)

	h.Mux.Handle("/auth/", goauth.Routes())
	h.Mux.HandleFunc("GET /test", h.SaySomething)

	return h.Mux
}

func (handler *Handler) SaySomething(w http.ResponseWriter, r *http.Request) {
	fmt.Print("HESJA")
}
