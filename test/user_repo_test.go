package test

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

func Test_UserRepository_CheckUserPassword(t *testing.T) {
	// 创建一个内存数据库连接
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
	defer db.Close()

	database.DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
        DriverName:           "postgres",
		Conn: db,
	}), &gorm.Config{})

	if err != nil {
		t.Fatalf("Failed to open mock DB: %v", err)
	}

	student := model.Student{
		Id:       1,
		Name:     "Curry",
		Password: "$2a$04$fn7SQX1dw4TFNlaEXBZZiuZDD2.b6TY4aYuhd2eCrbkwdrnpxMTmS", // Example hashed password
	}

	condition := &model.LoginStudent{
        Name:     "Curry", 
		Password: "12345678",
    }

			// Create a new instance of _UserRepository
	userRepo := repository.UserRepository()
	
	// 设置预期的查询和结果
	mock.ExpectQuery(`SELECT \* FROM "students" WHERE name = \$1 ORDER BY "students"\."id" LIMIT 1`).
	WithArgs(condition.Name).
	WillReturnRows(sqlmock.NewRows([]string{"id", "name", "password"}).AddRow(student.Id, student.Name, student.Password))
	
	
	// Call the CheckUserPassword function
	student, result, tokenResult := userRepo.CheckUserPassword(condition)
	log.Println("student:",student)
	log.Println("result:",result)
	log.Println("tokenResult:",tokenResult)
	// Assert the token result
	if tokenResult == "" {
		t.Errorf("Expected non-empty token result, got empty string: %v,%v,%v",student,result,tokenResult)
	}

	// Assert that all expectations were met
	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Failed to meet DB expectations: %v", err)
	}
}
