package ushield

import (
	"context"
	"github.com/ushield/aurora-admin/server/global"
	"github.com/ushield/aurora-admin/server/model/ushield"
	ushieldReq "github.com/ushield/aurora-admin/server/model/ushield/request"
)

type UserUsdtSubscriptionsService struct{}

// CreateUserUsdtSubscriptions 创建userUsdtSubscriptions表记录
// Author [yourname](https://github.com/yourname)
func (userUsdtSubscriptionsService *UserUsdtSubscriptionsService) CreateUserUsdtSubscriptions(ctx context.Context, userUsdtSubscriptions *ushield.UserUsdtSubscriptions) (err error) {
	err = global.GVA_DB.Create(userUsdtSubscriptions).Error
	return err
}

// DeleteUserUsdtSubscriptions 删除userUsdtSubscriptions表记录
// Author [yourname](https://github.com/yourname)
func (userUsdtSubscriptionsService *UserUsdtSubscriptionsService) DeleteUserUsdtSubscriptions(ctx context.Context, id string) (err error) {
	err = global.GVA_DB.Delete(&ushield.UserUsdtSubscriptions{}, "id = ?", id).Error
	return err
}

// DeleteUserUsdtSubscriptionsByIds 批量删除userUsdtSubscriptions表记录
// Author [yourname](https://github.com/yourname)
func (userUsdtSubscriptionsService *UserUsdtSubscriptionsService) DeleteUserUsdtSubscriptionsByIds(ctx context.Context, ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]ushield.UserUsdtSubscriptions{}, "id in ?", ids).Error
	return err
}

// UpdateUserUsdtSubscriptions 更新userUsdtSubscriptions表记录
// Author [yourname](https://github.com/yourname)
func (userUsdtSubscriptionsService *UserUsdtSubscriptionsService) UpdateUserUsdtSubscriptions(ctx context.Context, userUsdtSubscriptions ushield.UserUsdtSubscriptions) (err error) {
	err = global.GVA_DB.Model(&ushield.UserUsdtSubscriptions{}).Where("id = ?", userUsdtSubscriptions.Id).Updates(&userUsdtSubscriptions).Error
	return err
}

// GetUserUsdtSubscriptions 根据id获取userUsdtSubscriptions表记录
// Author [yourname](https://github.com/yourname)
func (userUsdtSubscriptionsService *UserUsdtSubscriptionsService) GetUserUsdtSubscriptions(ctx context.Context, id string) (userUsdtSubscriptions ushield.UserUsdtSubscriptions, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&userUsdtSubscriptions).Error
	return
}

// GetUserUsdtSubscriptionsInfoList 分页获取userUsdtSubscriptions表记录
// Author [yourname](https://github.com/yourname)
func (userUsdtSubscriptionsService *UserUsdtSubscriptionsService) GetUserUsdtSubscriptionsInfoList(ctx context.Context, info ushieldReq.UserUsdtSubscriptionsSearch) (list []ushield.UserUsdtSubscriptions, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&ushield.UserUsdtSubscriptions{})
	var userUsdtSubscriptionss []ushield.UserUsdtSubscriptions
	// 如果有条件搜索 下方会自动创建搜索语句

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&userUsdtSubscriptionss).Error
	return userUsdtSubscriptionss, total, err
}
func (userUsdtSubscriptionsService *UserUsdtSubscriptionsService) GetUserUsdtSubscriptionsPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
