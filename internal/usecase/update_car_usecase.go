package usecase

import (
	"carstore/internal/data"
	"carstore/internal/domain/model"
	"fmt"
)

type UpdateCarUsecase struct {
	carsRepo data.CarsRepository
}

func NewUpdateCarUsecase(repo data.CarsRepository) *UpdateCarUsecase {
	return &UpdateCarUsecase{
		carsRepo: repo,
	}
}

func (c *UpdateCarUsecase) UpdateCar(update model.CarUpdate) error {
	err := validateUpdate(update)
	if err != nil {
		return fmt.Errorf("validation update: %w", err)
	}
	err = c.carsRepo.Update(update)
	if err != nil {
		return fmt.Errorf("update car in repo: %w", err)
	}
	return nil
}

func validateUpdate(update model.CarUpdate) error {
	err := validateId(update.Id)
	if err != nil {
		return err
	}
	if update.Mark != nil {
		err := validateMark(*update.Mark)
		if err != nil {
			return err
		}
	}
	if update.Model != nil {
		err := validateModel(*update.Model)
		if err != nil {
			return err
		}
	}
	if update.Owner != nil {
		err := validateOwner(*update.Owner)
		if err != nil {
			return err
		}
	}
	if update.RegNum != nil {
		err := validateRegNum(*update.RegNum)
		if err != nil {
			return err
		}
	}
	return nil
}
