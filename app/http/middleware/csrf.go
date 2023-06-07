package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/csrf"
	adapter "github.com/gwatts/gin-adapter"
	"net/http"
)

var csrfMd func(http.Handler) http.Handler

func CSRF() gin.HandlerFunc {
	csrfMd = csrf.Protect(
		// 是用於生成 CSRF 令牌的密鑰。請注意，這只是示例密鑰，實際應用中應使用更長且更安全的密鑰。
		[]byte("32-byte-long-auth-key"),
		// 設置 CSRF 中間件在非 HTTPS 連接上也起作用，
		csrf.Secure(false),
		// 將 CSRF 令牌設置為 HttpOnly，以便 JavaScript 無法訪問該令牌，從而提高安全性。
		csrf.HttpOnly(true),
		// 完全禁止第三方的 Cookie 請求，基本上只有在網域和 URL 中的網域相同，才會傳遞 Cookie 請求
		csrf.SameSite(csrf.SameSiteStrictMode),
		// 設置 CSRF 錯誤處理程序，當 CSRF 令牌驗證失敗時調用該處理程序。在此示例中，返回一個包含錯誤信息的 HTTP 響應。
		csrf.ErrorHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte(`{"message": "Forbidden - CSRF token invalid"}`))
		})),
	)
	return adapter.Wrap(csrfMd)
}

func CSRFToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("X-CSRF-Token", csrf.Token(c.Request))
	}
}
