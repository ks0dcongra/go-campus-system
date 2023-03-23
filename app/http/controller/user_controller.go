package controller

import (
	"example1/app/http/middleware"
	"example1/app/model"
	"example1/app/model/responses"
	"example1/app/service"
	database "example1/database"
	"fmt"
	"net/http"

	"github.com/gomodule/redigo/redis"
	"github.com/pquerna/ffjson/ffjson"

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
		if requestData == "" {
			c.JSON(http.StatusOK, responses.Status(responses.ParameterErr, nil, "From DB"))
			return
		}
		redisKey := fmt.Sprintf("user_%s", requestData)
		var empty interface{}
		// 連線redis資料庫
		conn := database.RedisDefaultPool.Get()
		// 函式中沒東西可以執行後才會操作，資料庫用完再關閉
		defer conn.Close()
		// 尋找redis裡面有沒有rediskey，如果撈到redis有暫存就不用去撈資料庫了，
		// 如果沒有找到err就會存在就會進入if判斷，轉成Bytes是為了供ffjson套件使用
		data, err := redis.Bytes(conn.Do("GET", redisKey))
		if err != nil {
			student, status := service.NewUserService().ScoreSearch(requestData)
			if status != responses.Success {
				c.JSON(http.StatusOK, responses.Status(responses.Error, nil, "From DB"))
				return
			}
			// 加密成JSON檔，用ffjson比普通的json還快
			redisData, _ := ffjson.Marshal(student)
			// 設置redis的key、value，30秒後掰掰
			conn.Do("SETEX", redisKey, 30, redisData)
			c.JSON(http.StatusOK, responses.Status(responses.Success, student, "From DB"))
		} else {
			// 將Byte解密映射到type User上
			ffjson.Unmarshal(data, &empty)
			c.JSON(http.StatusOK, responses.Status(responses.Success, empty, "From Redis"))
		}
	}
}

// Migration
func (h *userController) Migration() gin.HandlerFunc {
	return func(c *gin.Context) {
		database.DB.Migrator().CreateTable(model.Course{}, model.Student{}, model.Score{})
		// database.DB.Migrator().CreateConstraint(model.Course{}, "id")
		// database.DB.DB()
		c.JSON(http.StatusOK, "Success create")
	}
}

func (h *userController) DropMigration() gin.HandlerFunc {
	return func(c *gin.Context) {
		database.DB.Migrator().DropTable(model.Course{}, model.Student{}, model.Score{})
		// database.DB.DB()
		c.JSON(http.StatusOK, "Success create")
	}
}
