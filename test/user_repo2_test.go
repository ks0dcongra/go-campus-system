package repository

import (
	"example1/app/model"
	"example1/app/repository"
	"example1/database"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Test_UserRepository_CheckUserPassword2(t *testing.T) {

	tests := []struct {
		Name     string
		Password string
	}{
		{
			Name:     "James",
			Password: "1234567890",
		},
		{
			Name:     "Curry",
			Password: "12345678",
		},
		
	}

	student := []model.Student{
		
		{
			Id:       2,
			Name:     "James",
			Password: "$2a$04$fn7SQX1dw4TFNlaEXBZZiuZDD2.b6TY4aYuhd2eCrbkwdrnpxMTmS", // Example hashed password
		},
		{
			Id:       1,
			Name:     "Curry",
			Password: "$2a$04$fn7SQX1dw4TFNlaEXBZZiuZDD2.b6TY4aYuhd2eCrbkwdrnpxMTmS", // Example hashed password
		},
	}

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


	database.DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:        "sqlmock_db_0",
		DriverName: "postgres",
		Conn:       db,
	}), &gorm.Config{})

	if err != nil {
		t.Fatalf("Failed to open mock DB: %v", err)
	}

	for i, tt := range tests {	
		
		tt := tt
		i := i
		t.Run(tt.Name, func(t *testing.T) {
			log.Println(i)
			log.Println(tt)
			name := tt.Name
			password := tt.Password
			t.Parallel()
			// 设置预期的查询和结果
			
			mock.ExpectQuery(`SELECT \* FROM "students" WHERE name = \$1 ORDER BY "students"\."id" LIMIT 1`).
				WithArgs(tt.Name).
				WillReturnRows(sqlmock.NewRows([]string{"id", "name", "password"}).
					AddRow(student[i].Id, student[i].Name, student[i].Password))
			// Create a new instance of _UserRepository
			userRepo := repository.UserRepository()
			// Call the CheckUserPassword function
			condition := &model.LoginStudent{
				Name:     name,
				Password: password,
			}
			gotStudent, gotResult, gotTokenResult := userRepo.CheckUserPassword(condition)

			log.Println("student:", gotStudent)
			log.Println("result:", gotResult)
			log.Println("tokenResult:", gotTokenResult)
			// Assert the token result
			if gotTokenResult == "密碼錯誤" {
				t.Errorf("Expected non-empty token result, got empty string: %v,%v,%v", gotStudent, gotResult, gotTokenResult)
			}		
			t.Cleanup(func() {
				// Assert that all expectations were met
				err = mock.ExpectationsWereMet()
				if err != nil {
					t.Errorf("Failed to meet DB expectations: %v", err)
				}
			})
		})	
	}
}
