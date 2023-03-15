package repository

import (
	"example1/app/model"
	"example1/database"
	"log"
	"time"
	"golang.org/x/crypto/bcrypt"
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
	log.Println("NOWAY:",password)
	log.Println("NOWAY2:",student)
	result = database.DB.Where("name = ?", name).First(&student)
	log.Println("NOWAY3:",student)
	pwdMatch, err := comparePasswords(student.Password,condition.Password)
	if !pwdMatch {
		result.Error = err
		return student, result
	}
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

func comparePasswords(hashedPwd string, plainPwd string) (bool, error) {
    // Since we'll be getting the hashed password from the DB it
    // will be a string so we'll need to convert it to a byte slice
    byteHash := []byte(hashedPwd)
	byteHash2 := []byte(plainPwd)
    err := bcrypt.CompareHashAndPassword(byteHash, byteHash2)
    if err != nil {
        log.Println(err)
        return false, err
    }
    return true, err
}