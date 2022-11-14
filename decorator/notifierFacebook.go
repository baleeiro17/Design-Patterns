package decorator

import "fmt"

type NotifierFacebook struct {
	Notifier Notifier
}

func (n *NotifierFacebook) Send(msg string) {
	n.Notifier.Send(msg)
	fmt.Printf("Send the message by Facebook: %s\n", msg)
}
