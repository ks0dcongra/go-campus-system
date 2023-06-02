package cookie

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetJWTTokenCookie(c *gin.Context, token string) {
	cookie := &http.Cookie{
		Name:  "x-csrf-token",
		Value: token,
		// 使得 JavaScript 无法访问该 Cookie，提高安全性。
		HttpOnly: true,
		// 仅在使用 HTTPS 连接时发送该 Cookie。
		Secure: false,
		// 表示对整个网站的所有路径都可见
		Path:   "/",
		MaxAge: 3600, // Expiration time in seconds
		// 仅允许与当前站点具有相同顶级域的请求发送该 Cookie。有 http.SameSiteNoneMode、http.SameSiteLaxMode 和 http.SameSiteStrictMode
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
	if cookie != "pass" {
		return false
	}
	return true
}
