package usecase

import (
	"carstore/internal/data"
	"fmt"
)

type DeleteCarUsecase struct {
	carsRepo data.CarsRepository
}

func (c *DeleteCarUsecase) DeleteCar(regNum string) error {
	// TODO: check regNum exists
	err := c.carsRepo.Delete(regNum)
	if err != nil {
		return fmt.Errorf("delete car from repo: %w", err)
	}
	return nil
}
