package usecase

import (
	"carstore/internal/data"
	"carstore/internal/domain/model"
	"fmt"
)

type UpdateCarUsecase struct {
	carsRepo data.CarsRepository
}

func (c *UpdateCarUsecase) UpdateCar(update model.CarUpdate) error {
	// TODO: validate fields
	err := c.carsRepo.Update(update)
	if err != nil {
		return fmt.Errorf("update car in repo: %w", err)
	}
	return nil
}
