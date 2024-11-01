package dao

import (
	"yky-gin/db"
	"yky-gin/models"
)

type AppVersionDAO struct{}

func (dao *AppVersionDAO) GetAppVersion() (*models.AppVersion, error) {
	var appVerson models.AppVersion
	result := db.Db.First(&appVerson)
	if result.Error != nil {
		return nil, result.Error
	}
	return &appVerson, nil
}
