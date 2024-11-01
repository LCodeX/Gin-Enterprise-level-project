package controllers

import (
	"yky-gin/services"
	"yky-gin/utils/resp"

	"github.com/gin-gonic/gin"
)

type PackagesController struct {
	PackagesService *services.PackageService
}

func NewPackagesController(packagesService *services.PackageService) *PackagesController {
	return &PackagesController{PackagesService: packagesService}
}

func (p *PackagesController) GetPackagesList(c *gin.Context) {
	packages, _ := p.PackagesService.GetPackageList()
	resp.RespHelper.OK(c, packages)
}
