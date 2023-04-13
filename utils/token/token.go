package token

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func GenerateToken(user_id int) (string, error) {

	token_lifespan, err := strconv.Atoi(os.Getenv("TOKEN_HOUR_LIFESPAN"))

	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{} 
	claims["authorized"] = true
	claims["user_id"] = user_id
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(token_lifespan)).Unix() // 设置过期时间
	claims["iss"] = "shehomebow"                                                     // 設置簽署人
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("API_SECRET")))

}
// JWT 只要驗這些嗎
// What is Unit Test with golang? 

// 验证 JWT Token
func TokenValid(c *gin.Context) (*jwt.Token, error) {
	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 這個條件判斷用於檢查 Token 是否使用了預期的簽名方法。如果 Token 的簽名方法不是 HMAC，則返回一個錯誤，並指明簽名方法不符合預期。
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// 方法符合預期的情況下，從環境變數中讀取 API 密鑰，並返回一個 []byte 類型的密鑰用於解析 Token。
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}

//  把postman上的Bearer Token 前綴處理掉，僅回傳Token
func ExtractToken(c *gin.Context) string {
	bearerToken := c.Request.Header.Get("Authorization")
	// 如果bearerToken符合長度標準就回傳bearerToken
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func ExtractTokenID(c *gin.Context) (uint, error) {
	// 把postman上的Bearer Token 前綴處理掉，僅回傳Token
	tokenString := ExtractToken(c)
	// 這個條件判斷用於檢查 Token 是否使用了預期的簽名方法。如果 Token 的簽名方法不是 HMAC，則返回一個錯誤，並指明簽名方法不符合預期。
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return 0, err
	}
	// 取出payload
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 32)
		if err != nil {
			return 0, err
		}
		return uint(uid), nil
	}
	return 0, nil
}
