package builder

type Director struct {
	builder IBuilder
}

func NewDirector(b IBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) setBuilder(b IBuilder) {
	d.builder = b
}

func (d *Director) BuildRan(ip string, port int) Ran {
	d.builder.setIp(ip)
	d.builder.setPort(port)
	d.builder.setAccess()
	return d.builder.getRan()
}
