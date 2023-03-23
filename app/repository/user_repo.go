package repository

import (
	"example1/app/model"
	"example1/database"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type _UserRepository struct {
}

// var new []model.ReturnStudent

func UserRepository() *_UserRepository {
	return &_UserRepository{}
}

// Login Check
func (h *_UserRepository) CheckUserPassword(condition *model.LoginStudent) (Student model.Student, result *gorm.DB) {
	name := condition.Name
	student := model.Student{}
	result = database.DB.Where("name = ?", name).First(&student)
	pwdMatch, err := comparePasswords(student.Password, condition.Password)
	if !pwdMatch {
		result.Error = err
		return student, result
	}
	return student, result
}

// Create User
func (h *_UserRepository) Create(data *model.CreateStudent) (id int, result *gorm.DB) {
	student := model.Student{
		Name:           data.Name,
		Password:       data.Password,
		Student_number: data.Student_number,
		CreatedTime:    time.Now(),
		UpdatedTime:    time.Now()}
	result = database.DB.Create(&student)
	return student.Id, result
}

// score search
func (h *_UserRepository) ScoreSearch(data string) (Student []model.ReturnStudent, result model.ReturnStudent) {
	studentSearch := model.ReturnStudent{}
	student := model.Student{}
	rows, _ := database.DB.Model(&student).Select("scores.score,students.name,courses.subject").
		Joins("left join scores on students.id = scores.student_id").
		Joins("left join courses on courses.id = scores.course_id").Where("students.id = ?", data).Rows()
	defer rows.Close()
	var new []model.ReturnStudent
	var result2 model.ReturnStudent
	if rows != nil {
		for rows.Next() {
			database.DB.ScanRows(rows, &studentSearch)
			result2 = model.ReturnStudent{Name: studentSearch.Name, Subject: studentSearch.Subject, Score: studentSearch.Score}
			new = append(new, result2)
		}
	}

	return new, result2
}

func comparePasswords(hashedPwd string, plainPwd string) (bool, error) {
	byteHash := []byte(hashedPwd)
	byteHash2 := []byte(plainPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, byteHash2)
	if err != nil {
		return false, err
	}
	return true, err
}
