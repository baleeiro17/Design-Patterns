package abstractfactory

type Client struct {
	User
	saldo float64
}

func (c *Client) getName() string {
	return c.Nome
}

func (c *Client) getSobrenome() string {
	return c.Sobrenome
}

func (c *Client) getPassword() string {
	return c.Password
}

func (c *Client) getSaldo() float64 {
	return c.saldo
}

func (c *Client) setSaldo(saldo float64) {
	c.saldo = saldo
}
