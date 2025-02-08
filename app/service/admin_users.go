package service

import (
	"context"
	"errors"
	"shadowCloud/app/models"
	"shadowCloud/internal/global"
	"shadowCloud/internal/tool"
	"time"
)

// 定义usersService 继承baseService
type adminUsersService struct {
	baseService // nolint:unused
}

// 登录操作
func (a *adminUsersService) Login(username string, password string, login_ip string) (bool, string) {
	userinfo, _ := models.GetAdminUserByUsername(username)
	if userinfo.Username == "" {
		return false, "用户不存在！"
	}
	if userinfo.Password != tool.EncryMd5(password) {
		return false, "密码错误！"
	}
	token := tool.GetRandomStr(16)
	// 保存token对应用户信息到redis
	cache_key := "login_token:" + token
	global.Rdb.HSet(context.Background(), cache_key, "id", int(userinfo.Id), "username", userinfo.Username, "nickname", userinfo.Nickname, "avatar", userinfo.Avatar, "email", userinfo.Email, "mobile", userinfo.Mobile)
	// 设置token有效期2个小时
	global.Rdb.Expire(context.Background(), cache_key, time.Minute*60*2)
	//更新用户最后登录时间和ip
	models.UpdateAdminUserLastLoginInfo(userinfo.Username, login_ip)
	return true, token
}
func (a *adminUsersService) GetUserInfoByToken(token string) (bool, map[string]string) {
	cache_key := "login_token:" + token
	//判断key是否存在
	if global.Rdb.Exists(context.Background(), cache_key).Val() == 0 {
		return false, nil
	}
	userinfo := global.Rdb.HGetAll(context.Background(), cache_key)
	if userinfo == nil {
		return false, nil
	}
	result, err := userinfo.Result()
	if err != nil {
		return false, nil
	}
	return true, result
}

// 创建用户
func (a *adminUsersService) CreateAdminUser(username string, password string, nickname string, email string, mobile string, avatar string) (int, error) {
	//判断用户名是否存在
	userinfo, _ := models.GetAdminUserByUsername(username)
	if userinfo.Username != "" {
		return 0, errors.New("用户名已存在！")
	}
	//判断手机号是否存在
	userinfo, _ = models.GetAdminUserByMobile(mobile)
	if userinfo.Mobile != "" {
		return 0, errors.New("手机号已存在！")
	}
	//判断邮箱是否存在
	userinfo, _ = models.GetAdminUserByEmail(email)
	if userinfo.Email != "" {
		return 0, errors.New("邮箱已存在！")
	}
	user := models.AdminUser{
		Username:      username,
		Password:      tool.EncryMd5(password),
		Nickname:      nickname,
		Email:         email,
		Mobile:        mobile,
		Avatar:        avatar,
		LastLoginIp:   "",
		LastLoginTime: models.LocalTime(time.Now()),
		CreatedAt:     models.LocalTime(time.Now()),
	}
	id, err := models.CreateAdminUser(&user)
	return int(id), err
}
