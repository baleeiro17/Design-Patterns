package decorator

type Notifier interface {
	Send(msg string)
}
