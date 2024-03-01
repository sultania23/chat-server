package server

import (
	"context"
	"fmt"
	"github.com/kavu/go_reuseport"
	"github.com/tuxoo/idler/internal/config"
	"net/http"
)

const (
	protocol = "tcp"
)

type HTTPServer struct {
	httpServer *http.Server
}

func NewHTTPServer(cfg *config.Config, handler http.Handler) *HTTPServer {
	return &HTTPServer{
		httpServer: &http.Server{
			Addr:           fmt.Sprintf(":%s", cfg.HTTP.Port),
			Handler:        handler,
			MaxHeaderBytes: cfg.HTTP.MaxHeaderMegabytes << 28,
		},
	}
}

func (s *HTTPServer) Run() error {
	listener, err := reuseport.NewReusablePortListener(protocol, s.httpServer.Addr)
	if err != nil {
		return err
	}

	return s.httpServer.Serve(listener)
}

func (s *HTTPServer) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
