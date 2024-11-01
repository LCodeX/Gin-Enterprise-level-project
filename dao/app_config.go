package dao

import (
	"yky-gin/db"
	"yky-gin/models"
)

type AppConfigDAO struct{}

func (dao *AppConfigDAO) GetAppConfig() (*models.AppConfig, error) {
	var appConfig *models.AppConfig
	result := db.Db.First(&appConfig)
	if result.Error != nil {
		return nil, result.Error
	}
	return appConfig, nil
}
