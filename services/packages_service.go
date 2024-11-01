package services

import (
	"yky-gin/dao"
	"yky-gin/models"
)

func GetPackageList() (models.Packages, error) {
	return dao.GetPackageList()
}
