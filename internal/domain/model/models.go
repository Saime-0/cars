package model

type CarCreate struct {
	RegNum string
	Mark   string
	Model  string
	Owner  string
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
