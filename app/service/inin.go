package service

import "shadowCloud/internal/global"

var (
	BaseService = &baseService{db: global.Db}

	AdminService      = &adminService{baseService: *BaseService}
	AdminUsersService = &adminUsersService{baseService: *BaseService}
)
