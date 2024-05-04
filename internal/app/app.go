package app

import (
	"carstore/internal/controller/httpv1"
	"carstore/internal/data"
	"carstore/internal/usecase"
	"carstore/lib/extapi"
	"context"
	"fmt"
	"log"
	"sync"
)

func Run(ctx context.Context) error {
	repo := data.NewCarRepoBase()
	externalapi := extapi.NewExternalApi("https://example.co")
	uc := usecasesImpls{
		CarsUsecase:      usecase.NewCarsUsecase(repo),
		UpdateCarUsecase: usecase.NewUpdateCarUsecase(repo),
		DeleteCarUsecase: usecase.NewDeleteCarUsecase(repo),
		AddCarUsecase:    usecase.NewAddCarUsecase(externalapi, repo),
	}
	app := newApplication(ctx, uc)
	hc := httpv1.NewHttpController(httpv1.Usecases{
		CarsUsecase:      uc.CarsUsecase,
		UpdateCarUsecase: uc.UpdateCarUsecase,
		DeleteCarUsecase: uc.DeleteCarUsecase,
		AddCarUsecase:    uc.AddCarUsecase,
	})
	go app.runController(hc, "HttpController")
	return app.gracefulShutdownApplication()
}

func (a *application) runController(c Controller, name string) {
	// a.wg.Add(1)
	// defer a.wg.Done()
	log.Printf("starting %s", name)
	err := c.Run(a.ctx)
	if err != nil {
		a.errorf("%s controller run: %s", name, err)
	}
	log.Printf("stopped %s", name)
}

type Controller interface {
	Run(ctx context.Context) error
}

func (a *application) gracefulShutdownApplication() error {
	var err error
	select {
	case <-a.ctx.Done():
		log.Println("ApplicationReceiveCtxDone")
	case err = <-a.errCh:
		a.cancelFunc()
		log.Println("ApplicationReceiveInternalError")
	}
	// a.wg.Wait()
	return err
}

func (a *application) errorf(format string, args ...any) {
	select {
	case a.errCh <- fmt.Errorf(format, args...):
	default:
	}
}

type usecasesImpls struct {
	*usecase.CarsUsecase
	*usecase.UpdateCarUsecase
	*usecase.DeleteCarUsecase
	*usecase.AddCarUsecase
}

type application struct {
	usecases usecasesImpls

	errCh      chan error
	wg         *sync.WaitGroup // for indicate all things (servers, handlers...) will stopped
	cancelFunc context.CancelFunc
	ctx        context.Context
}

func newApplication(ctx context.Context, usecases usecasesImpls) *application {
	ctx, cancelFunc := context.WithCancel(ctx)
	return &application{
		usecases:   usecases,
		errCh:      make(chan error),
		wg:         new(sync.WaitGroup),
		cancelFunc: cancelFunc,
		ctx:        ctx,
	}
}
