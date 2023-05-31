package routes

import (
	"example1/app/http/controller"
	"example1/app/http/middleware"
	// "example1/app/service"

	"github.com/gin-gonic/gin"
)

func ApiRoutes(router *gin.Engine) {

	
	// 设置静态文件目录（可选）
	router.Static("/static", "./static")
	// 设置HTML模板文件目录（可选）
	router.LoadHTMLGlob("view/*")

	router.Use(middleware.CSRF())
	router.Use(middleware.CSRFToken())

	// 測試HTML
	router.GET("/index", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"title": "Hi~訪客",
		})
	})

	userApi := router.Group("user/api")
	// [Session用]:每次進行user相關操作都會產生一個Session
	// userApi := router.Group("user/api", session.SetSession())
	// 虹堡 Authantication驗證是否為紅寶之人
	// userApi.Use(middleware.AuthRequired)
	// 原來的CreateUser
	// userService := service.NewUserService()



	// create
	userApi.POST("create", controller.NewUserController().CreateUser())

	// login
	userApi.POST("login", controller.NewUserController().LoginUser())

	// [Session用]:Session Auth
	// userApi.Use(session.AuthSession())
	// {
	// 	//logout
	// 	userApi.GET("logout", controller.UserController().LogoutUser())
	// 	// score search
	// 	userApi.GET("search/:id", controller.UserController().ScoreSearch())
	// }

	// [Token用]:驗證JWT是否正確設置?
	userApi.Use(middleware.JwtAuthMiddleware())
	{
		// logout
		userApi.GET("logout", controller.NewUserController().LogoutUser())
		// score search
		userApi.GET("search/:id", controller.NewUserController().ScoreSearch())
		// 模擬CSRF攻擊手法：delete
		userApi.DELETE("delete/:id", controller.NewUserController().DeleteUser())
	}
}
