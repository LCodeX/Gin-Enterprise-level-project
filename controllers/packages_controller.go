package controllers

import (
	"yky-gin/services"
	"yky-gin/utils/resp"

	"github.com/gin-gonic/gin"
)

type PackagesController struct{}

func (p PackagesController) GetPackagesList(c *gin.Context) {
	packages, _ := services.GetPackageList()
	resp.RespHelper.OK(c, packages)
}
