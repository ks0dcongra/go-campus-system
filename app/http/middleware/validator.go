package middleware

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

func UserPasd(field validator.FieldLevel) bool {
	// 右邊表達式為開頭必為大寫且長度最小為五最大為十：^[A-Z]\w{4,10}$
	// 下面表達式為
	if match, _ := regexp.MatchString(`^[a-zA-Z0-9_!@]+$`, field.Field().String()); match {
		return true
	}
	return false
}
