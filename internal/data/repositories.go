package data

import (
	"carstore/internal/domain/model"
	"errors"
	"math/rand"
	"sync"
	"time"
)

type CarsRepository interface {
	Cars(model.CarsFilter, model.CarsPagination) ([]model.Car, error)
	// CarById(id string) (car *model.Car, exist bool, err error)
	Add(model.CarCreate) error
	Delete(id string) error
	Update(model.CarUpdate) error
}

var _ CarsRepository = (*CarRepoBase)(nil)

type CarRepoBase struct {
	cars []model.Car
	mu   *sync.Mutex
}

func NewCarRepoBase() *CarRepoBase {
	return &CarRepoBase{
		cars: []model.Car{},
		mu:   new(sync.Mutex),
	}
}

// Add implements CarsRepository.
func (c *CarRepoBase) Add(newCar model.CarCreate) error {
	mayDelay()
	err := mayError()
	if err != nil {
		return err
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cars = append(c.cars, model.Car(newCar))
	return nil
}

// Cars implements CarsRepository.
func (c *CarRepoBase) Cars(f model.CarsFilter, p model.CarsPagination) ([]model.Car, error) {
	mayDelay()
	err := mayError()
	if err != nil {
		return nil, err
	}

	result := []model.Car{}
	c.mu.Lock()
	defer c.mu.Unlock()
	for _, car := range c.cars {
		if f.Id == nil || *f.Id == car.Id &&
			f.RegNum == nil || *f.RegNum == car.RegNum &&
			f.Mark == nil || *f.Mark == car.Mark &&
			f.Model == nil || *f.Model == car.Model &&
			f.Owner == nil || *f.Owner == car.Owner {
			result = append(result, car)
		}
	}
	return paginate(result, p.Page, p.PerPage), nil
}

func (c *CarRepoBase) Delete(id string) error {
	mayDelay()
	err := mayError()
	if err != nil {
		return err
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	for i := 0; i < len(c.cars); i++ {
		if id == c.cars[i].Id {
			c.cars = append(c.cars[:i], c.cars[i+1:]...)
			break
		}
	}
	return nil
}

func (c *CarRepoBase) Update(upd model.CarUpdate) error {
	mayDelay()
	err := mayError()
	if err != nil {
		return err
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	for i := 0; i < len(c.cars); i++ {
		if upd.Id == c.cars[i].Id {
			if upd.RegNum != nil {
				c.cars[i].RegNum = *upd.RegNum
			}
			if upd.Mark != nil {
				c.cars[i].Mark = *upd.Mark
			}
			if upd.Model != nil {
				c.cars[i].Model = *upd.Model
			}
			if upd.Owner != nil {
				c.cars[i].Owner = *upd.Owner
			}
			break
		}
	}
	return nil
}

var ErrSomeTest = errors.New("some test error")

func mayError() error {
	if rand.Intn(5) == 0 {
		return ErrSomeTest
	}
	return nil
}
func mayDelay() {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
}
