package builder

type Ran3gppBuilder struct {
	ip     string
	access string
	port   int
}

func newRan3gppBuilder() *Ran3gppBuilder {
	return &Ran3gppBuilder{}
}

func (r *Ran3gppBuilder) setIp(ip string) {
	r.ip = ip
}

func (r *Ran3gppBuilder) setAccess() {
	r.access = "3gpp"
}

func (r *Ran3gppBuilder) setPort(port int) {
	r.port = port
}

func (r *Ran3gppBuilder) getRan() Ran {
	return Ran{
		ip:     r.ip,
		port:   r.port,
		access: r.access,
	}
}
