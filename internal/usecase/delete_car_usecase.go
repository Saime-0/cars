package usecase

import (
	"carstore/internal/data"
	"fmt"
)

type DeleteCarUsecase struct {
	carsRepo data.CarsRepository
}

func NewDeleteCarUsecase(repo data.CarsRepository) *DeleteCarUsecase {
	return &DeleteCarUsecase{
		carsRepo: repo,
	}
}

func (c *DeleteCarUsecase) DeleteCar(id string) error {
	// TODO: mb check id exists?
	err := validateId(id)
	if err != nil {
		return fmt.Errorf("validatrion id: %w", err)
	}
	err = c.carsRepo.Delete(id)
	if err != nil {
		return fmt.Errorf("delete car from repo: %w", err)
	}
	return nil
}
