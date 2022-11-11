package main

type Client struct {
	User
	saldo float64
}

func (c *Client) getName() string {
	return c.nome
}

func (c *Client) getSobrenome() string {
	return c.sobrenome
}

func (c *Client) getPassword() string {
	return c.password
}

func (c *Client) getSaldo() float64 {
	return c.saldo
}

func (c *Client) setSaldo(saldo float64) {
	c.saldo = saldo
}
