package decorator

import "fmt"

type NotifierEmail struct {
}

func (n *NotifierEmail) Send(msg string) {
	fmt.Printf("Send the message by Email: %s\n", msg)
}
