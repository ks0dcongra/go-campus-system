package routes

import (
	"example1/app/http/controller"
	session "example1/app/http/middleware"

	"github.com/gin-gonic/gin"
)

func ApiRoutes(router *gin.Engine) {
	adminApi := router.Group("admin/api")
	userApi := router.Group("user/api", session.SetSession())
	//adminItem
	adminApi.POST("item", controller.AdminController().AddItem())
	adminApi.GET("item", controller.AdminController().GetItem())

	// 虹堡 Authantication
	// userApi.Use(middleware.AuthRequired)

	//userItem
	userApi.GET("item", controller.UserController().GetItem())

	// create
	userApi.POST("create", controller.UserController().CreateUser())

	// login
	userApi.POST("login", controller.UserController().LoginUser())

	userApi.Use(session.AuthSession())
	{
		//logout
		userApi.GET("logout", controller.UserController().LogoutUser())
		// score search
		userApi.GET("search/:id", controller.UserController().ScoreSearch())
	}
}
