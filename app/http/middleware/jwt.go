package middleware

import (
	"example1/utils/global"
	"example1/utils/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}

		// 取得header检查 Token 是否在黑名单中
		tokenString := c.GetHeader("Authorization")	
		if _, ok := global.Blacklist[tokenString]; ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token has been invalidated"})
			c.Abort()
			return
		}
		c.Next()
	}
}
