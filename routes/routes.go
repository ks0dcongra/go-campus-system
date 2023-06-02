package routes

import (
	"example1/app/http/controller"
	"example1/app/http/middleware"
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

	// 獲得CSRF Token 與 攻擊CSRF之網頁
	router.GET("/index", func(c *gin.Context) {
		// 設定CSRF
		// cookie.SetJWTTokenCookie(c, token)
		token := csrf.Token(c.Request)
		c.Header("X-CSRF-Token", token)
		c.HTML(200, "index.html", gin.H{
			"title":          token,
			csrf.TemplateTag: csrf.TemplateField(c.Request),
		})
	})
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
}
