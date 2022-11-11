package main

type Produtor struct {
	User
	saldoVendas float64
}

func (c *Produtor) getName() string {
	return c.nome
}

func (c *Produtor) getSobrenome() string {
	return c.sobrenome
}

func (c *Produtor) getPassword() string {
	return c.password
}

func (c *Produtor) GetSaldoVendas() float64 {
	return c.saldoVendas
}

func (c *Produtor) setSaldoVendas(saldo float64) {
	c.saldoVendas = saldo
}
