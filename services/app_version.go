package services

import (
	"yky-gin/dao"
	"yky-gin/models"
)

type AppVersionService struct {
	AppversionDAO *dao.AppVersionDAO
}

func NewAppVersionService(appVersionDao *dao.AppVersionDAO) *AppVersionService {
	return &AppVersionService{
		AppversionDAO: appVersionDao,
	}
}

func (appService *AppVersionService) GetAppVersion() (*models.AppVersion, error) {
	return appService.AppversionDAO.GetAppVersion()
}
