package usecase

import (
	"carstore/internal/data"
	"carstore/internal/domain/model"
	"carstore/lib/extapi"
	"fmt"
	"github.com/google/uuid"
)

type AddCarUsecase struct {
	externalApi *extapi.ExternalApi
	carsRepo    data.CarsRepository
}

func NewAddCarUsecase(externalapi *extapi.ExternalApi, repo data.CarsRepository) *AddCarUsecase {
	return &AddCarUsecase{
		externalApi: externalapi,
		carsRepo:    repo,
	}
}

func (c *AddCarUsecase) AddCar(regNum string) error {
	err := validateRegNum(regNum)
	if err != nil {
		return fmt.Errorf("validation car number: %w", err)
	}
	info, exists, err := c.externalApi.RegNumInfo(regNum)
	if err != nil {
		return fmt.Errorf("update car in repo: %w", err)
	}
	if !exists {
		return fmt.Errorf("not found info in external api")
	}
	err = c.carsRepo.Add(model.CarCreate{
		Id:     uuid.NewString(),
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
