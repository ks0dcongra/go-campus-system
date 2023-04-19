package test

import (
	"example1/app/model"
	"example1/app/repository"
	"example1/database"
	"log"
	"testing"
	"reflect"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Test_UserRepository_CheckUserPassword3(t *testing.T) {
	type args struct {
		condition *model.LoginStudent
	}
	tests := []struct {
		Name     string
		Password string
		student  []model.Student
		h               *repository.Export_UserRepository
		args            args
		wantStudent     model.Student
		wantResult      *gorm.DB
		wantTokenResult string
	}{
		{
			Name:     "James",
			Password: "12345678",
			student: []model.Student{
				{
					Id:       2,
					Name:     "James",
					Password: "$2a$04$fn7SQX1dw4TFNlaEXBZZiuZDD2.b6TY4aYuhd2eCrbkwdrnpxMTmS", // Example hashed password
				},
			},
			h:  repository.UserRepository(),
			args:   args{condition: &model.LoginStudent{Name: "James", Password: "12345678"}},
			wantStudent:     model.Student{Id: 1, Name: "test"},
			wantResult:      nil,
			wantTokenResult: "token",
		},
		{
			Name:     "Curry",
			Password: "12345678",
			student: []model.Student{
				{
					Id:       1,
					Name:     "Curry",
					Password: "$2a$04$fn7SQX1dw4TFNlaEXBZZiuZDD2.b6TY4aYuhd2eCrbkwdrnpxMTmS", // Example hashed password
				},
			},
			h:  repository.UserRepository(),
			args:   args{condition: &model.LoginStudent{Name: "Curry", Password: "12345678"}},
			wantStudent:     model.Student{Id: 1, Name: "test"},
			wantResult:      nil,
			wantTokenResult: "token",
		},
		// GPT示範案例
		// {
		// 	Name:            "test case 1",
		// 	Password:        "password",
		// 	student:         []model.Student{{ID: 1, Name: "test"}},
		// 	h:               repository.UserRepository(),
		// 	args:            args{condition: &model.LoginStudent{Username: "test", Password: "password"}},
		// 	wantStudent:     model.Student{ID: 1, Name: "test"},
		// 	wantResult:      nil,
		// 	wantTokenResult: "token",
		// },
	}

	// 開始創一個模擬database實體
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("Failed to open mock sql db, got error: %v", err)
	}
	if db == nil {
		t.Error("db is null")
	}
	if mock == nil {
		t.Error("mock is null")
	}

	// 取代掉我原本的資料庫database.DB
	database.DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:        "sqlmock_db_0",
		DriverName: "postgres",
		Conn:       db,
	}), &gorm.Config{})

	// 如果打不開Mock資料庫則報錯
	if err != nil {
		t.Fatalf("Failed to open mock DB: %v", err)
	}

	for i, tt := range tests {
		// tt := tt
		// i := i
		t.Run(tt.Name, func(t *testing.T) {
			log.Println(i)
			log.Println(tt)
			mock.ExpectQuery(`SELECT \* FROM "students" WHERE name = \$1 ORDER BY "students"\."id" LIMIT 1`).
				WithArgs(tt.Name).
				WillReturnRows(sqlmock.NewRows([]string{"id", "name", "password"}).
				AddRow(tt.student[i].Id, tt.student[i].Name, tt.student[i].Password))
				
			gotStudent, gotResult, gotTokenResult := tt.h.CheckUserPassword(tt.args.condition)
			if !reflect.DeepEqual(gotStudent, tt.wantStudent) {
				t.Errorf("_UserRepository.CheckUserPassword() gotStudent = %v, want %v", gotStudent, tt.wantStudent)
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("_UserRepository.CheckUserPassword() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
			if gotTokenResult != tt.wantTokenResult {
				t.Errorf("_UserRepository.CheckUserPassword() gotTokenResult = %v, want %v", gotTokenResult, tt.wantTokenResult)
			}
		})
	}
}
