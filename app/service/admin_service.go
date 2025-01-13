package service

import "shadowCloud/app/models"

// 定义adminService 继承baseService
type adminService struct {
	baseService
}

func (a *adminService) Profile() (user models.Test, err error) {
	user, err = models.GetTestOne(2)
	return user, err
}
