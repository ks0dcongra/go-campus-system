package middleware

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/csrf"
)

var csrfMd func(http.Handler) http.Handler

func CSRF() gin.HandlerFunc {
	return func(c *gin.Context){
		csrfMd = csrf.Protect(
			[]byte("32-byte-long-auth-key"),
			csrf.Secure(false),
			csrf.HttpOnly(true),
			csrf.ErrorHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusForbidden)
				w.Write([]byte(`{"message": "Forbidden - CSRF token invalid"}`))
			})),
		)
	}
}

func CSRFToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("X-CSRF-Token",csrf.Token(c.Request))
	}
}