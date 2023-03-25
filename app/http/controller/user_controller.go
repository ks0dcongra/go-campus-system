package controller

import (
	"example1/app/http/middleware"
	"example1/app/model"
	"example1/app/model/responses"
	"example1/app/service"
	"log"

	// database "example1/database"
	"fmt"
	// "log"
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
			c.JSON(http.StatusOK, responses.Status(responses.ParameterErr, nil, "From DB"))
			return
		}
		item, status := service.NewItemService().Get(requestData)
		if status != responses.Success {
			c.JSON(http.StatusOK, responses.Status(responses.Error, nil, "From DB"))
			return
		}
		c.JSON(http.StatusOK, responses.Status(responses.Success, item, "From DB"))
	}
}

// Login
func (h *userController) LoginUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestData := new(model.LoginStudent)
		if err := c.ShouldBindJSON(requestData); err != nil {
			c.JSON(http.StatusOK, responses.Status(responses.ParameterErr, nil, "From DB"))
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
			c.JSON(http.StatusOK, responses.Status(responses.Error, nil, "From DB"))
			return
		}
		c.JSON(http.StatusOK, responses.Status(responses.Success, gin.H{
			"Student":  student,
			"Sessions": middleware.GetSession(c),
		}, "From DB"))
	}
}

// Logout
func (h *userController) LogoutUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		middleware.ClearSession(c)
		c.JSON(http.StatusOK, responses.Status(responses.Success, gin.H{
			"message": "Logout Successfully.",
		}, "From DB"))
	}
}

// Create User
func (h *userController) CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestData := new(model.CreateStudent)
		if err := c.ShouldBindJSON(requestData); err != nil {
			c.JSON(http.StatusOK, responses.Status(responses.ParameterErr, nil, "From DB"))
			return
		}
		student_id, status := service.NewUserService().CreateUser(requestData)
		if status != responses.Success {
			c.JSON(http.StatusOK, responses.Status(responses.Error, nil, "From DB"))
			return
		}
		c.JSON(http.StatusOK, responses.Status(responses.Success, student_id, "From DB"))
	}
}



// ScoreSearch
func (h *userController) ScoreSearch() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestData := c.Param("id")
		if (requestData == "0" || requestData == "" ){
			c.JSON(http.StatusOK, responses.Status(responses.ParameterErr, nil, "From DB"))
			return
		}
		redisKey := fmt.Sprintf("user_%s", requestData)
		
		student, status := service.NewUserService().ScoreSearch(requestData, redisKey)
		log.Println("student:",student)
		log.Println("status:",status)
		if status == responses.Error {
			// 失敗
			c.JSON(http.StatusOK, responses.Status(responses.Error, nil, "From DB"))
		}else if status == responses.SuccessDb {
			// 成功但來自DB
			c.JSON(http.StatusOK, responses.Status(responses.SuccessDb, student, "From DB"))
		}else{
			// 成功但來自Redis
			c.JSON(http.StatusOK, responses.Status(responses.SuccessRedis, student, "From Redis"))
		}
	}
}


// func redis {
// 	database: Redis連線
// 	service:拿到連線結果
// 	service:判斷REdIS有沒有Key存在
//  service:取得Key存在的成果
// 	repo: 如果沒有去創造KEY去暫存於Redis中
// }
