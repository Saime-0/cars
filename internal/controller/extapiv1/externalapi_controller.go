package extapiv1

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type ExternalApiController struct {
	httpServer *http.Server
}

type Config struct {
	ServerPort string
}

func NewExternalApiController(cfg Config) *ExternalApiController {
	var mux = http.NewServeMux()
	httpServer := &http.Server{
		Addr:           "0.0.0.0:" + cfg.ServerPort,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	var eac = &ExternalApiController{
		httpServer: httpServer,
	}
	mux.HandleFunc("GET /info", eac.handleInfo)
	return eac
}

func (s *ExternalApiController) Run(ctx context.Context) error {
	go s.handleContext(ctx)
	err := s.httpServer.ListenAndServe()
	if err != nil && ctx.Err() != context.Canceled {
		return fmt.Errorf("http server: %w", err)
	}
	return nil
}

func (s *ExternalApiController) handleContext(ctx context.Context) {
	select {
	case <-ctx.Done():
		s.httpServer.Shutdown(context.Background())
	}
}
