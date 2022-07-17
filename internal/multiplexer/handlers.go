package multiplexer

import (
	"net/http"

	"github.com/morfin60/parallel-handlers/internal/config"
	"github.com/morfin60/parallel-handlers/internal/http/middleware"
)

func RegisterHandlers(handler http.Handler, cfg *config.Config) {
	checkUrlsHandler := middleware.LimitRequests(MultiplexingHandler(cfg.PoolSize), cfg.RequestsLimit)

	handler.(*http.ServeMux).HandleFunc("/multiplex", checkUrlsHandler)
}
