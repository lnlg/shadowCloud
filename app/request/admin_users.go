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

// 添加用户验证请求
type AddUserRequest struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Nickname string `form:"nickname" json:"nickname" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required,email"`
	Mobile   string `form:"mobile" json:"mobile" binding:"required,phone"`
	Avatar   string `form:"avatar" json:"avatar"`
}

func (AddUserRequest) GetMessages() ValidateErrorMessages {
	return ValidateErrorMessages{
		"Username.required": "用户名不能为空",
		"Password.required": "密码不能为空",
		"Nickname.required": "昵称不能为空",
		"Email.required":    "邮箱不能为空",
		"Email.email":       "邮箱格式不正确",
		"Mobile.required":   "手机号不能为空",
		"Mobile.phone":      "手机号格式错误",
	}
}
