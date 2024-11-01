package controllers

import (
	"yky-gin/services"
	"yky-gin/utils/resp"

	"github.com/gin-gonic/gin"
)

type AppConfigController struct {
	appConfigService *services.AppConfigService
}

func NewAppConfigController(appConfigService *services.AppConfigService) *AppConfigController {
	return &AppConfigController{
		appConfigService: appConfigService,
	}
}
func (p *AppConfigController) GetAppConfig(c *gin.Context) {
	appConfig, err := p.appConfigService.GetAppConfig()
	if err != nil {
		resp.RespHelper.Fail(c, resp.Error.Code, err)
		return
	}
	resp.RespHelper.OK(c, appConfig)
}
