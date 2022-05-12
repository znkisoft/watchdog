package ssm

type Monitor struct {
	config string
}

type Plugin interface {
	Exec(chan<- string) error
}
