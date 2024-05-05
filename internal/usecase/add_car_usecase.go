package usecase

import (
	"carstore/internal/data"
	"carstore/internal/domain/model"
	"carstore/lib/extapi"
	"fmt"
	"sync/atomic"
)

type AddCarUsecase struct {
	externalApi *extapi.ExternalApi
	carsRepo    data.CarsRepository
	atomId      atomic.Int32
}

func NewAddCarUsecase(externalapi *extapi.ExternalApi, repo data.CarsRepository) *AddCarUsecase {
	return &AddCarUsecase{
		externalApi: externalapi,
		carsRepo:    repo,
		atomId:      atomic.Int32{},
	}
}

func (c *AddCarUsecase) AddCar(regNum string) error {
	err := validateRegNum(regNum)
	if err != nil {
		return fmt.Errorf("validation car number: %w", err)
	}
	info, exists, err := c.externalApi.RegNumInfo(regNum)
	if err != nil {
		return fmt.Errorf("get info from external api: %w", err)
	}
	if !exists {
		return fmt.Errorf("not found info in external api")
	}
	id := c.atomId.Add(1)
	err = c.carsRepo.Add(model.CarCreate{
		Id:     fmt.Sprintf("%03d", id),
		RegNum: info.RegNum,
		Mark:   info.Mark,
		Model:  info.Model,
		Owner:  info.Owner,
	})
	if err != nil {
		return fmt.Errorf("add car to repo: %w", err)
	}
	return nil
}
