package controller_test

import (
	"bytes"
	"encoding/json"
	"example1/app/http/controller"
	"example1/app/model"
	"example1/app/model/responses"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ServiceMock struct {
	mock.Mock
}

// Mock因為傳入的interface需要指定所有method，不然會報錯
func (serviceMock *ServiceMock) Login(requestData *model.LoginStudent) (model.Student, string) {
	args := serviceMock.Called(requestData)
	return args.Get(0).(model.Student), args.Get(1).(string)
}
func (serviceMock *ServiceMock) CreateUser(*model.Student) (int, string)    { return 0, "" }
func (serviceMock *ServiceMock) ScoreSearch(string, uint) ([]interface{}, string) { return nil, "" }
func (serviceMock *ServiceMock) GetRedisKey(string) ([]byte, error)               { return nil, nil }
func (serviceMock *ServiceMock) SetRedisKey(string, []byte) error                 { return nil }
func (serviceMock *ServiceMock) ComparePasswords(string, string) (bool, error)    { return false, nil }
func (serviceMock *ServiceMock) HashAndSalt([]byte) string                        { return "" }

func Test_userController_LoginUser(t *testing.T) {
	tests := []struct {
		name                   string
		PostBody               model.LoginStudent
		MockResponse           model.Student
		MockResponseStatus     string
		ExpectedHttpStatus     int
		ExpectedResponseStatus string
	}{
		{
			name: "Success_login",
			PostBody: model.LoginStudent{
				Name:     "Jason",
				Password: "12345678",
			},
			MockResponse: model.Student{
				Id:       97,
				Name:     "Jason",
				Password: "12345678",
			},
			MockResponseStatus:     responses.Success,
			ExpectedHttpStatus:     http.StatusOK,
			ExpectedResponseStatus: responses.Success,
		},
		{
			name:     "Failure_login_parameterErr",
			PostBody: model.LoginStudent{},
			MockResponse: model.Student{
				Id:       98,
				Name:     "Emily",
				Password: "1234",
			},
			MockResponseStatus:     responses.ParameterErr,
			ExpectedHttpStatus:     http.StatusNotFound,
			ExpectedResponseStatus: responses.ParameterErr,
		},
		{
			name: "Failure_login_StatusNotFound",
			PostBody: model.LoginStudent{
				Name:     "Andy",
				Password: "123",
			},
			MockResponse: model.Student{
				Id:       99,
				Name:     "Andy",
				Password: "123",
			},
			MockResponseStatus:     responses.PasswordErr,
			ExpectedHttpStatus:     http.StatusNotFound,
			ExpectedResponseStatus: responses.PasswordErr,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 設定 gin 的測試模式
			gin.SetMode(gin.TestMode)

			serviceMock := new(ServiceMock)
			// 設定 ServiceMock 的 Login 方法的預期輸出

			serviceMock.On("Login", mock.Anything).
				Return(tt.MockResponse, tt.MockResponseStatus)

			// 使用工廠模式時的方法
			// controller := controller.NewUserController(serviceMock)
			controller := controller.UserController{
				UserService: serviceMock,
			}

			// 設置 POST 請求的 Body
			jsonValue, _ := json.Marshal(tt.PostBody)

			// create a request to test the controller
			req, _ := http.NewRequest("POST", "/user/api/login", bytes.NewBuffer(jsonValue))

			// create a ResponseRecorder to record the response
			w := httptest.NewRecorder()

			// create a fake gin.Context
			_, router := gin.CreateTestContext(w)

			// 將路由註冊到 gin 的引擎上
			router.POST("/user/api/login", controller.LoginUser())

			// call the controller function
			router.ServeHTTP(w, req)

			// 關閉請求的 Body
			req.Body.Close()

			// 解析response
			var responseBody map[string]interface{}
			err := json.Unmarshal(w.Body.Bytes(), &responseBody)

			// 驗證ctx status code, 程式的status code, 與marshall完後的資料是否有錯誤
			assert := assert.New(t)
			assert.Equal(tt.ExpectedHttpStatus, w.Code)
			assert.Equal(tt.ExpectedResponseStatus, responseBody["status"])
			assert.Nil(err)
		})
	}
}

func TestUserController_LogoutUser(t *testing.T) {
	tests := []struct {
		name                   string
		authHeader             string
		isBlacklisted          bool
		h                      *controller.UserController
		expectedStatus         int
		expectedBodyMessage    string
		expectResponsesSuccess string
	}{
		{
			name:                   "success_logout",
			h:                      controller.NewUserController(),
			expectedStatus:         http.StatusOK,
			expectedBodyMessage:    "Logout Successfully.",
			expectResponsesSuccess: responses.Success,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)

			req, _ := http.NewRequest("GET", "/user/api/logout", nil)
			req.Header.Set("Authorization", tt.authHeader)

			w := httptest.NewRecorder()
			_, router := gin.CreateTestContext(w)
			router.GET("/user/api/logout", tt.h.LogoutUser())
			router.ServeHTTP(w, req)

			var response responses.Response
			err := json.Unmarshal(w.Body.Bytes(), &response)

			assert.Nil(t, err)
			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.Equal(t, tt.expectResponsesSuccess, response.Status)
			assert.Equal(t, tt.expectedBodyMessage, response.Data.(map[string]interface{})["message"])
		})
		// test
	}
}
