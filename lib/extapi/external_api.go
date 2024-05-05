package extapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ExternalApi struct {
	host string
}

func NewExternalApi(host string) *ExternalApi {
	return &ExternalApi{
		host: host,
	}
}

type RegNumInfoResponse struct {
	RegNum string
	Mark   string
	Model  string
	Owner  string
}

type infoResponse struct {
	RegNum string `json:"regNum"`
	Mark   string `json:"mark"`
	Model  string `json:"model"`
	Owner  string `json:"owner"`
}

func (r *ExternalApi) RegNumInfo(regNum string) (*RegNumInfoResponse, bool, error) {
	resp, err := http.Get(r.host + "/info?regNum=" + regNum)
	if err != nil {
		return nil, false, fmt.Errorf("get data from external api: %w", err)
	}
	if resp.StatusCode == http.StatusNotFound {
		return nil, false, nil
	}
	var info infoResponse
	reader := json.NewDecoder(resp.Body)
	err = reader.Decode(&info)
	if err != nil {
		return nil, false, fmt.Errorf("decode response from external api: %w", err)
	}
	return &RegNumInfoResponse{
		RegNum: info.RegNum,
		Mark:   info.Mark,
		Model:  info.Model,
		Owner:  info.Owner,
	}, true, nil
}
