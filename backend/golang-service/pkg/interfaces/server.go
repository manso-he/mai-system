package interfaces

type Server interface {
	ListenAndServe() error
	Shutdown()
}
