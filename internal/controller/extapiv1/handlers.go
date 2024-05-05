package extapiv1

import (
	"encoding/json"
	"fmt"
	"github.com/brianvoe/gofakeit/v7"
	"log"
	"net/http"
)

type carInfo struct {
	RegNum string `json:"regNum"`
	Mark   string `json:"mark"`
	Model  string `json:"model"`
	Owner  string `json:"owner"`
}

func (d *ExternalApiController) handleInfo(w http.ResponseWriter, r *http.Request) {
	if !r.URL.Query().Has("regNum") {
		err := fmt.Errorf("query parameter 'regNum' is required")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	regNum := r.URL.Query().Get("regNum")
	responseData := carInfo{
		RegNum: regNum,
		Mark:   gofakeit.CarMaker(),
		Model:  gofakeit.CarModel(),
		Owner:  gofakeit.Name(),
	}

	b, err := json.Marshal(responseData)
	if err != nil {
		err = fmt.Errorf("marshal response data: %w", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_, err = w.Write(b)
	if err != nil {
		log.Printf("handle info: %w", err)
	}
}
