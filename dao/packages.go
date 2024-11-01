package dao

import (
	"yky-gin/db"
	"yky-gin/models"
)

type PackagesDao struct{}

func (dao *PackagesDao) GetPackageList() (models.Packages, error) {
	var packages models.Packages
	err := db.Db.Find(&packages).Error
	return packages, err
}
