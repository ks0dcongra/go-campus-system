package test

import (
	"example1/app/model"
	"example1/app/repository"
	"example1/database"
	"testing"
	"reflect"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
// TODO:dd
func Test_UserRepository_CheckUserPassword3(t *testing.T) {
	type args struct {
		condition *model.LoginStudent
	}
	tests := []struct {
		student  []model.Student
		args            args
		h               *repository.Export_UserRepository
		wantStudent     model.Student
		wantResult      *gorm.DB
		wantTokenResult string
	}{
		{
			student: []model.Student{
				{
					Id:       2,
					Name:     "James",
					Password: "$2a$04$fn7SQX1dw4TFNlaEXBZZiuZDD2.b6TY4aYuhd2eCrbkwdrnpxMTmS", // Password為丟入之預期狀況
				},
			},
			args:   args{condition: &model.LoginStudent{Name: "James", Password: "12345678"}},
			h:  repository.UserRepository(),
			wantStudent:     model.Student{Id: 2, Name: "James",Password: "$2a$04$fn7SQX1dw4TFNlaEXBZZiuZDD2.b6TY4aYuhd2eCrbkwdrnpxMTmS"},
			wantTokenResult: "Password Wrong!",
		},
		{
			student: []model.Student{
				{
					Id:       1,
					Name:     "Curry",
					Password: "$2a$04$fn7SQX1dw4TFNlaEXBZZiuZDD2.b6TY4aYuhd2eCrbkwdrnpxMTmS", // Password為丟入之預期狀況
				},
			},
			args:   args{condition: &model.LoginStudent{Name: "Curry", Password: "12345678"}},
			h:  repository.UserRepository(),
			wantStudent:     model.Student{Id: 1, Name: "Curry",Password: "$2a$04$fn7SQX1dw4TFNlaEXBZZiuZDD2.b6TY4aYuhd2eCrbkwdrnpxMTmS"},
			wantTokenResult: "Password Wrong!",
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
		t.Run("Check Repository SQL", func(t *testing.T) {

			// 設定Mock SQL預期回傳資料與欄位
			rows := sqlmock.NewRows([]string{"id", "name", "password"}).
				AddRow(tt.student[0].Id, tt.student[0].Name, tt.student[0].Password)

			// 設定Mock SQL撈取資料後預期之狀況
			mock.ExpectQuery(`SELECT \* FROM "students"`).
				WithArgs(tt.student[0].Name).
				WillReturnRows(rows)
				
			// 模擬東西丟入repository去跑
			gotStudent,_,gotTokenResult := tt.h.CheckUserPassword(tt.args.condition)
			// 判斷回傳的東西是否相同
			if !reflect.DeepEqual(gotStudent, tt.wantStudent) {
				t.Errorf("_UserRepository.CheckUserPassword() gotStudent = %v, want %v", gotStudent, tt.wantStudent)
			}
			// 判斷錯誤字串是否等於Password Wrong!
			if password_wrong := reflect.DeepEqual(gotTokenResult, tt.wantTokenResult); password_wrong {
				t.Errorf("_UserRepository.CheckUserPassword() gotTokenResult = %v, want %v", gotTokenResult, tt.wantTokenResult)
			}
			
			err = mock.ExpectationsWereMet()
			if err != nil {
			t.Errorf("Failed to meet DB expectations: %v", err)
			}
		})
	}
	
}
