package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)



func main() {
	carsRepo := // todo 
	deps := Deps{
		CarsUsecase:      carsUsecase,
		UpdateCarUsecase: UpdateCarUsecase{
			carsRepo: carsRepo,
		},
		DeleteCarUsecase: DeleteCarUsecase{
			carsRepo: carsRepo,
		},
		AddCarUsecase:    AddCarUsecase{
			externalApi: NewExternalApi("https://example.co"),
			carsRepo:    carsRepo,
		},
	}
	http.HandleFunc("GET /cars", deps.handleGetCars)
	http.HandleFunc("DELETE /cars/{regNum}", deps.handleDeleteCar)
	http.HandleFunc("PATH /cars", deps.handleUpdateCar)
	http.HandleFunc("POST /cars", deps.handleAddCar)
}

type Deps struct {
	CarsUsecase
	UpdateCarUsecase
	DeleteCarUsecase
	AddCarUsecase
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
	var input *UpdateCarInput
	reader := json.NewDecoder(r.Body)
	err := reader.Decode(input)
	if err != nil {
		text := fmt.Sprintf("decode request body: %w", err)
		http.Error(w, text, http.StatusBadRequest)
		return
	}
	err = d.UpdateCar(CarUpdate{
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

type UpdateCarInput struct {
	RegNum string  `json:"regNum"`
	Mark   *string `json:"mark,omitempty"`
	Model  *string `json:"model,omitempty"`
	Owner  *string `json:"owner,omitempty"`
}

func (d *Deps) handleAddCar(w http.ResponseWriter, r *http.Request) {
	var input *AddCarInput
	reader := json.NewDecoder(r.Body)
	err := reader.Decode(input)
	if err != nil {
		text := fmt.Sprintf("decode request body: %w", err)
		http.Error(w, text, http.StatusBadRequest)
		return
	}

	for _, regNum := range input.RegNums {
		// TODO: mb fetch errors
		err := d.AddCar(regNum)
		if err != nil {
			text := fmt.Sprintf("add car via usecase: %w", err)
			http.Error(w, text, http.StatusBadRequest)
			return
		}
	}
}

type AddCarInput struct {
	RegNums []string `json:"regNums"`
}

type AddCarUsecase struct {
	externalApi ExternalApi
	carsRepo    CarsRepository
}

func (c *AddCarUsecase) AddCar(regNum string) error {
	// TODO: validate input
	info, exists, err := c.externalApi.RegNumInfo(regNum)
	if err != nil {
		return fmt.Errorf("update car in repo: %w", err)
	}
	if !exists {
		return fmt.Errorf("not found info in external api")
	}
	err = c.carsRepo.Add(CarCreate{
		RegNum: info.RegNum,
		Mark:   info.Mark,
		Model:  info.Model,
		Owner:  info.Owner,
	})
	if err != nil {
		return fmt.Errorf("add car to repo: %w", err)
	}
	return nil
}

type UpdateCarUsecase struct {
	carsRepo CarsRepository
}

func (c *UpdateCarUsecase) UpdateCar(update CarUpdate) error {
	// TODO: validate fields
	err := c.carsRepo.Update(update)
	if err != nil {
		return fmt.Errorf("update car in repo: %w", err)
	}
	return nil
}

type DeleteCarUsecase struct {
	carsRepo CarsRepository
}

func (c *DeleteCarUsecase) DeleteCar(regNum string) error {
	// TODO: check regNum exists
	err := c.carsRepo.Delete(regNum)
	if err != nil {
		return fmt.Errorf("delete car from repo: %w", err)
	}
	return nil
}

type CarsUsecase struct {
	carsRepo    CarsRepository
}

func (c *CarsUsecase) Cars(filter CarsFilter, pagination CarsPagination) ([]CarDomain, error) {
	// TODO: validation filter, pagination
	cars, err := c.carsRepo.Cars(filter, pagination)
	if err != nil {
		return nil, fmt.Errorf("get cars from repo: %w", err)
	}
	return cars, nil
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

func (r *ExternalApi) RegNumInfo(regNum string) (*RegNumInfoResponse, bool, error) {
	resp, err := http.Get(r.host + "/info?regNum=" + regNum)
	if err != nil {
		return nil, false, fmt.Errorf("get data from external api: %w", err)
	}
	if resp.StatusCode == http.StatusNotFound {
		return nil, false, nil
	}
	var info *infoResponse
	reader := json.NewDecoder(resp.Body)
	err = reader.Decode(info)
	if err != nil {
		return nil, false, fmt.Errorf("decode response from external api: %w")
	}
	return &RegNumInfoResponse{
		RegNum: info.RegNum,
		Mark:   info.Mark,
		Model:  info.Model,
		Owner:  info.Owner,
	}, true, nil
}

type infoResponse struct {
	RegNum string `json:"regNum"`
	Mark   string `json:"mark"`
	Model  string `json:"model"`
	Owner  string `json:"owner"`
}

type CarsRepository interface {
	Cars(CarsFilter, CarsPagination) ([]CarDomain, error)
	Add(CarCreate) error
	Delete(regNum string) error
	Update(CarUpdate) error
}

type CarCreate struct {
	RegNum string
	Mark   string
	Model  string
	Owner  string
}
