package controller

import (
	"example1/app/http/middleware"
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
			return
		}
		student, status := service.NewUserService().Login(requestData)
		if student.Id == 0 {
			c.JSON(http.StatusNotFound, "Student Error2")
			return
		}
		// 用id來儲存session
		middleware.SaveSession(c, student.Id)
		if status != responses.Success {
			c.JSON(http.StatusOK, responses.Status(responses.Error, nil))
			return
		}
		c.JSON(http.StatusOK, responses.Status(responses.Success,  gin.H{
			"Student" : student,
			"Sessions":middleware.GetSession(c),
		}))
	}
}

// Logout
func (h *userController) LogoutUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		middleware.ClearSession(c)
		c.JSON(http.StatusOK, responses.Status(responses.Success,  gin.H{
			"message":"Logout Successfully.",
		}))
	}
}

// Create User
func (h *userController) CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestData := new(model.CreateStudent)
		log.Print("happy",requestData)
		if err := c.ShouldBindJSON(requestData); err != nil {
			c.JSON(http.StatusOK, responses.Status(responses.ParameterErr, nil))
			return
		}
		err := c.ShouldBindJSON(requestData)
		log.Print("happy0",err)
		student_id, status := service.NewUserService().CreateUser(requestData)
		log.Print("happy2",student_id)
		log.Print("happy3",status)
		if status != responses.Success {
			c.JSON(http.StatusOK, responses.Status(responses.Error, nil))
			return
		}
		c.JSON(http.StatusOK, responses.Status(responses.Success, student_id))
	}
}

// Get Score
// func ScoreSearch(c *gin.Context) gin.HandlerFunc{
// 	var students []model.Student
// 	c.BindJSON(&students)
// 	result := database.DB.Preload("courses").Find(&students)
// 	c.JSON(http.StatusOK, result)
// }

// Score Search
// func (h *userController) ScoreSearch() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var student []model.Student
// 		database.DB.Preload("Course").Find(&student, "id = ?", "3")
// 		c.JSON(http.StatusOK, &student)
// 	}
// }

// ScoreSearch
func (h *userController) ScoreSearch() gin.HandlerFunc {
	return func(c *gin.Context) {
		// requestData := new(model.SearchScoreStudent)
		// log.Println(requestData)
		requestData := c.Param("id")
		// log.Println("handsome:",c.Param("id"))
		// if err := c.ShouldBindJSON(requestData); err != nil {
		// 	c.JSON(http.StatusOK, responses.Status(responses.ParameterErr, nil))
		// 	return
		// }
		if requestData == "" {
			c.JSON(http.StatusOK, responses.Status(responses.ParameterErr, nil))
			return
		}
		// if user.Id == 0 {
		// 	c.JSON(http.StatusNotFound,"Error")
		// 	return
		// }
		// err := c.ShouldBindJSON(requestData)
		// log.Print("handsome2",err)
		student, status := service.NewUserService().ScoreSearch(requestData)
		// log.Print("happy3",status)
		if status != responses.Success {
			c.JSON(http.StatusOK, responses.Status(responses.Error, nil))
			return
		}
		c.JSON(http.StatusOK, responses.Status(responses.Success, student))
	}
}
// user := pojo.FindByUserId(c.Param("id"))
// 	if user.Id == 0 {
// 		c.JSON(http.StatusNotFound,"Error")
// 		return
// }
// c.JSON(http.StatusOK, responses.Status(responses.Success,  gin.H{
// 	"Student" : student,
// 	"Sessions":middleware.GetSession(c),
// }))