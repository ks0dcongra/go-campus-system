package routes

import (
	"example1/app/http/controller"
	"example1/app/http/middleware"
	"example1/app/service"

	"github.com/gin-gonic/gin"
)

func ApiRoutes(router *gin.Engine) {
	userApi := router.Group("user/api")
	userService := service.NewUserService()
	// [Session用]:每次進行user相關操作都會產生一個Session
	// userApi := router.Group("user/api", session.SetSession())

	// 虹堡 Authantication驗證是否為紅寶之人
	// userApi.Use(middleware.AuthRequired)

	// create
	userApi.POST("create", controller.NewUserController(userService).CreateUser())

	// login
	userApi.POST("login", controller.NewUserController(userService).LoginUser())

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
		//logout
		userApi.GET("logout", controller.NewUserController(userService).LogoutUser())
		// score search
		userApi.GET("search/:id", controller.NewUserController(userService).ScoreSearch())
	}
}
