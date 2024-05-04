package main

import (
	"carstore/internal/data"
	"carstore/internal/domain/model"
	"carstore/internal/usecase"
	"carstore/lib/extapi"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	repo := data.NewCarRepoBase()
	externalapi := extapi.NewExternalApi("https://example.co")
	deps := Deps{
		CarsUsecase:      usecase.NewCarsUsecase(repo),
		UpdateCarUsecase: usecase.NewUpdateCarUsecase(repo),
		DeleteCarUsecase: usecase.NewDeleteCarUsecase(repo),
		AddCarUsecase:    usecase.NewAddCarUsecase(externalapi, repo),
	}
	http.HandleFunc("GET /cars", deps.handleGetCars)
	http.HandleFunc("DELETE /cars/{regNum}", deps.handleDeleteCar)
	http.HandleFunc("PATH /cars", deps.handleUpdateCar)
	http.HandleFunc("POST /cars", deps.handleAddCar)
}

type Deps struct {
	*usecase.CarsUsecase
	*usecase.UpdateCarUsecase
	*usecase.DeleteCarUsecase
	*usecase.AddCarUsecase
}

func (d *Deps) handleGetCars(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	filter := model.CarsFilter{}
	if query.Has("id") {
		id := query.Get("id")
		filter.Id = &id
	}
	if query.Has("regNum") {
		regNum := query.Get("regNum")
		filter.RegNum = &regNum
	}
	if query.Has("mark") {
		mark := query.Get("mark")
		filter.Mark = &mark
	}
	if query.Has("model") {
		model := query.Get("model")
		filter.Model = &model
	}
	if query.Has("owner") {
		owner := query.Get("owner")
		filter.Owner = &owner
	}

	pagination := model.CarsPagination{}
	if query.Has("page") {
		pageStr := query.Get("page")
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			text := fmt.Sprintf("convert `page` parameter to int: %w", err)
			http.Error(w, text, http.StatusBadRequest)
			return
		}
		pagination.Page = page
	}
	if query.Has("perPage") {
		perPageStr := query.Get("perPage")
		perPage, err := strconv.Atoi(perPageStr)
		if err != nil {
			text := fmt.Sprintf("convert `perPage` parameter to int: %w", err)
			http.Error(w, text, http.StatusBadRequest)
			return
		}
		pagination.PerPage = perPage
	}
	cars, err := d.Cars(filter, pagination)
	if err != nil {
		text := fmt.Sprintf("get cars via usecase: %w", err)
		http.Error(w, text, http.StatusBadRequest)
		return
	}
	b, err := json.Marshal(cars)
	if err != nil {
		text := fmt.Sprintf("marshal response data: %w", err)
		http.Error(w, text, http.StatusBadRequest)
		return
	}
	_, err = w.Write(b)
	if err != nil {
		log.Printf("handle get cars: %w", err)
	}
}

func (d *Deps) handleDeleteCar(w http.ResponseWriter, r *http.Request) {
	regNum := r.PathValue("regNum")
	err := d.DeleteCar(regNum)
	if err != nil {
		text := fmt.Sprintf("delete car via usecase: %w", err)
		http.Error(w, text, http.StatusBadRequest)
		return
	}
}

func (d *Deps) handleUpdateCar(w http.ResponseWriter, r *http.Request) {
	var input *updateCarInput
	reader := json.NewDecoder(r.Body)
	err := reader.Decode(input)
	if err != nil {
		text := fmt.Sprintf("decode request body: %w", err)
		http.Error(w, text, http.StatusBadRequest)
		return
	}
	err = d.UpdateCar(model.CarUpdate{
		RegNum: input.RegNum,
		Mark:   input.Mark,
		Model:  input.Model,
		Owner:  input.Owner,
	})
	if err != nil {
		text := fmt.Sprintf("update car via usecase: %w", err)
		http.Error(w, text, http.StatusBadRequest)
		return
	}

}

type updateCarInput struct {
	Id     string  `json:"id"`
	RegNum *string `json:"regNum"`
	Mark   *string `json:"mark"`
	Model  *string `json:"model"`
	Owner  *string `json:"owner"`
}

func (d *Deps) handleAddCar(w http.ResponseWriter, r *http.Request) {
	var input *addCarInput
	reader := json.NewDecoder(r.Body)
	err := reader.Decode(input)
	if err != nil {
		text := fmt.Sprintf("decode request body: %w", err)
		http.Error(w, text, http.StatusBadRequest)
		return
	}

	for _, regNum := range input.RegNums {
		// TODO: mb fetch errors?
		err := d.AddCar(regNum)
		if err != nil {
			text := fmt.Sprintf("add car via usecase: %w", err)
			http.Error(w, text, http.StatusBadRequest)
			return
		}
	}
}

type addCarInput struct {
	RegNums []string `json:"regNums"`
}
