package routes

import (
	"example1/app/http/controller"
	// "example1/app/http/middleware"

	"github.com/gin-gonic/gin"
)

func ApiRoutes(router *gin.Engine) {
	adminApi := router.Group("admin/api")
	userApi := router.Group("user/api")
	//admin
	adminApi.POST("item", controller.AdminController().AddItem())
	adminApi.GET("item", controller.AdminController().GetItem())
	//user
	// userApi.Use(middleware.AuthRequired)
	userApi.GET("item", controller.UserController().GetItem())

	//login && logout
	userApi.POST("login",controller.UserController().LoginUser())
	// userApi.GET("logout",controller.UserController().LogoutUser())

}
