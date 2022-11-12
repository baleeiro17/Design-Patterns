package abstractfactory

type ClientFactory struct{}

func (c *ClientFactory) MakeUser(user User) IUser {

	return &Client{
		User: User{
			user.Nome,
			user.Sobrenome,
			user.Password,
		},
	}
}
