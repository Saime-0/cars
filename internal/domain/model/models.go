package model

type Car struct {
	Id     string
	RegNum string
	Mark   string
	Model  string
	Owner  string
}

type CarCreate struct {
	Id     string
	RegNum string
	Mark   string
	Model  string
	Owner  string
}

type CarUpdate struct {
	Id     string
	RegNum *string
	Mark   *string
	Model  *string
	Owner  *string
}

type CarsFilter struct {
	Id     *string
	RegNum *string
	Mark   *string
	Model  *string
	Owner  *string
}

type CarsPagination struct {
	Page    int
	PerPage int
}
