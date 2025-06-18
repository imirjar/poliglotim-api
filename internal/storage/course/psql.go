package course

type Psql struct {
	Conn string
}

func (p *Psql) Ping() error {
	return nil
}

func (p *Psql) Connect() error {
	return nil
}
