package dao

import (
	"yky-gin/db"
	"yky-gin/models"
)

type UserDAO struct{}

func (dao *UserDAO) FindByUsername(username string) (*models.User, error) {
	var user models.User
	result := db.Db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (dao *UserDAO) FindByPhoneNumber(phoneNumber string) (*models.User, error) {
	var user models.User
	result := db.Db.Where("phone_number = ?", phoneNumber).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (dao *UserDAO) Update(user *models.User) error {
	if err := db.Db.Save(user).Error; err != nil {
		return err
	}
	return nil
}

func (dao *UserDAO) CreateUser(user *models.User) error {
	result := db.Db.Create(user)
	return result.Error
}
