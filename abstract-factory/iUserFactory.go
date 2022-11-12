package abstractfactory

type IUserFactory interface {
	MakeUser(user User) IUser
}

func GetUserFactory(typeUser int) (IUserFactory, error) {

	switch typeUser {

	case 0:
		return &ClientFactory{}, nil

	case 1:
		return &GestorFactory{}, nil

	case 2:
		return &ProdutorFactory{}, nil

	default:
		return nil, nil

	}

}
