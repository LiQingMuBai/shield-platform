package ushield

import (
	"context"
	"github.com/ushield/aurora-admin/server/global"
	"github.com/ushield/aurora-admin/server/model/ushield"
	ushieldReq "github.com/ushield/aurora-admin/server/model/ushield/request"
)

type UserPackageSubscriptionsService struct{}

// CreateUserPackageSubscriptions 创建userPackageSubscriptions表记录
// Author [yourname](https://github.com/yourname)
func (userPackageSubscriptionsService *UserPackageSubscriptionsService) CreateUserPackageSubscriptions(ctx context.Context, userPackageSubscriptions *ushield.UserPackageSubscriptions) (err error) {
	err = global.GVA_DB.Create(userPackageSubscriptions).Error
	return err
}

// DeleteUserPackageSubscriptions 删除userPackageSubscriptions表记录
// Author [yourname](https://github.com/yourname)
func (userPackageSubscriptionsService *UserPackageSubscriptionsService) DeleteUserPackageSubscriptions(ctx context.Context, id string) (err error) {
	err = global.GVA_DB.Delete(&ushield.UserPackageSubscriptions{}, "id = ?", id).Error
	return err
}

// DeleteUserPackageSubscriptionsByIds 批量删除userPackageSubscriptions表记录
// Author [yourname](https://github.com/yourname)
func (userPackageSubscriptionsService *UserPackageSubscriptionsService) DeleteUserPackageSubscriptionsByIds(ctx context.Context, ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]ushield.UserPackageSubscriptions{}, "id in ?", ids).Error
	return err
}

// UpdateUserPackageSubscriptions 更新userPackageSubscriptions表记录
// Author [yourname](https://github.com/yourname)
func (userPackageSubscriptionsService *UserPackageSubscriptionsService) UpdateUserPackageSubscriptions(ctx context.Context, userPackageSubscriptions ushield.UserPackageSubscriptions) (err error) {
	err = global.GVA_DB.Model(&ushield.UserPackageSubscriptions{}).Where("id = ?", userPackageSubscriptions.Id).Updates(&userPackageSubscriptions).Error
	return err
}

// GetUserPackageSubscriptions 根据id获取userPackageSubscriptions表记录
// Author [yourname](https://github.com/yourname)
func (userPackageSubscriptionsService *UserPackageSubscriptionsService) GetUserPackageSubscriptions(ctx context.Context, id string) (userPackageSubscriptions ushield.UserPackageSubscriptions, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&userPackageSubscriptions).Error
	return
}

// GetUserPackageSubscriptionsInfoList 分页获取userPackageSubscriptions表记录
// Author [yourname](https://github.com/yourname)
func (userPackageSubscriptionsService *UserPackageSubscriptionsService) GetUserPackageSubscriptionsInfoList(ctx context.Context, info ushieldReq.UserPackageSubscriptionsSearch) (list []ushield.UserPackageSubscriptions, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&ushield.UserPackageSubscriptions{})
	var userPackageSubscriptionss []ushield.UserPackageSubscriptions
	// 如果有条件搜索 下方会自动创建搜索语句

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&userPackageSubscriptionss).Error
	return userPackageSubscriptionss, total, err
}
func (userPackageSubscriptionsService *UserPackageSubscriptionsService) GetUserPackageSubscriptionsPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
