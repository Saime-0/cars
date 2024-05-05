package httpv1

import (
	"carstore/internal/usecase"
	"context"
	"fmt"
	"net/http"
	"time"
)

type HttpController struct {
	uc         Usecases
	httpServer *http.Server
}

type Usecases struct { // не стал делать интерфейсы для юзкейсов
	*usecase.CarsUsecase
	*usecase.UpdateCarUsecase
	*usecase.DeleteCarUsecase
	*usecase.AddCarUsecase
}

type Config struct {
	ServerPort string
}

func NewHttpController(cfg Config, uc Usecases) *HttpController {
	var mux = http.NewServeMux()
	httpServer := &http.Server{
		Addr:           "0.0.0.0:" + cfg.ServerPort,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	var httpController = &HttpController{
		httpServer: httpServer,
		uc:         uc,
	}
	registerHttpServerRoutes(mux, httpController)
	return httpController
}

func (s *HttpController) Run(ctx context.Context) error {
	go s.handleContext(ctx)
	err := s.httpServer.ListenAndServe()
	if err != nil && ctx.Err() != context.Canceled {
		return fmt.Errorf("http server: %w", err)
	}
	return nil
}

func (s *HttpController) handleContext(ctx context.Context) {
	select {
	case <-ctx.Done():
		s.httpServer.Shutdown(context.Background())
	}
}

func registerHttpServerRoutes(mux *http.ServeMux, hc *HttpController) {
	// неоднозначный параметр hc можно заменить на общую структуру зависимостей...
	// хендлеров,а существующие методы-хендлеры заменить...
	// функциями принимающими зависимости (юзкейсы или другие небходимые вещи)...
	// и вовзвращающие замыкание (функцию-хендлер)
	mux.HandleFunc("GET /cars", hc.handleGetCars)
	mux.HandleFunc("DELETE /cars/{regNum}", hc.handleDeleteCar)
	mux.HandleFunc("PATH /cars", hc.handleUpdateCar)
	mux.HandleFunc("POST /cars", hc.handleAddCar)
}
