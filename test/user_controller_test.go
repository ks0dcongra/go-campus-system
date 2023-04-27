package test

import (
	"bytes"

	
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"

	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/stretchr/testify/assert"
	// "github.com/stretchr/testify/mock"
)

func Test_userController_LoginUser(t *testing.T) {
	tests := []struct {
		name string
		h    *controller.UserController
		want gin.HandlerFunc
	}{
		{
			name: "James",
			h:  controller.NewUserController(),
			want: nil,
		},
		
	}
	// TODO:
	// 初始化測試用的 Gin 引擎和 HTTP 請求、響應
	router := gin.New()
	router.POST("/login", controller.NewUserController().LoginUser())

	// 將 JSON 編碼的數據轉換為 io.Reader 接口，並使用 httptest.NewRequest() 創建一個包含 JSON 內容的 POST 請求
	jsonData := []byte(`{"Name":"Mar234","Password":"12345678"}`)
	
	req := httptest.NewRequest("POST", "/login", bytes.NewBuffer(jsonData))
	req.Header.Set("Authorization", "Bearer testToken")
	// req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	// c, _ := gin.CreateTestContext(w)
	// c.Params = []gin.Param{gin.Param{Key: "k", Value: "v"}}

	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "helloworld")
	}

	handler(w, req)
	router.ServeHTTP(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)
	expectedStatus := http.StatusOK

	fmt.Println("1",req)
	fmt.Println("2",w)
	fmt.Println("3",resp)
	fmt.Println("4",resp.StatusCode)
	fmt.Println("5",resp.Body)
	fmt.Println("7",resp.Header.Get("Content-Type"))
	fmt.Println("8",string(body))
	fmt.Println("9",expectedStatus)
	fmt.Println("10",w.Code)
	assert.Equal(t, expectedStatus, w.Code)
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := "nil"; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userController.LoginUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
