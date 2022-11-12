package abstractfactory

type Produtor struct {
	User
	saldoVendas float64
}

func (c *Produtor) getName() string {
	return c.Nome
}

func (c *Produtor) getSobrenome() string {
	return c.Sobrenome
}

func (c *Produtor) getPassword() string {
	return c.Password
}

func (c *Produtor) GetSaldoVendas() float64 {
	return c.saldoVendas
}

func (c *Produtor) setSaldoVendas(saldo float64) {
	c.saldoVendas = saldo
}
