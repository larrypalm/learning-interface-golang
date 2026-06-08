package main

import (
	"context"
	"fmt"
	"learn-interfaces-go/internal/handler"
	"learn-interfaces-go/internal/math"
	"learn-interfaces-go/internal/notify"
	"learn-interfaces-go/internal/store"
	"learn-interfaces-go/internal/validate"
	"log"
	"net/http"
)

func main() {
	rect := math.NewRectangle(1.123, 2)
	fmt.Println(math.TotalArea[math.Rectangle](rect))

	circle := math.NewCircle(1123.123)
	fmt.Println(math.TotalArea[math.Circle](circle))

	circleSizeCategory := math.SizeCategory(circle)
	fmt.Println(circleSizeCategory)

	rectPerimeter := math.TotalPerimeter(rect)
	fmt.Println(rectPerimeter)

	emailer := notify.NewEmailNotifier("hejsan")
	test := []notify.EmailNotifier{emailer}
	notify.Broadcast[notify.EmailNotifier](test, "test me")

	smser := notify.NewSmsNotifier(0707112233)
	smser2 := notify.NewSmsNotifier(0707223344)
	smsNotifiers := []notify.SmsNotifier{smser, smser2}
	notify.Broadcast[notify.SmsNotifier](smsNotifiers, "Hello")

	validators := []validate.EmailValidator{{}}
	errors := validate.ValidateAll[validate.EmailValidator](validators, "notanemail")
	for _, error := range errors {
		fmt.Println(error)
	}

	lengthValidators := []validate.LengthValidator{{Min: 5}}
	lengthErrors := validate.ValidateAll[validate.LengthValidator](lengthValidators, "test")
	for _, err := range lengthErrors {
		fmt.Println(err)
	}

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
