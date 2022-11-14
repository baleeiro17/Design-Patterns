package main

import (
	"Design-Patterns/decorator"
)

func main() {

	notifierEmail := &decorator.NotifierEmail{}

	notifierFace := &decorator.NotifierFacebook{
		notifierEmail,
	}

	notifierSlack := &decorator.NotifierSlack{
		notifierFace,
	}

	notifierSlack.Send("Olá")

}
