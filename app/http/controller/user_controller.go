package controller

import (
	"example1/app/model"
	"example1/app/model/responses"
	"example1/app/service"
	"example1/utils/global"
	"example1/utils/token"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
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
		student, status, tokenResult := service.NewUserService().Login(requestData)
		if student.Id == 0 {
			c.JSON(http.StatusNotFound, responses.Status(responses.Error, gin.H{"message": "student not found!"}))
			return
		}
		// [Session用]:用id存至session暫存
		// middleware.SaveSession(c, student.Id)
		if status != responses.Success {
			c.JSON(http.StatusOK, responses.Status(responses.Error, nil))
			return
		}
		c.JSON(http.StatusOK, responses.Status(responses.Success, gin.H{
			"Student": student,
			// [Session用]:拿到上面session暫存
			// "Sessions": middleware.GetSession(c),
			// [Token用]:回傳的參數
			"Token": tokenResult,
		}))
	}
}

// Logout
func (h *userController) LogoutUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		// [Session用]:清除目前Session
		// middleware.ClearSession(c)
		// [Token用]:取得Header
		tokenString := c.GetHeader("Authorization")
		global.Blacklist[tokenString] = true // 将 Token 加入黑名单
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
		// [Token用]:先將uint轉換成int再運用strconv轉換成string。
		user_id, err := token.ExtractTokenID(c)
		str_user_id := strconv.Itoa(int(user_id))
		// [Token用]:限制只有本人能查詢分數，如果Token login時所暫存的user_id與傳入c的user_id不相符，則回傳只限本人查詢分數。
		if str_user_id != requestData {
			c.JSON(http.StatusOK, responses.Status(responses.ScoreTokenErr, nil))
			return
		}
		// [Token用]:Token那邊出錯了!
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.Status(responses.TokenErr, nil))
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
