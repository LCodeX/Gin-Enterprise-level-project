package controllers

import (
	"yky-gin/dao"
	"yky-gin/services"
)

type App struct {
	AppVersionController *AppVersionController
	PackagesController   *PackagesController
	UserController       *UserController
}

func NewApp() *App {
	return &App{
		AppVersionController: NewAppVersionController(
			services.NewAppVersionService(&dao.AppVersionDAO{}),
		),
		PackagesController: NewPackagesController(
			services.NewPackagesService(&dao.PackagesDao{}),
		),
		UserController: NewUserController(services.NewUserService(&dao.UserDAO{})),
	}
}
