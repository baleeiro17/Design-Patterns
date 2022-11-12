package abstractfactory

type IUser interface {
	getName() string
	getSobrenome() string
	getPassword() string
}

type User struct {
	Nome      string
	Sobrenome string
	Password  string
}
