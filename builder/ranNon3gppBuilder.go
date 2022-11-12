package builder

type RanNon3gppBuilder struct {
	ip     string
	access string
	port   int
}

func newRanNon3gppBuilder() *RanNon3gppBuilder {
	return &RanNon3gppBuilder{}
}

func (r *RanNon3gppBuilder) setIp(ip string) {
	r.ip = ip
}

func (r *RanNon3gppBuilder) setAccess() {
	r.access = "non3gpp"
}

func (r *RanNon3gppBuilder) setPort(port int) {
	r.port = port
}

func (r *RanNon3gppBuilder) getRan() Ran {
	return Ran{
		ip:     r.ip,
		port:   r.port,
		access: r.access,
	}
}
