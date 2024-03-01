package server

import (
	"context"
	"fmt"
	"github.com/tuxoo/idler/internal/config"
	"net/http"
)

type WSServer struct {
	wsServer *http.Server
}

func NewWSServer(cfg *config.Config, wsHandler http.Handler) *WSServer {
	wsServer := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.WS.Port),
		Handler: wsHandler,
	}

	return &WSServer{
		wsServer: wsServer,
	}
}

func (s *WSServer) Run() error {
	return s.wsServer.ListenAndServe()
}

func (s *WSServer) Shutdown(ctx context.Context) error {
	return s.wsServer.Shutdown(ctx)
}
