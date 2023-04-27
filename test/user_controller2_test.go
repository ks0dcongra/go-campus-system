package test

import (
	"bytes"
	"encoding/json"
	"example1/app/http/controller"
	"example1/app/model"
	
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"

	// "net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type EmployeeRepoImplMock struct {
    mock.Mock
}

func (serviceMock *EmployeeRepoImplMock) Login(login *model.LoginStudent) (model.Student, string){
    args := serviceMock.Called(login)
    log.Println("zzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz",args.Get(0))
	
    return args.Get(0).(model.Student),"0"
}

func Test_userController_LoginUser2(t *testing.T) {
	serviceMock := new(EmployeeRepoImplMock)
    serviceMock.On("Login", mock.Anything).
        Return(model.Student{Id: 99, Name: "Jack", Password: "12345678"},"0")
    

    controller := controller.UserController{
        UserService: serviceMock,
    }
	
    
    
	PostBody := map[string]interface{}{
		"Name":"Mar234",
		"Password":"12345678",
	}
	gin.SetMode(gin.TestMode)

	body, _ := json.Marshal(PostBody)
	// create a request to test the controller
	req,_ := http.NewRequest("POST", "/user/api/login", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	// create a ResponseRecorder to record the response
    rr := httptest.NewRecorder()
	// create a fake gin.Context
    c, router := gin.CreateTestContext(rr)
    c.Request = req
	// router := gin.New()
	
	var m map[string]interface{}
    err := json.NewDecoder(req.Body).Decode(&m)
	router.POST("/user/api/login", controller.LoginUser())
	// call the controller function
	router.ServeHTTP(rr, c.Request)
    // req.Body.Close()
    fmt.Println(err, m)

	log.Println(rr)
	log.Println(rr.Code)
	log.Println(rr.HeaderMap)
	log.Println(rr.Body)
	log.Println(rr.Header())
    // actial := controller.LoginUser
    // log.Println("AAA", actial)
	status := "0"
    assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, status, rr.Code)
	// var response responses.Response
	// err := json.Unmarshal(rr.Body.Bytes(), &response)
    // assert.Nil(t, err)
    // assert.Equal(t, responses.Success, response.Status)
}

// use mock to implments CalculatorService's method
// func (repoMock *EmployeeRepoImplMock) FindEmployeesAgeGreaterThan(age int) []model.Employee {
	
// 	args := repoMock.Called(age)
// 	log.Println("argsargsargsargsargsargsargsargsargsargsargs",args.Get(0).([]model.Employee))
// 	return args.Get(0).([]model.Employee)
// }

// func TestGetSrEmployeeNumbers_Age400(t *testing.T) {

// 	repoMock := new(EmployeeRepoImplMock)
// 	repoMock.On("FindEmployeesAgeGreaterThan", 40).
// 		Return([]model.Employee{
// 			{ID: 99, Name: "Jack", Age: 70},
// 		})

	

// 	es := controller.EmployeeSerivceImpl{
// 		EmpRepo: repoMock,
// 	}

// 	actial := es.GetSrEmployeeNumbers(40)
//     log.Println("BBB",actial)
// 	expected := 1
// 	assert.Equal(t, expected, actial)
// }


// package controller_test

// import (
//     "example1/app/controller"
//     "example1/app/model"
//     "example1/app/service"
//     "example1/utils/global"
//     "github.com/gin-gonic/gin"
//     "github.com/stretchr/testify/assert"
//     "github.com/stretchr/testify/mock"
//     "net/http"
//     "net/http/httptest"
//     "testing"
// )

// type MockUserService struct {
//     mock.Mock
// }

// func (m *MockUserService) Login(email string, password string) (*model.User, error) {
//     args := m.Called(email, password)
//     return args.Get(0).(*model.User), args.Error(1)
// }

// func TestLogin(t *testing.T) {
//     gin.SetMode(gin.TestMode)

//     // 建立模擬物件
//     mockService := &MockUserService{}

//     // 設定模擬物件預期的回傳值
//     mockUser := &model.User{
//         ID:       1,
//         Name:     "John",
//         Email:    "john@example.com",
//         Password: "$2a$10$7atbux/ZKjRcQdpfLxy.xuhCddiBvkwG9XamKTB.5.IV7TX.j1fFe", // 123456
//     }
//     mockService.On("Login", "john@example.com", "123456").Return(mockUser, nil)

//     // 初始化 controller 物件，並將上述模擬物件注入 controller
//     controller := &controller.UserController{
//         UserService: mockService,
//     }

//     // 建立 HTTP 請求
//     req, _ := http.NewRequest("POST", "/api/v1/login", nil)
//     req.Header.Set("Content-Type", "application/json")
//     req.SetBasicAuth("john@example.com", "123456")
//     w := httptest.NewRecorder()

//     // 執行測試
//     router := gin.Default()
//     router.POST("/api/v1/login", controller.Login)
//     router.ServeHTTP(w, req)

//     // 驗證回應
//     assert.Equal(t, http.StatusOK, w.Code)
//     assert.JSONEq(t, `{
//         "code": 200,
//         "message": "success",
//         "data": {
//             "id": 1,
//             "name": "John",
//             "email": "john@example.com",
//             "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTY5MTkyMDIsInVzZXJfaWQiOjEsInVzZXJuYW1lIjoiSm9obiJ9.I8nMnZ5QXZmUvV7RAB0-ZUyA6QlffxIqV7Y-8tnU1V0"
//         }
//     }`, w.Body.String())

//     // 驗證模擬物件有被正確呼叫
//     mockService.AssertExpectations(t)
// }