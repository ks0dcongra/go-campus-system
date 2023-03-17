package routes

import (
	"example1/app/http/controller"
	session "example1/app/http/middleware"

	"github.com/gin-gonic/gin"
)

func ApiRoutes(router *gin.Engine) {
	adminApi := router.Group("admin/api")
	userApi := router.Group("user/api", session.SetSession())
	//admin
	adminApi.POST("item", controller.AdminController().AddItem())
	adminApi.GET("item", controller.AdminController().GetItem())
	//user
	// 虹堡 Authantication
	// userApi.Use(middleware.AuthRequired) 
	userApi.GET("item", controller.UserController().GetItem())

	// create
	userApi.POST("create", controller.UserController().CreateUser())
	
	// login
	userApi.POST("login",controller.UserController().LoginUser())

	// score search
	userApi.GET("search/:id",controller.UserController().ScoreSearch())

	//logout
	userApi.Use(session.AuthSession())
	{
		userApi.GET("logout",controller.UserController().LogoutUser())
	}
	
}
