package middleware

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

func UserPasd(field validator.FieldLevel) bool {
	if match, _ := regexp.MatchString(`^[A-Z]\w{4,10}$`, field.Field().String()); match {
		return true
	}
	return false
}
