package controller_test

import (
	"bytes"
	"encoding/json"
	"example1/app/http/controller"
	"example1/app/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ServiceMock struct {
    mock.Mock
}

func (serviceMock *ServiceMock) Login(requestData *model.LoginStudent) (model.Student, string){
    args := serviceMock.Called(requestData)
    return args.Get(0).(model.Student),args.Get(1).(string)
}

func Test_userController_LoginUser2(t *testing.T) {
	
	tests := []struct {
		name string
		PostBody model.LoginStudent
		MockResponse model.Student
		MockResponseStatus string
		wantHttpStatusOK int
		wantResponseStatus string
	}{
		{
			name : "test_case_1",
			PostBody : model.LoginStudent{
				Name:"Mar234",
				Password:"12345678",
			},
			MockResponse : model.Student{
				Id: 99, 
				Name: "Jack", 
				Password: "12345678",
			},
			MockResponseStatus : "0",
			wantHttpStatusOK : http.StatusOK,
			wantResponseStatus : "0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 設定 gin 的測試模式
			gin.SetMode(gin.TestMode)

			serviceMock := new(ServiceMock)
			// 設定 ServiceMock 的 Login 方法的預期輸出
			
			serviceMock.On("Login", mock.Anything).
				Return(tt.MockResponse,tt.MockResponseStatus)
				
			controller := controller.NewUserController(serviceMock)
			// 使用非工廠模式時的方法
			// controller := controller.UserController{
			//     UserService: serviceMock,
			// }

			// 設置 POST 請求的 Body
			jsonValue, _ := json.Marshal(tt.PostBody)

			// create a request to test the controller
			req,_ := http.NewRequest("POST", "/user/api/login", bytes.NewBuffer(jsonValue))

			// create a ResponseRecorder to record the response
			w := httptest.NewRecorder()
			
			// create a fake gin.Context
			c, router := gin.CreateTestContext(w)
			c.Request = req

			// 將路由註冊到 gin 的引擎上
			router.POST("/user/api/login", controller.LoginUser())

			// call the controller function
			router.ServeHTTP(w, c.Request)

			// 關閉請求的 Body
			req.Body.Close()

			// 解析response
			var responseBody map[string]interface{}
			err := json.Unmarshal(w.Body.Bytes(),&responseBody)
			
			// 驗證ctx status code, 程式的status code, 與marshall完後的資料是否有錯誤
			assert := assert.New(t)
			assert.Equal(tt.wantHttpStatusOK, w.Code)
			assert.Equal(tt.wantResponseStatus, responseBody["status"])
			assert.Nil(err)
		})
	}
}
