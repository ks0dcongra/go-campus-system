package routes

import (
	"example1/app/http/controller"
	"example1/app/http/middleware"
	"github.com/gin-gonic/gin"
)

func ApiRoutes(router *gin.Engine) {
	adminApi := router.Group("admin/api")

	userApi := router.Group("user/api")
	// [Session用]:每次進行user相關操作都會產生一個Session
	// userApi := router.Group("user/api", session.SetSession())

	//adminItem
	adminApi.POST("item", controller.AdminController().AddItem())
	adminApi.GET("item", controller.AdminController().GetItem())

	// 虹堡 Authantication驗證是否為紅寶之人
	// userApi.Use(middleware.AuthRequired)

	//userItem
	userApi.GET("item", controller.UserController().GetItem())

	// create
	userApi.POST("create", controller.UserController().CreateUser())

	// login
	userApi.POST("login", controller.UserController().LoginUser())

	// [Session用]:Session Auth
	// userApi.Use(session.AuthSession())
	// {
	// 	//logout
	// 	userApi.GET("logout", controller.UserController().LogoutUser())
	// 	// score search
	// 	userApi.GET("search/:id", controller.UserController().ScoreSearch())
	// }

	//logout
	userApi.GET("logout", controller.UserController().LogoutUser())
	
	// [Token用]:驗證JWT是否正確設置?
	userApi.Use(middleware.JwtAuthMiddleware())
	{
		// score search
		userApi.GET("search/:id", controller.UserController().ScoreSearch())
	}
}
