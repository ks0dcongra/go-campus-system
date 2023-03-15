package service

import (
	"example1/app/model"
	"example1/app/model/responses"
	"example1/app/repository"
	"fmt"
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
	// fmt.Printf("Datatype of i2 : %T\n", data.Password)
	// pwd := getPwd(data.Password)
	// var pwd string = data.Password
	pwd := []byte(data.Password)
	hash := hashAndSalt(pwd)
	fmt.Println("hash solution",hash)
	data.Password = hash
	fmt.Println("hash solution2",data)
	fmt.Println("hash solution3",data)
	// fmt.Printf("Datatype of i : %T\n", pwd)
	// fmt.Println("Salted Hashoop", data)
	// fmt.Println("Salted Hashoop2", data.Password)
	student_id, db := repository.UserRepository().Create(data)
	if db.Error!= nil {
		return -1 , responses.Error
	}
	return student_id,responses.Success
}

// func getPwd(password string) []byte {
// 	// Prompt the user to enter a password
// 	fmt.Println("Enter a password")
// 	// Variable to store the users input
// 	var pwd string = password
// 	// Read the users input
// 	// _, err := fmt.printf(&pwd)
// 	// if err != nil {
// 	// 	log.Println(err)
// 	// }
// 	// Return the users input as a byte slice which will save us
// 	// from having to do this conversion later on
// 	return []byte(pwd)
// }


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