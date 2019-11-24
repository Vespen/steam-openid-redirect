package server

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
)

// Server structure.
type server struct {
	addr    string
	port    string
	handler http.Handler
}

// server.Server.ListenAndServe
func (s *server) ListenAndServe() error {
	addr := fmt.Sprintf("%s:%s", s.addr, s.port)

	logrus.Infof(`Starting server at: "%s"`, addr)

	if err := http.ListenAndServe(addr, s.handler); err != nil {
		return err
	}

	return nil
}

// Creates server with the provided parameters.
func NewServer(addr, port string, handler http.Handler) Server {
	if handler == nil {
		panic("handler cannot be nil")
	}

	return &server{
		addr:    addr,
		port:    port,
		handler: handler,
	}
}
