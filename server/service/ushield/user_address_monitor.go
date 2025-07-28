package ushield

import (
	"context"
	"github.com/ushield/aurora-admin/server/global"
	"github.com/ushield/aurora-admin/server/model/ushield"
	ushieldReq "github.com/ushield/aurora-admin/server/model/ushield/request"
)

type UserAddressMonitorService struct{}

// CreateUserAddressMonitor 创建userAddressMonitor表记录
// Author [yourname](https://github.com/yourname)
func (userAddressMonitorService *UserAddressMonitorService) CreateUserAddressMonitor(ctx context.Context, userAddressMonitor *ushield.UserAddressMonitor) (err error) {
	err = global.GVA_DB.Create(userAddressMonitor).Error
	return err
}

// DeleteUserAddressMonitor 删除userAddressMonitor表记录
// Author [yourname](https://github.com/yourname)
func (userAddressMonitorService *UserAddressMonitorService) DeleteUserAddressMonitor(ctx context.Context, id string) (err error) {
	err = global.GVA_DB.Delete(&ushield.UserAddressMonitor{}, "id = ?", id).Error
	return err
}

// DeleteUserAddressMonitorByIds 批量删除userAddressMonitor表记录
// Author [yourname](https://github.com/yourname)
func (userAddressMonitorService *UserAddressMonitorService) DeleteUserAddressMonitorByIds(ctx context.Context, ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]ushield.UserAddressMonitor{}, "id in ?", ids).Error
	return err
}

// UpdateUserAddressMonitor 更新userAddressMonitor表记录
// Author [yourname](https://github.com/yourname)
func (userAddressMonitorService *UserAddressMonitorService) UpdateUserAddressMonitor(ctx context.Context, userAddressMonitor ushield.UserAddressMonitor) (err error) {
	err = global.GVA_DB.Model(&ushield.UserAddressMonitor{}).Where("id = ?", userAddressMonitor.Id).Updates(&userAddressMonitor).Error
	return err
}

// GetUserAddressMonitor 根据id获取userAddressMonitor表记录
// Author [yourname](https://github.com/yourname)
func (userAddressMonitorService *UserAddressMonitorService) GetUserAddressMonitor(ctx context.Context, id string) (userAddressMonitor ushield.UserAddressMonitor, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&userAddressMonitor).Error
	return
}

// GetUserAddressMonitorInfoList 分页获取userAddressMonitor表记录
// Author [yourname](https://github.com/yourname)
func (userAddressMonitorService *UserAddressMonitorService) GetUserAddressMonitorInfoList(ctx context.Context, info ushieldReq.UserAddressMonitorSearch) (list []ushield.UserAddressMonitor, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&ushield.UserAddressMonitor{})
	var userAddressMonitors []ushield.UserAddressMonitor
	// 如果有条件搜索 下方会自动创建搜索语句

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&userAddressMonitors).Error
	return userAddressMonitors, total, err
}
func (userAddressMonitorService *UserAddressMonitorService) GetUserAddressMonitorPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
