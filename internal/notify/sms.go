package notify

import "fmt"

type SmsNotifier struct {
	PhoneNr int
}

func NewSmsNotifier(phonenr int) SmsNotifier {
	return SmsNotifier{
		PhoneNr: phonenr,
	}
}

func (s SmsNotifier) Send(message string) error {
	fmt.Printf("Sending message to: %v. Message: %s\n", s.PhoneNr, message)
	return nil
}
