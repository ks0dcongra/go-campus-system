package service_test

import (
	"errors"
	"example1/app/model"
	"example1/app/model/responses"
	"example1/app/service"
	"example1/utils/token"
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
		student []model.Student
		args    args
		h       *service.UserService
	}{
		{
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
		t.Run("測試Hash", func(t *testing.T) {
			mockUserService := &MockUserService{
				UserServiceHashToken: service.NewUserService(),
			}
			_, pwdErr := mockUserService.UserServiceHashToken.ComparePasswords(tt.student[0].Password, tt.args.condition.Password)
			assert := assert.New(t)
			assert.Nil(pwdErr)
		})
		t.Run("測試Token", func(t *testing.T) {
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
		student []model.Student
		args    args
		h       *service.UserService
	}{
		{
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
		t.Run("測試Hash", func(t *testing.T) {
			mockUserService := &MockUserService{
				UserServiceHashToken: service.NewUserService(),
			}
			_, pwdErr := mockUserService.UserServiceHashToken.ComparePasswords(tt.student[0].Password, tt.args.condition.Password)
			assert := assert.New(t)
			assert.NotNil(pwdErr)
		})
		t.Run("測試Token", func(t *testing.T) {
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
		Name:    "test@example.com",
		Password: "12345678",
	}
	mockStudent := model.Student{
		Id:	1,
		Name: "test@example.com",
		Score: nil,
		Student_number: "",	
		Token: "",
		Password:  "$2a$04$fn7SQX1dw4TFNlaEXBZZiuZDD2.b6TY4aYuhd2eCrbkwdrnpxMTmS", // hashed "password"
		CreatedTime: time.Now(),
		UpdatedTime: time.Now(),
	}

	t.Run("returns student and success when login is successful", func(t *testing.T) {
		repoMock := new(UserRepositoryMock)
		repoMock.On("Login", condition).Return(mockStudent, nil)

		userService := service.UserService{
			UserRepository: repoMock,
		}
		actual, status := userService.Login(condition)

		expectedStatus := responses.Success
		expectedStudent := mockStudent

		assert.Equal(t, expectedStatus, status)
		assert.Equal(t, expectedStudent.Id, actual.Id)
		assert.Equal(t, expectedStudent.Name, actual.Name)
		assert.Equal(t, expectedStudent.Password, actual.Password)
	})

	t.Run("returns DbErr when repository returns an error", func(t *testing.T) {
		repoMock := new(UserRepositoryMock)
		repoMock.On("Login", condition).Return(model.Student{}, errors.New("Database error"))

		userService := service.UserService{
			UserRepository: repoMock,
		}
		actual, status := userService.Login(condition)

		expectedStatus := responses.DbErr

		assert.Equal(t, expectedStatus, status)
		assert.Equal(t, model.Student{}, actual)
	})

	t.Run("returns PasswordErr when password is incorrect", func(t *testing.T) {
		repoMock := new(UserRepositoryMock)
		// mockStudent.Password = "wrong password"
		repoMock.On("Login", condition).Return(mockStudent, nil)

		userService := service.UserService{
			UserRepository: repoMock,
		}

		condition.Password = "wrongpassword"
		_, status := userService.Login(condition)

		expectedStatus := responses.PasswordErr

		assert.Equal(t, expectedStatus, status)
	})
}
