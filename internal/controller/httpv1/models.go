package httpv1

type updateCarInput struct {
	Id     string  `json:"id"`
	RegNum *string `json:"regNum"`
	Mark   *string `json:"mark"`
	Model  *string `json:"model"`
	Owner  *string `json:"owner"`
}

type addCarInput struct {
	RegNums []string `json:"regNums"`
}
