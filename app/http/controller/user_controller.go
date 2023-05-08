package controller

import (
	"example1/app/model"
	"example1/app/model/responses"
	"example1/app/service"
	"example1/utils/global"
	"example1/utils/token"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

type UserController struct {
	UserService service.UserServiceInterface
}

func NewUserController() *UserController {
	return &UserController{
		UserService: service.NewUserService(),
	}
}

// Login
func (h *UserController) LoginUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestData := new(model.LoginStudent)
		// var login model.LoginStudent
		if err := c.ShouldBindJSON(requestData); err != nil {
			c.JSON(http.StatusNotFound, responses.Status(responses.ParameterErr, nil))
			return
		}
		student, status := h.UserService.Login(requestData)
		// student, status:= service.NewUserService().Login(requestData)
		if status == responses.Success {
			c.JSON(http.StatusOK, responses.Status(responses.Success, gin.H{
				"Student": student,
				// [Session用]:拿到上面session暫存
				// [Session用]:用id存至session暫存
				// middleware.SaveSession(c, student.Id)
				// "Sessions": middleware.GetSession(c),
			}))
			return
		}
		c.JSON(http.StatusNotFound, responses.Status(status, nil))
	}
}

// Logout
func (h *UserController) LogoutUser() gin.HandlerFunc {
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
func (h *UserController) CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestData := new(model.CreateStudent)
		if err := c.ShouldBindJSON(requestData); err != nil {
			c.JSON(http.StatusNotFound, responses.Status(responses.ParameterErr, nil))
			return
		}
		student_id, status := service.NewUserService().CreateUser(requestData)
		if status != responses.Success {
			c.JSON(http.StatusNotFound, responses.Status(responses.Error, nil))
			return
		}
		c.JSON(http.StatusOK, responses.Status(responses.Success, student_id))
	}
}

// ScoreSearch
func (h *UserController) ScoreSearch() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestData := c.Param("id")
		if requestData == "0" || requestData == "" {
			c.JSON(http.StatusOK, responses.Status(responses.ParameterErr, nil))
			return
		}

		// 創建 JwtFactory 實例
		JwtFactory := token.Newjwt()
		// [Token用]:先將uint轉換成int再運用strconv轉換成string。
		user_id, err := JwtFactory.ExtractTokenID(c)
		// [Token用]:Token出錯了!
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.Status(responses.TokenErr, nil))
		}

		student, status := service.NewUserService().ScoreSearch(requestData, user_id)

		if status == responses.SuccessDb || status == responses.SuccessRedis {
			c.JSON(http.StatusOK, responses.Status(status, student))
		} else {
			c.JSON(http.StatusNotFound, responses.Status(status, student))
		}
	}
}
