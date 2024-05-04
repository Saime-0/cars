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

func (c *CarsUsecase) Cars(filter model.CarsFilter, pagination model.CarsPagination) ([]model.CarDomain, error) {
	err := validateFilter(filter)
	if err != nil {
		return nil, fmt.Errorf("validation filter: %w", err)
	}
	err = validatePagination(pagination)
	if err != nil {
		return nil, fmt.Errorf("validation filter: %w", err)
	}
	cars, err := c.carsRepo.Cars(filter, pagination)
	if err != nil {
		return nil, fmt.Errorf("get cars from repo: %w", err)
	}
	return cars, nil
}

var errMarkValidation = errors.New("mark incorrect")
var errModelValidation = errors.New("model incorrect")
var errOwnerValidation = errors.New("owner incorrect")

func validateFilter(filter model.CarsFilter) error {
	if filter.Mark != nil && len(*filter.Mark) < 1 {
		return errMarkValidation
	}
	if filter.Model != nil && len(*filter.Model) < 1 {
		return errModelValidation
	}
	if filter.Owner != nil && len(*filter.Owner) < 1 {
		return errOwnerValidation
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
