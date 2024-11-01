package router

import (
	"yky-gin/controllers"
	middlewares "yky-gin/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	app := controllers.NewApp()
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1 := r.Group("api/v1")
	{
		user := v1.Group("/user")
		{
			user.POST("/register", app.UserController.Register)
			user.POST("/login", app.UserController.Login)
			user.POST("/forget-password", app.UserController.ForgetPassword)
		}
		v1.Use(middlewares.JWTAuth())
		{
			v1.POST("/update-password", app.UserController.UpdateUserPassword)
			v1.GET("/packages", app.PackagesController.GetPackagesList)
			v1.GET("/app-version", app.AppVersionController.GetAppVersion)
			v1.GET("/app-config", app.AppConfigController.GetAppConfig)
		}
	}
	return r
}
