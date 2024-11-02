package services

import (
	"errors"
	"yky-gin/dao"
	"yky-gin/models"
	"yky-gin/models/dto"
	"yky-gin/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserDAO *dao.UserDAO
}

func NewUserService(userDAO *dao.UserDAO) *UserService {
	return &UserService{UserDAO: userDAO}
}

func (s *UserService) Register(password, phone_number string, sms_code string) (*models.User, string, error) {
	if _, err := s.UserDAO.FindByPhoneNumber(phone_number); err == nil {
		return nil, "", errors.New("user already exists")
	}
	if sms_code != "test" {
		return nil, "", errors.New("sms code error")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, "", err
	}

	user := &models.User{
		Password:    string(hashedPassword),
		PhoneNumber: phone_number,
	}

	if err := s.UserDAO.CreateUser(user); err != nil {
		return nil, "", err
	}
	token, _ := utils.GenerateToken(user.ID)
	return user, token, nil
}

func (s *UserService) Login(phone_number string, password string) (interface{}, error) {
	user, err := s.UserDAO.FindByPhoneNumber(phone_number)
	if err != nil {
		return "", errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid password")
	}
	token, _ := utils.GenerateToken(user.ID)
	return gin.H{
		"token":    token,
		"userInfo": user,
	}, nil
}

func (s *UserService) UpdateUserPassword(req dto.UpdatePasswordRequest, user_id uint64) error {
	user, err := s.UserDAO.FindByUserId(user_id)
	if err != nil {
		return errors.New("user not found")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.OldPassword)); err != nil {
		return errors.New("old password is not correct")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	if err := s.UserDAO.Update(user); err != nil {
		return errors.New("failed to update password")
	}
	return nil
}

func (s *UserService) ForgotPassword(req dto.ForgetPasswordRequest) error {
	if req.SmsCode != "test" {
		return errors.New("sms code error")
	}
	user, err := s.UserDAO.FindByPhoneNumber(req.PhoneNumber)

	if err != nil {
		return errors.New("user not found")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	if err := s.UserDAO.Update(user); err != nil {
		return errors.New("failed to update password")
	}
	return nil
}
