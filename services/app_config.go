package services

import (
	"yky-gin/dao"
	"yky-gin/models"
)

type AppConfigService struct {
	appConfigDao *dao.AppConfigDAO
}

func NewAppConfigService(appConfigDao *dao.AppConfigDAO) *AppConfigService {
	return &AppConfigService{
		appConfigDao: appConfigDao,
	}
}

func (s *AppConfigService) GetAppConfig() (*models.AppConfig, error) {
	return s.appConfigDao.GetAppConfig()
}
