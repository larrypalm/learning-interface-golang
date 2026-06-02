package main

import (
	"context"
	"fmt"
	"learn-interfaces-go/internal/handler"
	"learn-interfaces-go/internal/store"
	"log"
	"net/http"
)

func main() {
	ctx := context.Background()
	store, err := store.New(ctx)
	if err != nil {
		return
	}

	handler := handler.New(store)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	fmt.Println("Server running on port 8080")
	log.Fatal(srv.ListenAndServe())
}
