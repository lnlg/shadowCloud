package validator

import (
	"regexp"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// 初始化验证器
func InitValidator() *validator.Validate {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("phone", ValidatePhone)
		return v
	}
	return nil
}

// 验证手机号
func ValidatePhone(fl validator.FieldLevel) bool {
	phone := fl.Field().String()
	if len(phone) != 11 {
		return false
	}
	// 正则表达式验证手机号
	pattern := `^1[3-9][0-9]{9}$`
	ok, err := regexp.MatchString(pattern, phone)
	if err != nil {
		return false
	}
	return ok
}
