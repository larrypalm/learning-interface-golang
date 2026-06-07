package notify

import "fmt"

type EmailNotifier struct {
	To string
}

func NewEmailNotifier(to string) EmailNotifier {
	return EmailNotifier{To: to}
}

func (e EmailNotifier) Send(message string) error {
	fmt.Printf("Sending email to %s: %s\n", e.To, message)
	return nil
}
