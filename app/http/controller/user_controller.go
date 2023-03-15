package controller

import (
	"example1/app/model"
	"example1/app/model/responses"
	"example1/app/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userController struct {
}

func UserController() *userController {
	return &userController{}
}


func (h *userController) GetItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestData := new(model.SearchItem)
		if err := c.ShouldBindJSON(requestData); err != nil {
			c.JSON(http.StatusOK, responses.Status(responses.ParameterErr, nil))
			return
		}
		item, status := service.NewItemService().Get(requestData)
		if status != responses.Success {
			c.JSON(http.StatusOK, responses.Status(responses.Error, nil))
			return
		}
		c.JSON(http.StatusOK, responses.Status(responses.Success, item))
	}
}

// Login
func (h *userController) LoginUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestData := new(model.LoginStudent)
		if err := c.ShouldBindJSON(requestData); err != nil {
			c.JSON(http.StatusOK, responses.Status(responses.ParameterErr, nil))
			log.Println(err)
			return
		}

		// err := c.ShouldBindJSON(requestData)
		// log.Println(err)
		// if err == nil{
		// 	c.JSON(http.StatusOK, responses.Status(responses.ParameterErr, nil))
		// 	return
		// }
		// if err != nil{
		// 	log.Println("bad")
		// }
		// log.Println("KOYO22222", requestData)
		student, status := service.NewUserService().Login(requestData)
		if student.Id == 0 {
			c.JSON(http.StatusNotFound, "Student Error2")
			return
		}
		log.Println("KOYO22223", requestData)
		//  middleware.SaveSession(c, user.Id)
		if status != responses.Success {
			c.JSON(http.StatusOK, responses.Status(responses.Error, nil))
			return
		}
		log.Println("KOYO22224", requestData)
		c.JSON(http.StatusOK, responses.Status(responses.Success,  gin.H{
			"Student" : student,
			// "Sessions": middleware.GetSession(c),
		}))
	}
}

// pojo
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
// func (h *userController) LogoutUser() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		requestData := new(model.SearchItem)
// 		if err := c.ShouldBindJSON(requestData); err != nil {
// 			c.JSON(http.StatusOK, responses.Status(responses.ParameterErr, nil))
// 			return
// 		}
// 		item, status := service.NewItemService().Logout(requestData)
// 		if status != responses.Success {
// 			c.JSON(http.StatusOK, responses.Status(responses.Error, nil))
// 			return
// 		}
// 		c.JSON(http.StatusOK, responses.Status(responses.Success, item))
// 	}
// }

