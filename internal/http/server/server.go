package server

import (
	"net/http"

	"github.com/morfin60/parallel-handlers/internal/config"
)

type Server struct {
	server *http.Server
	config *config.Config
}

func New(config *config.Config) *Server {
	mux := http.NewServeMux()

	return &Server{
		server: &http.Server{Addr: config.Address, Handler: mux},
		config: config,
	}
}

// Start serving httpp
func (s *Server) Start() error {
	return s.server.ListenAndServe()
}

func (s *Server) Handler() http.Handler {
	return s.server.Handler
}
