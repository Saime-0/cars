package httpv1

import (
	"carstore/internal/domain/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func (d *HttpController) handleGetCars(w http.ResponseWriter, r *http.Request) {
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
			err = fmt.Errorf("convert `page` parameter to int: %w", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		pagination.Page = page
	}
	if query.Has("perPage") {
		perPageStr := query.Get("perPage")
		perPage, err := strconv.Atoi(perPageStr)
		if err != nil {
			err = fmt.Errorf("convert `perPage` parameter to int: %w", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		pagination.PerPage = perPage
	}
	cars, err := d.uc.Cars(filter, pagination)
	if err != nil {
		err = fmt.Errorf("get cars via usecase: %w", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	b, err := json.Marshal(cars)
	if err != nil {
		err = fmt.Errorf("marshal response data: %w", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_, err = w.Write(b)
	if err != nil {
		log.Printf("write get cars response: %s", err)
	}
}

func (d *HttpController) handleDeleteCar(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	err := d.uc.DeleteCar(id)
	if err != nil {
		err = fmt.Errorf("delete car via usecase: %w", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (d *HttpController) handleUpdateCar(w http.ResponseWriter, r *http.Request) {
	var input updateCarInput
	reader := json.NewDecoder(r.Body)
	err := reader.Decode(&input)
	if err != nil {
		err = fmt.Errorf("decode request body: %w", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = d.uc.UpdateCar(model.CarUpdate{
		Id:     input.Id,
		RegNum: input.RegNum,
		Mark:   input.Mark,
		Model:  input.Model,
		Owner:  input.Owner,
	})
	if err != nil {
		err = fmt.Errorf("update car via usecase: %w", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}

func (d *HttpController) handleAddCar(w http.ResponseWriter, r *http.Request) {
	var input addCarInput
	reader := json.NewDecoder(r.Body)
	err := reader.Decode(&input)
	if err != nil {
		err = fmt.Errorf("decode request body: %w", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, regNum := range input.RegNums {
		// TODO: mb fetch errors?
		err := d.uc.AddCar(regNum)
		if err != nil {
			err = fmt.Errorf("add car via usecase: %w", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
}
