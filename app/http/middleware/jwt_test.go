package middleware_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"example1/app/http/middleware"
	"example1/app/model/responses"
	"example1/utils/global"
	"example1/utils/token"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestJwtAuthMiddleware(t *testing.T) {
	// Define test cases
	tests := []struct {
		name           string
		authHeader     string
		isBlacklisted  bool
		expectedStatus int
		expectedBody   responses.Response
	}{
		{
			name:           "Authorized access",
			authHeader:     fmt.Sprintf("Bearer %v", generateJwtToken()),
			isBlacklisted:  false,
			expectedStatus: http.StatusOK,
			expectedBody: responses.Response{
				Status:  responses.Success,
				Message: "Success",
			},
		},
		{
			name:           "Unauthorized access",
			authHeader:     "invalid_token",
			isBlacklisted:  false,
			expectedStatus: http.StatusUnauthorized,
			expectedBody: responses.Response{
				Status:  responses.TokenErr,
				Message: "Token issue, can't pass authorization",
			},
		},
		{
			name:           "Expired token",
			authHeader:     fmt.Sprintf("Bearer %v", generateJwtToken()),
			isBlacklisted:  true,
			expectedStatus: http.StatusUnauthorized,
			expectedBody: responses.Response{
				Status:  responses.TokenExpired,
				Message: "Token has been Expired because blaklist has this token",
			},
		},
	}
	
	gin.SetMode(gin.TestMode)

	router := gin.New()
	

	router.Use(middleware.JwtAuthMiddleware())
	{
		router.GET("/test", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status":  responses.Success,
				"message": "Success",
			})
		})
	}

	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.isBlacklisted {
				global.Blacklist[tt.authHeader] = true
			}

			req, err := http.NewRequest("GET", "/test", nil)
			assert.NoError(t, err)

			req.Header.Set("Authorization", tt.authHeader)

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			// 驗證status code
			assert.Equal(t, tt.expectedStatus, w.Code)
			
			var body responses.Response
			err = json.Unmarshal(w.Body.Bytes(), &body)

			// 驗證unmarshal的過程是否有error
			assert.NoError(t, err)
			// 驗證預期的body與response的body是否相同
			assert.Equal(t, tt.expectedBody, body)
			// 刪除當前測試案例中所使用的 JWT token。這麼做是為了確保每個測試案例之間的獨立性，以免前一個測試案例的黑名單 token 影響到後一個測試案例。
			delete(global.Blacklist, tt.authHeader)
		})
	}
}

func generateJwtToken() string {
	jwtFactory := token.Newjwt()
	token, _ := jwtFactory.GenerateToken(1)
	return token
}