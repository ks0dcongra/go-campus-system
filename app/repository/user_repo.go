package repository

import (
	"example1/app/model"
	"example1/database"
	"log"
	"time"

	"gorm.io/gorm"
)

type _UserRepository struct {
}

func UserRepository() *_UserRepository {
	return &_UserRepository{}
}


// Login Check
func (h *_UserRepository) CheckUserPassword(condition *model.LoginStudent) (Student model.Student, result *gorm.DB) {
	name := condition.Name
	password := condition.Password
	student := model.Student{}
	result = database.DB.Where("name = ? and password = ?", name, password).First(&student)
	return student, result
}

// Create User
func (h *_UserRepository) Create(data *model.CreateStudent) (id int, result *gorm.DB) {
	log.Print("happy4",data)
	student := model.Student{
		Name: data.Name,
		Password: data.Password,
		Student_number: data.Student_number,
		CreatedTime: time.Now(),
		UpdatedTime: time.Now()}
	log.Print("happy5",student)
	log.Print("happy6",&student)
	result = database.DB.Create(&student)
	return student.Id, result
}