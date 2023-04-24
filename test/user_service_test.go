package test

import (
	"example1/app/model"
	"example1/app/service"

	"fmt"
	"log"
	"reflect"
	"testing"
	"example1/database"
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// type MockUserService struct{
// 	UserService service.UserServiceInterface
// }

func (m *MockUserService) CheckUserPassword() error {
	condition := &model.LoginStudent{Name: "James", Password: "12345678"}
	log.Println("tokenResult")

	// 開始創一個模擬database實體
	db, mock, err := sqlmock.New()
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	if db == nil {
		return fmt.Errorf("db is null") 
	} 
	
	if mock == nil {
		return fmt.Errorf("mock is null")
	}
	// 取代掉我原本的資料庫database.DB
	database.DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:        "sqlmock_db_0",
		DriverName: "postgres",
		Conn:       db,
	}), &gorm.Config{})

	// 如果打不開Mock資料庫則報錯
	if err != nil {
		fmt.Errorf("Failed to open mock DB: %v", err)
	}

	// 設定Mock SQL預期回傳資料與欄位
	rows := sqlmock.NewRows([]string{"id", "name", "password"}).
		AddRow("1", "James", "$2a$04$fn7SQX1dw4TFNlaEXBZZiuZDD2.b6TY4aYuhd2eCrbkwdrnpxMTmS")

	// 設定Mock SQL撈取資料後預期之狀況
	mock.ExpectQuery(`SELECT \* FROM "students"`).
		WithArgs("James").
		WillReturnRows(rows)

	_, _, tokenResult := m.UserService.Login(condition)
	
	log.Println(tokenResult)
	return fmt.Errorf("Failed to open mock DB")
}

func TestUserService_Login(t *testing.T){
	type args struct {
		condition *model.LoginStudent
	}
	tests := []struct {
		name            string
		h               *service.UserService
		args            args
		wantStudent     model.Student
		wantStatus      string
		wantTokenResult string
	}{
		// TODO: Add test cases.
	}
	mockUserService := &MockUserService{}

	tokenResult := mockUserService.CheckUserPassword()

	log.Println("hihi",tokenResult)
	// condition := &model.LoginStudent{Name: "James", Password: "12345678"}
	// student, status, tokenResult := m.UserService.Login(condition)
	
	// func (m *MockUserService ) CheckUserPassword() string {
	// 	condition := &model.LoginStudent{Name: "James", Password: "12345678"}
	// 	_, _, tokenResult := m.UserService.Login(condition)
	// 	return tokenResult
	// }
	t.Errorf("UserService.Login() gotStudent = %v, want %v",tokenResult,tokenResult)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStudent, gotStatus, gotTokenResult := tt.h.Login(tt.args.condition)
			if !reflect.DeepEqual(gotStudent, tt.wantStudent) {
				t.Errorf("UserService.Login() gotStudent = %v, want %v", gotStudent, tt.wantStudent)
			}
			if gotStatus != tt.wantStatus {
				t.Errorf("UserService.Login() gotStatus = %v, want %v", gotStatus, tt.wantStatus)
			}
			if gotTokenResult != tt.wantTokenResult {
				t.Errorf("UserService.Login() gotTokenResult = %v, want %v", gotTokenResult, tt.wantTokenResult)
			}
		})
	}
}

// Service: 判斷condition