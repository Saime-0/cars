package usecase

import (
	"carstore/internal/data"
	"carstore/internal/domain/model"
	"errors"
	"fmt"
)

type CarsUsecase struct {
	carsRepo data.CarsRepository
}

func NewCarsUsecase(repo data.CarsRepository) *CarsUsecase {
	return &CarsUsecase{
		carsRepo: repo,
	}
}

func (c *CarsUsecase) Cars(filter model.CarsFilter, pagination model.CarsPagination) ([]model.Car, error) {
	err := validateFilter(filter)
	if err != nil {
		return nil, fmt.Errorf("validation filter: %w", err)
	}
	err = validatePagination(pagination)
	if err != nil {
		return nil, fmt.Errorf("validation pagination: %w", err)
	}
	cars, err := c.carsRepo.Cars(filter, pagination)
	if err != nil {
		return nil, fmt.Errorf("get cars from repo: %w", err)
	}
	return cars, nil
}

func validateFilter(filter model.CarsFilter) error {
	if filter.Id != nil {
		err := validateId(*filter.Id)
		if err != nil {
			return err
		}
	}
	if filter.Mark != nil {
		err := validateMark(*filter.Mark)
		if err != nil {
			return err
		}
	}
	if filter.Model != nil {
		err := validateModel(*filter.Model)
		if err != nil {
			return err
		}
	}
	if filter.Owner != nil {
		err := validateOwner(*filter.Owner)
		if err != nil {
			return err
		}
	}
	if filter.RegNum != nil {
		err := validateRegNum(*filter.RegNum)
		if err != nil {
			return err
		}
	}
	return nil
}

var errPageValidation = errors.New("page incorrect")
var errPerPageValidation = errors.New("perPage incorrect")

func validatePagination(pagination model.CarsPagination) error {
	if pagination.Page < 0 {
		return errPageValidation
	}
	if pagination.PerPage < 0 {
		return errPerPageValidation
	}
	return nil
}
