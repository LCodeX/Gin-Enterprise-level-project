package router

import (
	"yky-gin/controllers"
	"yky-gin/dao"
	middlewares "yky-gin/middleware"
	"yky-gin/services"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1 := r.Group("api/v1")
	{
		user := v1.Group("/user")
		{
			userController := controllers.NewUserController(services.NewUserService(&dao.UserDAO{}))
			user.POST("/register", userController.Register)
			user.POST("/login", userController.Login)
		}
		packages := v1.Group("/packages")
		packages.Use(middlewares.JWTAuth())
		{
			packages.GET("/list", controllers.PackagesController{}.GetPackagesList)
		}
	}
	return r
}
