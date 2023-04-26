package controller

import (
	"example1/app/model"
	"example1/app/model/responses"
	"example1/app/service"
	"example1/utils/global"
	"example1/utils/token"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"example1/app/http/repo"
)
type EmployeeService interface {
	GetSrEmployeeNumbers(int) int
}

type EmployeeSerivceImpl struct {
	EmpRepo repo.EmployeeRepo
}

func (es *EmployeeSerivceImpl) GetSrEmployeeNumbers(age int) int {
	log.Println("niceifosdifdsofsf:::::::::")
	srEmps := es.EmpRepo.FindEmployeesAgeGreaterThan(age)
	return len(srEmps)
}

type UserControllerInterface interface {
	LoginUser() gin.HandlerFunc
}

type UserController struct {
	UserService service.UserServiceInterface 
}

func NewUserController() *UserController {
	return &UserController{}
}

func (h *UserController) GetItem() gin.HandlerFunc {
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
// TODO:
/* 
1. 先 refactor controller
2. 偽造物件坐在LoginUser()上
3. 測router的status code就好
*/

// Login TEST Mock 
// func (h *UserController) LoginUser() (model.Student, string) {
// 	log.Println("nice shot100")
// 	jsonData := &model.LoginStudent{Name:"Mar234",Password:"12345678"}
// 	student, status:= h.UserService.Login(jsonData)
// 	return student, status
// }

// Login
func (h *UserController) LoginUser() gin.HandlerFunc {
	log.Println("nice shot1")
	return func(c *gin.Context) {
		log.Println("nice shot2")
		requestData := new(model.LoginStudent)
		if err := c.ShouldBindJSON(requestData); err != nil {
			c.JSON(http.StatusOK, responses.Status(responses.ParameterErr, nil))
			return
		}
		// student, status:= h.UserService.Login(requestData)
		student, status:= service.NewUserService().Login(requestData)

		if status == responses.Success{
			c.JSON(http.StatusOK, responses.Status(responses.Success, gin.H{
				"Student": student,
				// [Session用]:拿到上面session暫存
				// [Session用]:用id存至session暫存
				// middleware.SaveSession(c, student.Id)
				// "Sessions": middleware.GetSession(c),
				// [Token用]:回傳的參數
				// "Token": tokenResult,
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
