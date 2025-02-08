package models

import (
	"shadowCloud/internal/global"
	"shadowCloud/internal/tool"
)

type AdminUser struct {
	Id            uint      `json:"id"`
	Username      string    `json:"username"`
	Password      string    `json:"password"`
	Nickname      string    `json:"nickname"`
	Avatar        string    `json:"avatar"`
	Email         string    `json:"email"`
	Mobile        string    `json:"mobile"`
	LastLoginIp   string    `json:"last_login_ip"`
	LastLoginTime LocalTime `json:"last_login_time"`
	CreatedAt     LocalTime `json:"created_at"`
	UpdatedAt     LocalTime `json:"updated_at"`
	IsDeleted     int       `json:"is_deleted"`
}

// 设置表明
func (AdminUser) TableName() string {
	return "admin_users"
}

// 获取用户信息
func GetAdminUserByUsername(username string) (AdminUser, error) {
	var user AdminUser
	err := global.Db.Where("`username` = ? and `is_deleted` = 0", username).First(&user).Error
	return user, err
}
func GetAdminUserByMobile(mobile string) (AdminUser, error) {
	var user AdminUser
	err := global.Db.Where("`mobile` = ? and `is_deleted` = 0", mobile).First(&user).Error
	return user, err
}
func GetAdminUserByEmail(email string) (AdminUser, error) {
	var user AdminUser
	err := global.Db.Where("`email` = ? and `is_deleted` = 0", email).First(&user).Error
	return user, err
}

// 更新最后登录时间和ip
func UpdateAdminUserLastLoginInfo(username string, ip string) error {
	return global.Db.Model(&AdminUser{}).Where("`username` = ?", username).Update("last_login_ip", ip).Update("last_login_time", tool.GetNowDate()).Error
}

// 创建用户
func CreateAdminUser(user *AdminUser) (uint, error) {
	err := global.Db.Create(user).Error
	return user.Id, err
}
