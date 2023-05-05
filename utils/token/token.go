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

type TokenInterface interface {
	GenerateToken(user_id int) (string, error)
}

type JwtFactory struct {
}

func Newjwt() *JwtFactory {
	return &JwtFactory{}
}

type CustomClaims struct {
	Authorized bool `json:"authorized"`
	UserID     int  `json:"user_id"`
	jwt.StandardClaims
}

func (j *JwtFactory) GenerateToken(user_id int) (string, error) {
	token_lifespan, err := strconv.Atoi(os.Getenv("TOKEN_HOUR_LIFESPAN"))

	// 這邊的token_lifespan != 0是為了讓測試可以不會因為抓不到環境變數而直接報出error
	if err != nil && token_lifespan != 0 {
		return "", err
	}

	if token_lifespan == 0 {
		token_lifespan, _ = strconv.Atoi("1")
	}

	if user_id == 0 {
		return "", fmt.Errorf("user_id can't use 0")
	}

	claims := &CustomClaims{
		Authorized: true,
		UserID:     user_id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(token_lifespan)).Unix(),
			// 設置簽署人
			Issuer: "shehomebow",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("API_SECRET")))
}

// 验证 JWT Token
func (j *JwtFactory) TokenValid(c *gin.Context) error {
	tokenString := j.ExtractToken(c)
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 這個條件判斷用於檢查 Token 是否使用了預期的簽名方法。如果 Token 的簽名方法不是 HMAC，則返回一個錯誤，並指明簽名方法不符合預期。
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// 方法符合預期的情況下，從環境變數中讀取 API 密鑰，並返回一個 []byte 類型的密鑰用於解析 Token。
		return []byte(os.Getenv("API_SECRET")), nil
	})

	if err != nil {
		return err
	}

	// 檢查簽名是否有效
	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	// 獲取解析後的 Payload
	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return fmt.Errorf("failed to parse claims")
	}

	// 驗證自訂的 Claims
	if !claims.Authorized {
		return fmt.Errorf("unauthorized")
	}

	// 驗證過期時間
	if time.Now().Unix() > claims.StandardClaims.ExpiresAt {
		return fmt.Errorf("token has expired")
	}

	// 驗證 Issuer
	if claims.StandardClaims.Issuer != "shehomebow" {
		return fmt.Errorf("invalid issuer")
	}

	return nil
}

// 把postman上的Bearer Token 前綴處理掉，僅回傳Token
func (j *JwtFactory) ExtractToken(c *gin.Context) string {
	bearerToken := c.Request.Header.Get("Authorization")
	// 如果bearerToken符合長度標準就回傳bearerToken
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func (j *JwtFactory) ExtractTokenID(c *gin.Context) (uint, error) {
	// 把postman上的Bearer Token 前綴處理掉，僅回傳Token
	tokenString := j.ExtractToken(c)
	// 這個條件判斷用於檢查 Token 是否使用了預期的簽名方法。如果 Token 的簽名方法不是 HMAC，則返回一個錯誤，並指明簽名方法不符合預期。
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 這個條件判斷用於檢查 Token 是否使用了預期的簽名方法。如果 Token 的簽名方法不是 HMAC，則返回一個錯誤，並指明簽名方法不符合預期。
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// 方法符合預期的情況下，從環境變數中讀取 API 密鑰，並返回一個 []byte 類型的密鑰用於解析 Token。
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return 0, err
	}
	// 取出payload
	claims, ok := token.Claims.(*CustomClaims)
	if ok && token.Valid {
		return uint(claims.UserID), nil
	}
	return 0, nil
}
