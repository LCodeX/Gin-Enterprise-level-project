package dto

type ForgetPasswordRequest struct {
	PhoneNumber string `json:"phone_number" binding:"required,zh_phone_number"`
	Password    string `json:"password" binding:"required"`
	SmsCode     string `json:"sms_code" binding:"required"`
}
