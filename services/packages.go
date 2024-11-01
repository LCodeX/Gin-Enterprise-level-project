package services

import (
	"yky-gin/dao"
	"yky-gin/models"
)

type PackageService struct {
	PackageDao *dao.PackagesDao
}

func NewPackagesService(packagesDao *dao.PackagesDao) *PackageService {
	return &PackageService{
		PackageDao: packagesDao,
	}
}

func (p *PackageService) GetPackageList() (models.Packages, error) {
	return p.PackageDao.GetPackageList()
}
