package main

type ClientFactory struct{}

func (c *ClientFactory) makeUser(user User) IUser {

	return &Client{
		User: User{
			user.nome,
			user.sobrenome,
			user.password,
		},
	}
}
