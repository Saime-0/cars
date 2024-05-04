package usecase

import (
	"errors"
	"regexp"
)

var regNumRegexp = regexp.MustCompile("/^[АВЕКМНОРСТУХ]\\d{3}(?<!000)[АВЕКМНОРСТУХ]{2}\\d{2,3}$/ui")
var errRegNumValidation = errors.New("regNum incorrect!")

func validateRegNum(regNum string) error {
	if !regNumRegexp.Match([]byte(regNum)) {
		return errRegNumValidation
	}
	return nil
}
