package service

import (
	"example1/app/model"
	"example1/app/model/responses"
	"example1/app/repository"
	"log"
	"golang.org/x/crypto/bcrypt"
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
	pwd := []byte(data.Password)
	hash := hashAndSalt(pwd)
	data.Password = hash
	student_id, db := repository.UserRepository().Create(data)
	if db.Error!= nil {
		return -1 , responses.Error
	}
	return student_id,responses.Success
}

// scoreSearch
func (h *UserService) ScoreSearch(data string) (student model.ReturnStudent , status string) {
	student, db := repository.UserRepository().ScoreSearch(data)
	if db.Error!= nil {
		return student , responses.Error
	}
	return student ,responses.Success
}

func hashAndSalt(pwd []byte) string {
	// Use GenerateFromPassword to hash & salt pwd.
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost. 
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}
