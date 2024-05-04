package data

import (
	"carstore/internal/domain/model"
)

type CarsRepository interface {
	Cars(model.CarsFilter, model.CarsPagination) ([]model.Car, error)
	// CarById(id string) (car *model.Car, exist bool, err error)
	Add(model.CarCreate) error
	Delete(id string) error
	Update(model.CarUpdate) error
}

var _ CarsRepository = (*CarRepoBase)(nil)

type CarRepoBase struct{}

func NewCarRepoBase() *CarRepoBase {
	return &CarRepoBase{}
}

// Add implements CarsRepository.
func (c *CarRepoBase) Add(model.CarCreate) error {
	panic("unimplemented")
}

// Cars implements CarsRepository.
func (c *CarRepoBase) Cars(model.CarsFilter, model.CarsPagination) ([]model.Car, error) {
	panic("unimplemented")
}

// Delete implements CarsRepository.
func (c *CarRepoBase) Delete(id string) error {
	panic("unimplemented")
}

// Update implements CarsRepository.
func (c *CarRepoBase) Update(model.CarUpdate) error {
	panic("unimplemented")
}
