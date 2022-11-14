package main

import (
	"Design-Patterns/abstract-factory"
	"fmt"
)

const (
	ClienteType = iota
	GestorType
	ProdutorType
)

func main() {

	gestorFactory, _ := abstractfactory.GetUserFactory(GestorType)

	gestor := gestorFactory.MakeUser(abstractfactory.User{
		Nome:      "John",
		Sobrenome: "Smith",
		Password:  "123",
	})

	fmt.Println(gestor)
}
