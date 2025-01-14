package request

// 自定义验证结构体
type TestParam struct {
	Username   string `form:"username" json:"username" binding:"required"`
	Password   string `form:"password" json:"password" binding:"required"`
	RePassword string `form:"re_password" json:"re_password" binding:"required,eqfield=Password"`
	Phone      string `form:"phone" json:"phone" binding:"required,phone"`
}

// 实现Validator接口 自定义错误信息
func (TestParam) GetMessages() ValidateErrorMessages {
	return ValidateErrorMessages{
		"Username.required":   "用户名称不能为空",
		"Password.required":   "密码不能为空",
		"RePassword.required": "二次密码不能为空",
		"RePassword.eqfield":  "两次密码不一致",
		"Phone.required":      "手机号不能为空",
		"Phone.phone":         "手机号格式错误",
	}
}

// 自定义验证结构体-可以给字段加一些其他tag信息，方面form，json的解析
type TestBindTest struct {
	Username   string `form:"username" json:"username" binding:"required"`
	Password   string `form:"password" json:"password" binding:"required"`
	RePassword string `form:"re_password" json:"re_password" binding:"required,eqfield=Password"`
	Email      string `form:"email" json:"email" binding:"required,email"`
	Age        int    `form:"age" json:"age" binding:"required,gte=0,lte=130"`
	Phone      string `form:"phone" json:"phone" binding:"required,phone"`
}
