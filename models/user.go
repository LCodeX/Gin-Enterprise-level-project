package models

import (
	"time"
)

type User struct {
	ID                   uint64     `json:"id" gorm:"primaryKey"`
	AvatarURL            string     `json:"avatar_url"`
	Nickname             string     `json:"nickname"`
	PhoneNumber          string     `json:"phone_number" gorm:"unique"`
	City                 string     `json:"city"`
	IdentityType         string     `json:"identity_type"`
	MainBusiness         string     `json:"main_business"`
	Username             string     `json:"username" gorm:"unique"`
	Password             string     `json:"-"`
	IsVip                bool       `json:"is_vip"`
	MembershipExpiryDate *time.Time `json:"membership_ex_date"`
	IsDisabled           bool       `json:"is_disabled"`
	CreatedAt            time.Time  `json:"-"`
	UpdatedAt            time.Time  `json:"-"`
	DeletedAt            *time.Time `json:"-"`
}

func (User) TableName() string {
	return "app_user"
}
