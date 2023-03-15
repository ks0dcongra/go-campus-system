package service

import (
	"example1/app/model"
	"example1/app/model/responses"
	"example1/app/repository"
	"log"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

// Login
func (h *UserService) Login(condition *model.LoginStudent) (student model.Student, status string) {
	log.Println("KOYO22225", condition)
	
	student, db := repository.UserRepository().CheckUserPassword(condition)
	if db.Error!= nil {
		return student , responses.Error
	}
	log.Println("KOYO22226", student)
	
	return student , responses.Success
}

// func LoginUser(c *gin.Context){
// 	name := c.PostForm("name")
// 	password := c.PostForm("password")
// 	user := pojo.CheckUserPassword(name, password)
// 	if user.Id == 0 {
// 		c.JSON(http.StatusNotFound, "Error")
// 		return
// 	}
// 	middlewares.SaveSession(c, user.Id)
// 	c.JSON(http.StatusOK, gin.H{
// 		"message" : "Login Successfully",
// 		"User" : user,
// 		"Sessions": middlewares.GetSession(c),
// 	})
// }

// Logout
// func (h *ItemService) Logout(condition *model.Student) (item model.Item, status string) {
// 	item, db := repository.ItemRepository().GetByID(condition)
// 	if db.Error!= nil {
// 		return item, responses.Error
// 	}
// 	return item,responses.Success
// }
