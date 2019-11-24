package server

// Server structure.
type Server interface {
	// Listens and serves.
	ListenAndServe() error
}
