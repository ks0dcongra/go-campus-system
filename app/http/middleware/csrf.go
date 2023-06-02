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
		// 是用于生成 CSRF 令牌的密钥。请注意，这只是示例密钥，实际应用中应使用更长且更安全的密钥。
		[]byte("32-byte-long-auth-key"),
		// 设置 CSRF 中间件在非 HTTPS 连接上也起作用，
		csrf.Secure(false),
		// 将 CSRF 令牌设置为 HttpOnly，以便 JavaScript 无法访问该令牌，从而提高安全性。
		csrf.HttpOnly(true),
		csrf.SameSite(csrf.SameSiteStrictMode),
		// 设置 CSRF 错误处理程序，当 CSRF 令牌验证失败时调用该处理程序。在此示例中，返回一个包含错误信息的 HTTP 响应。
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
