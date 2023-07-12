package routes

import (
	"example1/app/http/controller"
	"example1/app/http/middleware"
	"example1/app/model"
	"example1/database"
	"example1/utils/random"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/csrf"
)

func ApiRoutes(router *gin.Engine) {
	// 设置静态文件目录（可选）
	router.Static("/static", "./static")
	// 设置HTML模板文件目录（可选）
	router.LoadHTMLGlob("view/*")

	router.Use(middleware.CSRF())
	router.Use(middleware.CSRFToken())

	userApi := router.Group("user/api")
	// [Session用]:每次進行user相關操作都會產生一個Session
	// userApi := router.Group("user/api", session.SetSession())
	// 虹堡 Authantication驗證是否為紅寶之人
	// userApi.Use(middleware.AuthRequired)
	// 原來的CreateUser
	// userService := service.NewUserService()
	// [Session用]:Session Auth
	// userApi.Use(session.AuthSession())
	// {
	// // logout
	// 	userApi.GET("logout", controller.UserController().LogoutUser())
	// // score search
	// 	userApi.GET("search/:id", controller.UserController().ScoreSearch())
	// }

	// create
	userApi.POST("create", controller.NewUserController().CreateUser())

	// 模擬CSRF攻擊手法：delete
	userApi.DELETE("delete/:id", controller.NewUserController().DeleteUser())

	// login
	userApi.POST("login", controller.NewUserController().LoginUser())

	// [Token用]:驗證JWT是否正確設置?
	userApi.Use(middleware.JwtAuthMiddleware())
	{
		// logout
		userApi.GET("logout", controller.NewUserController().LogoutUser())
		// score search
		userApi.GET("search/:id", controller.NewUserController().ScoreSearch())
	}

	// 獲得CSRF Token 與 攻擊CSRF之網頁
	router.GET("/index", func(c *gin.Context) {
		// 設定CSRF
		// cookie.SetJWTTokenCookie(c, token)
		name := c.Query("name")
		var students []model.Student
		database.DB.Find(&students)
		token := csrf.Token(c.Request)
		c.Header("X-CSRF-Token", token)
		c.HTML(200, "index.html", gin.H{
			"name":           name,
			"csrf":           token,
			"students":       students,
			csrf.TemplateTag: csrf.TemplateField(c.Request),
		})
	})

	// OWASP ZAP SQLinjection Testing
	userApi.POST("/SQLinJSON", func(c *gin.Context) {
		var specStudents model.SQLinjectionStudent
		var students []model.Student
		if err := c.ShouldBindJSON(&specStudents); err != nil {
			c.JSON(http.StatusNotAcceptable, err.Error())
			return
		}
		// 對輸入進行 HTML 轉義
		escapedName := template.HTMLEscapeString(specStudents.Name)
		fmt.Println(escapedName)
		database.DB.Where("name = ?", escapedName).Find(&students)

		c.JSON(http.StatusOK, gin.H{
			"students": students,
		},
		)
	})

	// 獲得SQL injection 網頁
	router.GET("/GetSQLinjection", func(c *gin.Context) {
		var students []model.Student
		// Get All Students
		database.DB.Find(&students)
		token := csrf.Token(c.Request)
		var specStudents []model.Student
		// Get Specific Students
		name := c.Query("name")
		query := fmt.Sprintf("SELECT * FROM students WHERE name='%s'", name)
		database.DB.Raw(query).Scan(&specStudents)
		c.HTML(200, "SQLinjection.html", gin.H{
			"csrf":           token,
			"students":       students,
			"specStudents":   specStudents,
			csrf.TemplateTag: csrf.TemplateField(c.Request),
		})
	})

	// 發出POST到SQL injection 網頁
	router.POST("/PostSQLinjection", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		log.Println("password:", password)
		randomNum := random.RandInt(1, 100)
		student := model.Student{
			Name:           username,
			Password:       password,
			Student_number: strconv.Itoa(randomNum),
			CreatedTime:    time.Now(),
			UpdatedTime:    time.Now()}
		log.Println(student)

		database.DB.Create(&student)
		c.Redirect(http.StatusMovedPermanently, "/GetSQLinjection")
	})
}
