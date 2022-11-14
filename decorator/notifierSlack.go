package decorator

import "fmt"

type NotifierSlack struct {
	Notifier Notifier
}

func (n *NotifierSlack) Send(msg string) {
	n.Notifier.Send(msg)
	fmt.Printf("Send the message by Slack: %s\n", msg)
}
