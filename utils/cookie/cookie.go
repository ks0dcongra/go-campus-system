package cookie

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func SetJWTTokenCookie(c *gin.Context, token string) {
	cookie := &http.Cookie{
		Name:     "jwt-token",
		Value:    token,
		HttpOnly: false,
		Secure:   false, // Set to true if using HTTPS
		Path:     "/",
		MaxAge:   3600, // Expiration time in seconds
		SameSite: http.SameSiteNoneMode,
	}
	http.SetCookie(c.Writer, cookie)
}

func GetJWTTokenCookie(c *gin.Context) bool {
	cookie, err := c.Cookie("jwt-token")
	if err != nil {
		return false
	}
	if cookie != "pass"{
		return false
	}
	return true
}