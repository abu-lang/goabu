package goabu

type Agent interface {
	Start() error
	Join() error
	ForAll([]byte) error
	ReceivedActions() (<-chan chan []byte, <-chan chan string)
	Stop() error
	IsRunning() bool
	SetLogLevel(int)
}
