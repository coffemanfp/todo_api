package server

type Engine interface {
	Run(addr ...string) error
}
