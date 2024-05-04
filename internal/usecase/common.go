package usecase

import (
	"errors"
	"regexp"
)

var errIdValidation = errors.New("id incorrect")
var errRegNumValidation = errors.New("regNum incorrect")
var errMarkValidation = errors.New("mark incorrect")
var errModelValidation = errors.New("model incorrect")
var errOwnerValidation = errors.New("owner incorrect")

var regNumRegexp = regexp.MustCompile("^[АВЕКМНОРСТУХ][0-9]{3}[АВЕКМНОРСТУХ]{2}[0-9]{2,3}$")

// "^[АВЕКМНОРСТУХ]\\d{3}(?<!000)[АВЕКМНОРСТУХ]{2}\\d{2,3}$")
func validateRegNum(regNum string) error {
	if !regNumRegexp.Match([]byte(regNum)) {
		return errRegNumValidation
	}
	return nil
}

func validateId(id string) error {
	if len(id) < 1 {
		return errIdValidation
	}
	return nil
}
func validateMark(mark string) error {
	if len(mark) < 1 {
		return errMarkValidation
	}
	return nil
}
func validateModel(model string) error {
	if len(model) < 1 {
		return errModelValidation
	}
	return nil
}
func validateOwner(owner string) error {
	if len(owner) < 1 {
		return errOwnerValidation
	}
	return nil
}
