package validator

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var customErrorMessages = map[string]string{
	"name_not_admin":  "用户名不能为 'admin'",
	"required":        "该字段为必填项",
	"email":           "请输入有效的邮箱地址",
	"zh_phone_number": "手机号格式错误",
}

func TranslateError(err error) string {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldErr := range validationErrors {
			if msg, exists := customErrorMessages[fieldErr.Tag()]; exists {
				return fmt.Sprintf("字段 '%s': %s", fieldErr.Field(), msg)
			}
		}
	}
	return err.Error()
}
