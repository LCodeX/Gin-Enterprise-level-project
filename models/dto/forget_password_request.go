package dto

type ForgetPasswordRequest struct {
	PhoneNumber string `json:"phone_number" binding:"required,zh_phone_number"`
	Password    string `json:"password" binding:"required"`
	SmsCode     string `json:"sms_code" binding:"required"`
}

type UpdatePasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=6,max=20"`
}
