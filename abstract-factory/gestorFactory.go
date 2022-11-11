package main

type GestorFactory struct{}

func (g *GestorFactory) makeUser(user User) IUser {

	return &Gestor{
		User: User{
			user.nome,
			user.sobrenome,
			user.password,
		},
	}
}
