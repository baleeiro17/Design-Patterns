package abstractfactory

type Gestor struct {
	User
}

func (c *Gestor) getName() string {
	return c.Nome
}

func (c *Gestor) getSobrenome() string {
	return c.Sobrenome
}

func (c *Gestor) getPassword() string {
	return c.Password
}
