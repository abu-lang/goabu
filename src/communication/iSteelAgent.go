package communication

type ISteelAgent interface {
	Start() error
	Join() error
	// ForAll(...) error TODO
	Stop() error
	IsRunning() bool
}
