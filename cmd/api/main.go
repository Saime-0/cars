package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	carsUsecase := CarsUsecase{
		externalApi: NewExternalApi("https://example.co"),
	}
	deps := Deps{
		CarsUsecase: carsUsecase,
	}
	http.HandleFunc("GET /cars", deps.handleGetCars)
	http.HandleFunc("DELETE /cars/{id}", deleteDataHandler)
	http.HandleFunc("PATH /cars/{id}", updateDataHandler)
	http.HandleFunc("POST /cars/{id}", addDataHandler)
}

type Deps struct {
	CarsUsecase
}

func (d *Deps) handleGetCars(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	filter := CarsFilter{}
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

	pagination := CarsPagination{}
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
		text := fmt.Sprintf("handle get cars: %w", err)
		http.Error(w, text, http.StatusBadRequest)
		return
	}
	b, err := json.Marshal(cars)
	if err != nil {
		text := fmt.Sprintf("handle get cars: %w", err)
		http.Error(w, text, http.StatusBadRequest)
		return
	}
	_, err = w.Write(b)
	if err != nil {
		log.Printf("handle get cars: %w", err)
	}
}

func deleteDataHandler(w http.ResponseWriter, r *http.Request) {
	r.URL.Query().Get()
}

func updateDataHandler(w http.ResponseWriter, r *http.Request) {
	r.URL.Query().Get()
}

func addDataHandler(w http.ResponseWriter, r *http.Request) {
	r.URL.Query().Get()
}

type AddCarUsecase struct{}

func (c *AddCarUsecase) AddCar() error {

}

type UpdateCarUsecase struct{}

func (c *UpdateCarUsecase) UpdateCar(update CarUpdate) error {
	// TODO
}

type DeleteCarUsecase struct{}

func (c *DeleteCarUsecase) DeleteCar(regNum string) error {
	// TODO
}

type CarsUsecase struct {
	externalApi *ExternalApi
}

func (c *CarsUsecase) Cars(filter CarsFilter, pagination CarsPagination) ([]CarDomain, error) {
	// TODO
}

type CarUpdate struct {
	RegNum string
	Mark   *string
	Model  *string
	Owner  *string
}

type CarDomain struct {
	RegNum string
	Mark   string
	Model  string
	Owner  string
}

type CarsFilter struct {
	RegNum *string
	Mark   *string
	Model  *string
	Owner  *string
}

type CarsPagination struct {
	Page    int
	PerPage int
}

type RegNumInfoResponse struct {
	RegNum string
	Mark   string
	Model  string
	Owner  string
}

type ExternalApi struct {
	host string
}

func NewExternalApi(host string) *ExternalApi {
	return &ExternalApi{
		host: host,
	}
}

func (r *ExternalApi) RegNumInfo(regNum string) (*RegNumInfoResponse, error) {
	resp, err := http.Get(r.host + "/info?regNum=" + regNum)
	if err != nil {
		return nil, fmt.Errorf("get data from external api: %w", err)
	}
	var info *infoResponse
	reader := json.NewDecoder(resp.Body)
	err = reader.Decode(info)
	if err != nil {
		return nil, fmt.Errorf("decode response from external api: %w")
	}
	return &RegNumInfoResponse{
		RegNum: info.RegNum,
		Mark:   info.Mark,
		Model:  info.Model,
		Owner:  info.Owner,
	}, nil
}

type infoResponse struct {
	RegNum string `json:"regNum"`
	Mark   string `json:"mark"`
	Model  string `json:"model"`
	Owner  string `json:"owner"`
}
