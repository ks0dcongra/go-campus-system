package controller

import (
	"example1/app/http/middleware"
	"example1/app/model"
	"example1/app/model/responses"
	"example1/app/service"
	"fmt"
	"net/http"
	_ "github.com/joho/godotenv/autoload"
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
		if requestData == "0" || requestData == "" {
			c.JSON(http.StatusOK, responses.Status(responses.ParameterErr, nil))
			return
		}
		redisKey := fmt.Sprintf("user_%s", requestData)
		student, status := service.NewUserService().ScoreSearch(requestData, redisKey)
		if status == responses.Error {
			// 失敗
			c.JSON(http.StatusOK, responses.Status(responses.Error, nil))
		} else if status == responses.SuccessDb {
			// 成功但來自DB
			c.JSON(http.StatusOK, responses.Status(responses.SuccessDb, student))
		} else {
			// 成功但來自Redis
			c.JSON(http.StatusOK, responses.Status(responses.SuccessRedis, student))
		}
	}
}
