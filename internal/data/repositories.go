package data

import (
	"carstore/internal/domain/model"
)

type CarsRepository interface {
	Cars(model.CarsFilter, model.CarsPagination) ([]model.CarDomain, error)
	Add(model.CarCreate) error
	Delete(regNum string) error
	Update(model.CarUpdate) error
}

var _ CarsRepository = (*CarRepoBase)(nil)

type CarRepoBase struct{}

// Add implements CarsRepository.
func (c *CarRepoBase) Add(model.CarCreate) error {
	panic("unimplemented")
}

// Cars implements CarsRepository.
func (c *CarRepoBase) Cars(model.CarsFilter, model.CarsPagination) ([]model.CarDomain, error) {
	panic("unimplemented")
}

// Delete implements CarsRepository.
func (c *CarRepoBase) Delete(regNum string) error {
	panic("unimplemented")
}

// Update implements CarsRepository.
func (c *CarRepoBase) Update(model.CarUpdate) error {
	panic("unimplemented")
}
