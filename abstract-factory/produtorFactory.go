package main

type ProdutorFactory struct{}

func (p *ProdutorFactory) makeUser(user User) IUser {

	return &Produtor{
		User: User{
			user.nome,
			user.sobrenome,
			user.password,
		},
	}
}
