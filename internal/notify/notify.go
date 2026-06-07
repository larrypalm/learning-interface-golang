package notify

import "fmt"

type Notifier interface {
	Send(message string) error
}

func Broadcast[T Notifier](notifiers []T, message string) {
	for _, n := range notifiers {
		if err := n.Send(message); err != nil {
			fmt.Println(err)
		}
	}
}
