package service_test

import (
	"errors"
	"example1/app/model"
	"example1/app/model/responses"
	"example1/app/service"
	"example1/utils/token"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	UserServiceHashToken service.UserServiceInterface
	JwtFactory           token.TokenInterface
}

func TestUserService_HashAndTokenFunction_Success(t *testing.T) {
	type args struct {
		condition *model.LoginStudent
	}
	tests := []struct {
		testName1 string
		testName2 string
		student   []model.Student
		args      args
		h         *service.UserService
	}{
		{
			testName1: "測試Hash: Test Case1",
			testName2: "測試Token: Test Case1",
			student: []model.Student{
				{
					Id:       2,
					Name:     "James",
					Password: "$2a$04$fn7SQX1dw4TFNlaEXBZZiuZDD2.b6TY4aYuhd2eCrbkwdrnpxMTmS", // Password為丟入之預期狀況
				},
			},
			args: args{condition: &model.LoginStudent{Name: "James", Password: "12345678"}},
			h:    service.NewUserService(),
		},
		{
			testName1: "測試Hash: Test Case2",
			testName2: "測試Token: Test Case2",
			student: []model.Student{
				{
					Id:       1,
					Name:     "Curry",
					Password: "$2a$04$fn7SQX1dw4TFNlaEXBZZiuZDD2.b6TY4aYuhd2eCrbkwdrnpxMTmS", // Password為丟入之預期狀況
				},
			},
			args: args{condition: &model.LoginStudent{Name: "Curry", Password: "12345678"}},
			h:    service.NewUserService(),
		},
	}

	for _, tt := range tests {
		t.Run("tt.testName1", func(t *testing.T) {
			mockUserService := &MockUserService{
				UserServiceHashToken: service.NewUserService(),
			}
			_, pwdErr := mockUserService.UserServiceHashToken.ComparePasswords(tt.student[0].Password, tt.args.condition.Password)
			assert := assert.New(t)
			assert.Nil(pwdErr)
		})
		t.Run("tt.testName2", func(t *testing.T) {
			mockUserService := &MockUserService{
				JwtFactory: token.Newjwt(),
			}
			_, tokenErr := mockUserService.JwtFactory.GenerateToken(tt.student[0].Id)
			assert := assert.New(t)
			assert.Nil(tokenErr)
		})
	}
}

func TestUserService_HashAndTokenFunction_Failure(t *testing.T) {
	type args struct {
		condition *model.LoginStudent
	}
	tests := []struct {
		testName1 string
		testName2 string
		student   []model.Student
		args      args
		h         *service.UserService
	}{
		{
			testName1: "測試Hash: Test Case1",
			testName2: "測試Token: Test Case1",
			student: []model.Student{
				{
					Id:       0,
					Name:     "James",
					Password: "$0002a$04$fn7SQX1dw4TFNlaEXBZZiuZDD2.b6TY4aYuhd2eCrbkwdrnpxMTmS", // Password為丟入之預期狀況
				},
			},
			args: args{condition: &model.LoginStudent{Name: "James", Password: "123456789"}},
			h:    service.NewUserService(),
		},
		{
			testName1: "測試Hash: Test Case2",
			testName2: "測試Token: Test Case2",
			student: []model.Student{
				{
					Id:       0,
					Name:     "Curry",
					Password: "$1112a$04$fn7SQX1dw4TFNlaEXBZZiuZDD2.b6TY4aYuhd2eCrbkwdrnpxMTmS", // Password為丟入之預期狀況
				},
			},
			args: args{condition: &model.LoginStudent{Name: "Curry", Password: "012345678"}},
			h:    service.NewUserService(),
		},
	}

	for _, tt := range tests {
		t.Run("tt.testName1", func(t *testing.T) {
			mockUserService := &MockUserService{
				UserServiceHashToken: service.NewUserService(),
			}
			_, pwdErr := mockUserService.UserServiceHashToken.ComparePasswords(tt.student[0].Password, tt.args.condition.Password)
			assert := assert.New(t)
			assert.NotNil(pwdErr)
		})
		t.Run("tt.testName2", func(t *testing.T) {
			mockUserService := &MockUserService{
				JwtFactory: token.Newjwt(),
			}
			_, tokenErr := mockUserService.JwtFactory.GenerateToken(tt.student[0].Id)
			assert := assert.New(t)
			assert.NotNil(tokenErr)
		})
	}
}

type UserRepositoryMock struct {
	mock.Mock
}

func (m *UserRepositoryMock) Login(condition *model.LoginStudent) (model.Student, error) {
	args := m.Called(condition)
	student, ok := args.Get(0).(model.Student)
	if !ok {
		return model.Student{}, errors.New("Invalid return value")
	}
	return student, args.Error(1)
}

func TestUserService_Login(t *testing.T) {
	condition := &model.LoginStudent{
		Name:     "test@example.com",
		Password: "12345678",
	}
	mockStudent := model.Student{
		Id:             1,
		Name:           "test@example.com",
		Score:          nil,
		Student_number: "testStudent_number",
		Token:          "testToken",
		Password:       "$2a$04$fn7SQX1dw4TFNlaEXBZZiuZDD2.b6TY4aYuhd2eCrbkwdrnpxMTmS", // hashed "password"
		CreatedTime:    time.Now(),
		UpdatedTime:    time.Now(),
	}

	tests := []struct {
		name            string
		DBError         error
		expectedStatus  string
		expectedStudent model.Student
	}{
		{
			name:            "returns student and success when login is successful",
			DBError:         nil,
			expectedStatus:  responses.Success,
			expectedStudent: mockStudent,
		},
		{
			name:            "returns DbErr when repository returns an error",
			DBError:         fmt.Errorf("SQL Error"),
			expectedStatus:  responses.DbErr,
			expectedStudent: model.Student{},
		},
		{
			name:            "returns TokenError when GenerateToken error",
			DBError:         nil,
			expectedStatus:  responses.TokenErr,
			expectedStudent: model.Student{},
		},
		{
			name:            "returns PasswordErr when password is incorrect",
			DBError:         nil,
			expectedStatus:  responses.PasswordErr,
			expectedStudent: model.Student{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repoMock := new(UserRepositoryMock)
			if tt.name == "returns TokenError when GenerateToken error" {
				mockStudent.Id = 0
			}
			repoMock.On("Login", condition).Return(mockStudent, tt.DBError)

			userService := service.UserService{
				UserRepository: repoMock,
			}
			if tt.name == "returns PasswordErr when password is incorrect" {
				condition.Password = "wrong_password"
			}
			actual, status := userService.Login(condition)

			assert.Equal(t, tt.expectedStatus, status)
			assert.Equal(t, tt.expectedStudent.Student_number, actual.Student_number)
		})
	}
}
