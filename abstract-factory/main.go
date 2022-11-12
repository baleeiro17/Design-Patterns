package abstractfactory

import "fmt"

const (
	ClienteType = iota
	GestorType
	ProdutorType
)

func main() {

	gestorFactory, _ := GetUserFactory(GestorType)

	gestor := gestorFactory.MakeUser(User{
		Nome:      "John",
		Sobrenome: "Smith",
		Password:  "123",
	})

	fmt.Println(gestor)
}
