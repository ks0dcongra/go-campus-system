package repository

import (
	"example1/app/model"
	"example1/database"
	"gorm.io/gorm"
)

type _UserRepository struct {
}

func UserRepository() *_UserRepository {
	return &_UserRepository{}
}


// 確認使用者登入密碼正確
func (h *_UserRepository) CheckUserPassword(condition *model.LoginStudent) (Student model.Student, result *gorm.DB) {
	name := condition.Name
	password := condition.Password
	student := model.Student{}
	result = database.DB.Where("name = ? and password = ?", name, password).First(&student)
	return student, result
}
