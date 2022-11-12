package abstractfactory

type GestorFactory struct{}

func (g *GestorFactory) MakeUser(user User) IUser {

	return &Gestor{
		User: User{
			user.Nome,
			user.Sobrenome,
			user.Password,
		},
	}
}
