package test

import (
	"example1/app/http/controller"
	// "example1/app/repository"
	// "example1/app/service"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/stretchr/testify/assert" 
    "github.com/stretchr/testify/mock"
)

// 丟出特定的錯誤

type MyMockedObject struct{
	mock.Mock
}

func (m *MyMockedObject) DoSomething(number int) (bool, error) {
	args := m.Called(number)
	return args.Bool(0), args.Error(1)
}

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
	

	
	// 初始化 controller、service 和 repository
	// repo := &repository.UserRepository{}
	// service := &service.UserService{UserRepository:repo}
	// controller := &controller.UserController{UserService: service}
	// fmt.Println("1",controller)

	
	
	// 初始化測試用的 Gin 引擎和 HTTP 請求、響應
	router := gin.New()
	router.POST("/login", controller.NewUserController().LoginUser())
	// router.GET("/login",controller.LoginUser())
	req := httptest.NewRequest("POST", "/login", nil)
	
	
	// // req := httptest.NewRequest("GET", "user/api/login", nil)
	w := httptest.NewRecorder()
	// // router.ServeHTTP(w, req)
	// req.Header.Set("Authorization", "Bearer testToken")
	// fmt.Println("2",req.Header.Get("Authorization"))
	// handler := func(w http.ResponseWriter, r *http.Request) {
	// 	io.WriteString(w, "hihihi")
	// }
	// handler(w, req)
	router.ServeHTTP(w, req)
	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)
	// fmt.Println("3",req)
	fmt.Println("33333",resp)
	fmt.Println("4",resp.StatusCode)
	fmt.Println("5",resp.Body)
	fmt.Println("7",resp.Header.Get("Content-Type"))
	fmt.Println("8",string(body))

	expectedStatus := http.StatusOK
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
