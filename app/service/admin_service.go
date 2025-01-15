package service

import "shadowCloud/app/models"

// 定义adminService 继承baseService
type adminService struct {
	baseService // nolint:unused
}

func (a *adminService) Profile() (user models.Test, err error) {
	user, err = models.GetTestOne(2)
	return user, err
}
