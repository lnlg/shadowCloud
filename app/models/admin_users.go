package models

import "shadowCloud/internal/global"

type AdminUser struct {
	ID            int       `json:"id"`
	Username      string    `json:"username"`
	Password      string    `json:"password"`
	Nickname      string    `json:"nickname"`
	Avatar        string    `json:"avatar"`
	Email         string    `json:"email"`
	Mobile        string    `json:"mobile"`
	LastLoginIp   LocalTime `json:"last_login_ip"`
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
	err := global.Db.Where("username = ? and is_deleted = 0", username).First(&user).Error
	return user, err
}
