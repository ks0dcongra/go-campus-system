package middleware

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

func UserPasd(field validator.FieldLevel) bool {
	// 右邊表達式為開頭必為大寫且長度最小為五最大為十：^[A-Z]\w{4,10}$，右邊表達式為^[a-zA-Z0-9_!@]+$
	if match, _ := regexp.MatchString(`^(.*(?:\b[Aa][Nn][Dd]\b|\b[Oo][Rr]\b|\b[Ss][Ee][Ll][Ee][Cc][Tt]\b|\b[Ww][Hh][Ee][Rr][Ee]\b|\b[Hh][Aa][Vv][Ii][Nn][Gg]\b|\b[Uu][Nn][Ii][Oo][Nn]\b)|.*\d+=\d+|.*[a-zA-Z]+=[a-zA-Z]+|.*--|\s)+.*$`, field.Field().String()); match {
		return false
	}
	return true
}
