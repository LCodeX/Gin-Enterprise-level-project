package router

import (
	"yky-gin/controllers"
	middlewares "yky-gin/middleware"
	"yky-gin/utils/logger"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router(env string) *gin.Engine {
	app := controllers.NewApp()
	gin.SetMode(map[bool]string{true: gin.ReleaseMode, false: gin.DebugMode}[env == "PROD"])
	r := gin.New()
	// infoWriter := io.MultiWriter(os.Stdout, logger.GinInfoLogger().Out)
	// errorWriter := io.MultiWriter(os.Stdout, logger.GinErrorLogger().Out)
	if env == "PROD" {
		r.Use(gin.Logger())
		//r.Use(gin.LoggerWithWriter(infoWriter))
		r.Use(gin.RecoveryWithWriter(logger.GinErrorLogger().Out))
	} else {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/panic", func(c *gin.Context) {
		panic("This is a test panic")
	})
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
