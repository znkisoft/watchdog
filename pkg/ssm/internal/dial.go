package internal

// reference:
// https://serverfault.com/questions/312930/is-ping-a-reliable-way-to-check-if-a-server-is-available
// better way to check server status is to check specified port and not use ICMP ping method

// Dialer is to check server status
type Dialer struct {
	semaphore chan struct{}
}

func NewDialer(max int) *Dialer {
	return &Dialer{
		semaphore: make(chan struct{}, max),
	}
}

func (d *Dialer) Dial(network, addr string) {

}
