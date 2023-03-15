package service

import (
	"example1/app/model"
	"example1/app/model/responses"
	"example1/app/repository"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

// Login
func (h *UserService) Login(condition *model.LoginStudent) (student model.Student, status string) {
	student, db := repository.UserRepository().CheckUserPassword(condition)
	if db.Error!= nil {
		return student , responses.Error
	}
	return student , responses.Success
}

// CreateUser
func (h *UserService) CreateUser(data *model.CreateStudent) (student_id int , status string) {
	student_id, db := repository.UserRepository().Create(data)
	if db.Error!= nil {
		return -1 , responses.Error
	}
	return student_id,responses.Success
}
