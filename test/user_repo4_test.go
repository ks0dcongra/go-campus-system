package test

import (
	"example1/app/model"
	"example1/app/repository"
	"example1/database"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestUserRepository_SQL_Success(t *testing.T) {
	type args struct {
		condition *model.LoginStudent
	}
	tests := []struct {
		student  []model.Student
		args            args
		h               *repository.UserRepository
		wantStudent     model.Student
		wantErr         error
	}{
		{
			student: []model.Student{
				{
					Id:       2,
					Name:     "James",
					Password: "$2a$04$fn7SQX1dw4TFNlaEXBZZiuZDD2.b6TY4aYuhd2eCrbkwdrnpxMTmS", // Password為丟入之預期狀況
				},
			},
			args:   args{condition: &model.LoginStudent{Name: "James"}},
			h:  repository.NewUserRepository(),
			wantStudent:     model.Student{Id: 2, Name: "James",Password: "$2a$04$fn7SQX1dw4TFNlaEXBZZiuZDD2.b6TY4aYuhd2eCrbkwdrnpxMTmS"},
			wantErr:  nil,
		},
		{
			student: []model.Student{
				{
					Id:       1,
					Name:     "Curry",
					Password: "$2a$04$fn7SQX1dw4TFNlaEXBZZiuZDD2.b6TY4aYuhd2eCrbkwdrnpxMTmS", // Password為丟入之預期狀況
				},
			},
			args:   args{condition: &model.LoginStudent{Name: "Curry"}},
			h:  repository.NewUserRepository(),
			wantStudent:     model.Student{Id: 1, Name: "Curry",Password: "$2a$04$fn7SQX1dw4TFNlaEXBZZiuZDD2.b6TY4aYuhd2eCrbkwdrnpxMTmS"},
			wantErr: nil,
		},
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

	for _, tt := range tests {
		t.Run("測試SQL", func(t *testing.T) {
			// 設定Mock SQL預期回傳資料與欄位
			rows := sqlmock.NewRows([]string{"id", "name", "password"}).
				AddRow(tt.student[0].Id, tt.student[0].Name, tt.student[0].Password)
				
			// 設定Mock SQL撈取資料後預期之狀況
			mock.ExpectQuery(`SELECT \* FROM "students"`).
				WithArgs(tt.student[0].Name).
				WillReturnRows(rows)
			
			gotStudent, err := tt.h.Login(tt.args.condition)
			if err != tt.wantErr {
				t.Errorf("UserRepository.CheckUserPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotStudent, tt.wantStudent) {
				t.Errorf("UserRepository.CheckUserPassword() = %v, want %v", gotStudent, tt.wantStudent)
			}
		})
	}
}