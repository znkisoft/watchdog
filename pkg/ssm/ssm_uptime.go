package ssm

type Uptimer interface {
	Uptime() (int, error)
}

func (p *Plugin) Uptime() (int, error) {
	return 3213129, nil
}
