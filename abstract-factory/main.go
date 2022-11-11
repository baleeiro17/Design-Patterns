package main

import "fmt"

const (
	ClienteType = iota
	GestorType
	ProdutorType
)

func main() {

	gestorFactory, _ := GetUserFactory(GestorType)

	gestor := gestorFactory.makeUser(User{
		nome:      "John",
		sobrenome: "Smith",
		password:  "123",
	})

	fmt.Println(gestor)
}
