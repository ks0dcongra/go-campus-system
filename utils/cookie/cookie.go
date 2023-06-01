package cookie

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetJWTTokenCookie(c *gin.Context, token string) {
	cookie := &http.Cookie{
		Name:     "jwt-token",
		Value:    token,
		HttpOnly: true,
		Secure:   true, // Set to true if using HTTPS
		Path:     "/",
		MaxAge:   3600, // Expiration time in seconds
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(c.Writer, cookie)
}

func GetJWTTokenCookie(c *gin.Context) bool {
	cookie, err := c.Cookie("jwt-token")
	log.Println(cookie)
	if err != nil {
		return false
	}
	if cookie != "pass"{
		return false
	}
	return true
}