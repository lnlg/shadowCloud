package request

// 后台登录验证请求
type LoginAdminRequest struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func (LoginAdminRequest) GetMessages() ValidateErrorMessages {
	return ValidateErrorMessages{
		"Username.required": "用户名不能为空",
		"Password.required": "密码不能为空",
	}
}
