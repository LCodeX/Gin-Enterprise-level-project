package controllers

import (
	"yky-gin/services"
	"yky-gin/utils/resp"

	"github.com/gin-gonic/gin"
)

type AppVersionController struct {
	appVersionService *services.AppVersionService
}

func NewAppVersionController(appVersionService *services.AppVersionService) *AppVersionController {
	return &AppVersionController{appVersionService: appVersionService}
}

func (p *AppVersionController) GetAppVersion(c *gin.Context) {
	packages, _ := p.appVersionService.GetAppVersion()
	resp.RespHelper.OK(c, packages)
}
