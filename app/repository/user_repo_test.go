package repository_test

import (
	"example1/app/model"
	"example1/app/repository"
	"example1/database"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestUserRepository_SQL_Success(t *testing.T) {
	type args struct {
		condition *model.LoginStudent
	}
	tests := []struct {
		name		string
		student     model.Student
		args        args
		h           *repository.UserRepository
		expectedStudent model.Student
		expectedErr     error
	}{
		{
			name: "Success_Case1_SQL測試",
			student: model.Student{
					Id:       1,
					Name:     "James",
					Password: "$2a$04$fn7SQX1dw4TFNlaEXBZZiuZDD2.b6TY4aYuhd2eCrbkwdrnpxMTmS", // Password為丟入之預期狀況
			},
			args:        args{condition: &model.LoginStudent{Name: "James", Password: "12345678"}},
			h:           repository.NewUserRepository(),
			expectedStudent: model.Student{Id: 1, Name: "James", Password: "$2a$04$fn7SQX1dw4TFNlaEXBZZiuZDD2.b6TY4aYuhd2eCrbkwdrnpxMTmS"},
		},
		{
			name: "Success_Case2_SQL測試",
			student: model.Student{
					Id:       2,
					Name:     "Curry",
					Password: "000$2a$04$fn7SQX1dw4TFNlaEXBZZiuZDD2.b6TY4aYuhd2eCrbkwdrnpxMTmS", // Password為丟入之預期狀況
			},
			args:        args{condition: &model.LoginStudent{Name: "Curry", Password: "12345678"}},
			h:           repository.NewUserRepository(),
			expectedStudent: model.Student{Id: 2, Name: "Curry", Password: "000$2a$04$fn7SQX1dw4TFNlaEXBZZiuZDD2.b6TY4aYuhd2eCrbkwdrnpxMTmS"},
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
		t.Run("tt.name", func(t *testing.T) {
			// 設定Mock SQL預期回傳資料與欄位
			rows := sqlmock.NewRows([]string{"id", "name", "password"}).
				AddRow(tt.student.Id, tt.student.Name, tt.student.Password)

			// 設定Mock SQL撈取資料後預期之狀況
			mock.ExpectQuery(`SELECT \* FROM "students"`).
				WithArgs(tt.student.Name).
				WillReturnRows(rows)

			gotStudent, err := tt.h.Login(tt.args.condition)

			assert := assert.New(t)
			assert.Nil(err)
			assert.Equal(tt.expectedStudent, gotStudent)

			sqlerr := mock.ExpectationsWereMet()
			if !assert.Nil(sqlerr) {
				t.Errorf("Failed to meet DB expectations: %v", err)
			}
		})
	}
}

func TestUserRepository_SQL_Failure(t *testing.T) {
	type args struct {
		condition *model.LoginStudent
	}
	tests := []struct {
		name		string
		student     model.Student
		args        args
		h           *repository.UserRepository
		expectedStudent model.Student
	}{
		{
			name:	"Failure_Case1_SQL測試",
			student: model.Student{
					Id:       1,
					Name:     "James",
					Password: "$2a$04$fn7SQX1dw4TFNlaEXBZZiuZDD2.b6TY4aYuhd2eCrbkwdrnpxMTmS", // Password為丟入之預期狀況
			},
			args:        args{condition: &model.LoginStudent{ Password: "12345678"}},
			h:           repository.NewUserRepository(),
			expectedStudent: model.Student{Id: 0, Name: "", Password: ""},
		},
		{
			name:	"Failure_Case2_SQL測試",
			student: model.Student{
					Id:       2,
					Name:     "Curry",
					Password: "000$2a$04$fn7SQX1dw4TFNlaEXBZZiuZDD2.b6TY4aYuhd2eCrbkwdrnpxMTmS", // Password為丟入之預期狀況
			},
			args:        args{condition: &model.LoginStudent{}},
			h:           repository.NewUserRepository(),
			expectedStudent: model.Student{Id: 0, Name: "", Password: ""},
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
				AddRow(tt.student.Id, tt.student.Name, tt.student.Password)

			// 設定Mock SQL撈取資料後預期之狀況
			mock.ExpectQuery(`SELECT \* FROM "students"`).
				WithArgs(tt.student.Name).
				WillReturnRows(rows)

			gotStudent, err := tt.h.Login(tt.args.condition)

			assert := assert.New(t)
			assert.NotNil(err)
			assert.Equal(tt.expectedStudent, gotStudent)

			sqlerr := mock.ExpectationsWereMet()
			if !assert.NotNil(sqlerr) {
				t.Errorf("Failed to meet DB expectations: %v", err)
			}
		})
	}
}