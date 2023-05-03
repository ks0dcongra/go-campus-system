package service_test

import (
	"example1/app/model"
	"example1/app/service"
	"example1/utils/token"
	"github.com/stretchr/testify/assert"
	"testing"
)

type MockUserService struct {
	UserServiceHashToken service.UserServiceHashTokenInterface
	JwtFactory           token.TokenInterface
}

func TestUserService_Login_Success(t *testing.T) {
	type args struct {
		condition *model.LoginStudent
	}
	tests := []struct {
		student []model.Student
		args    args
		h       *service.UserService
		// wantPwdErr  error
		// wantTokenErr error
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
				UserServiceHashToken: service.NewUserServiceHashToken(),
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

func TestUserService_Login_Failure(t *testing.T) {
	type args struct {
		condition *model.LoginStudent
	}
	tests := []struct {
		student []model.Student
		args    args
		h       *service.UserService
		// wantPwdErr  error
		// wantTokenErr error
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
				UserServiceHashToken: service.NewUserServiceHashToken(),
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

// 測試貓咪interface
// type MockCatService struct {
// 	service.UserServiceHashToken
// }
// func AnimalTestInterface(userServiceHashTokenInterface service.UserServiceHashTokenInterface)  (bool, error){
//     a,b := userServiceHashTokenInterface.ComparePasswords("a","b")
// 	return a,b
// }
// func Cat(){
// 	cat := &MockCatService{}
// 	a,b := AnimalTestInterface(cat)
// 	log.Println(a,b)
// }
