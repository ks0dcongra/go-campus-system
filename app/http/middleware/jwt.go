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
		// 創建 JwtFactory 實例
		JwtFactory := token.Newjwt()

		// 驗證Token合法是否
		err := JwtFactory.TokenValid(c)
		if err != nil {
			log.Println("JwtAuthMiddleware() err:",err)
			c.JSON(http.StatusUnauthorized, responses.Status(responses.TokenErr, nil))
			c.Abort()
			return
		}

		// 取得header检查 Token 是否在黑名单中
		tokenString := c.GetHeader("Authorization")
		if _, ok := global.Blacklist[tokenString]; ok {
			c.JSON(http.StatusUnauthorized, responses.Status(responses.TokenExpired, nil))
			c.Abort()
			return
		}
		c.Next()
	}
}
