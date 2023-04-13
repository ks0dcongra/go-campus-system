package middleware

import (
	"example1/app/model/responses"
	"example1/utils/global"
	"example1/utils/token"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 驗證Token合法是否
		token, err := token.TokenValid(c)
		if err != nil {
			log.Println("err:",err)
			c.JSON(http.StatusUnauthorized, responses.Status(responses.TokenErr, err))
			c.Abort()
			return
		}
	
		// 檢查簽名是否有效
		if !token.Valid {
			c.JSON(http.StatusUnauthorized, responses.Status(responses.TokenErr, gin.H{"message":"token signature invalid"}))
			c.Abort()
			return
		}

		// 取得header检查 Token 是否在黑名单中
		tokenString := c.GetHeader("Authorization")
		if _, ok := global.Blacklist[tokenString]; ok {
			c.JSON(http.StatusUnauthorized, responses.Status(responses.TokenInvalid, nil))
			c.Abort()
			return
		}
		c.Next()
	}
}
