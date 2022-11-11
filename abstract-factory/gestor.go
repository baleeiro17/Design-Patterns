package main

type Gestor struct {
	User
}

func (c *Gestor) getName() string {
	return c.nome
}

func (c *Gestor) getSobrenome() string {
	return c.sobrenome
}

func (c *Gestor) getPassword() string {
	return c.password
}
