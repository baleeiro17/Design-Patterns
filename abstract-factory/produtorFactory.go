package abstractfactory

type ProdutorFactory struct{}

func (p *ProdutorFactory) MakeUser(user User) IUser {

	return &Produtor{
		User: User{
			user.Nome,
			user.Sobrenome,
			user.Password,
		},
	}
}
