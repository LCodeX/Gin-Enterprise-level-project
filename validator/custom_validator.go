package validator

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func StartsWithLetter(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	match, _ := regexp.MatchString(`^[a-zA-Z][a-zA-Z0-9]*$`, value)
	return match
}
