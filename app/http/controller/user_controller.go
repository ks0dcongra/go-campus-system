package controller

import (
	"example1/app/http/middleware"
	"example1/app/model"
	"example1/app/model/responses"
	"example1/app/service"
	"github.com/gin-gonic/gin"
	"net/http"
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
		c.JSON(http.StatusOK, responses.Status(responses.Success, gin.H{
			"Student":  student,
			"Sessions": middleware.GetSession(c),
		}))
	}
}

// Logout
func (h *userController) LogoutUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		middleware.ClearSession(c)
		c.JSON(http.StatusOK, responses.Status(responses.Success, gin.H{
			"message": "Logout Successfully.",
		}))
	}
}

// Create User
func (h *userController) CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestData := new(model.CreateStudent)
		if err := c.ShouldBindJSON(requestData); err != nil {
			c.JSON(http.StatusOK, responses.Status(responses.ParameterErr, nil))
			return
		}
		student_id, status := service.NewUserService().CreateUser(requestData)
		if status != responses.Success {
			c.JSON(http.StatusOK, responses.Status(responses.Error, nil))
			return
		}
		c.JSON(http.StatusOK, responses.Status(responses.Success, student_id))
	}
}

// ScoreSearch
func (h *userController) ScoreSearch() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestData := c.Param("id")
		if requestData == "" {
			c.JSON(http.StatusOK, responses.Status(responses.ParameterErr, nil))
			return
		}
		student, status := service.NewUserService().ScoreSearch(requestData)

		if status != responses.Success {
			c.JSON(http.StatusOK, responses.Status(responses.Error, nil))
			return
		}
		c.JSON(http.StatusOK, responses.Status(responses.Success, student))
	}
}
