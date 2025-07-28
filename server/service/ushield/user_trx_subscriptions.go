package ushield

import (
	"context"
	"github.com/ushield/aurora-admin/server/global"
	"github.com/ushield/aurora-admin/server/model/ushield"
	ushieldReq "github.com/ushield/aurora-admin/server/model/ushield/request"
)

type UserTrxSubscriptionsService struct{}

// CreateUserTrxSubscriptions 创建userTrxSubscriptions表记录
// Author [yourname](https://github.com/yourname)
func (userTrxSubscriptionsService *UserTrxSubscriptionsService) CreateUserTrxSubscriptions(ctx context.Context, userTrxSubscriptions *ushield.UserTrxSubscriptions) (err error) {
	err = global.GVA_DB.Create(userTrxSubscriptions).Error
	return err
}

// DeleteUserTrxSubscriptions 删除userTrxSubscriptions表记录
// Author [yourname](https://github.com/yourname)
func (userTrxSubscriptionsService *UserTrxSubscriptionsService) DeleteUserTrxSubscriptions(ctx context.Context, id string) (err error) {
	err = global.GVA_DB.Delete(&ushield.UserTrxSubscriptions{}, "id = ?", id).Error
	return err
}

// DeleteUserTrxSubscriptionsByIds 批量删除userTrxSubscriptions表记录
// Author [yourname](https://github.com/yourname)
func (userTrxSubscriptionsService *UserTrxSubscriptionsService) DeleteUserTrxSubscriptionsByIds(ctx context.Context, ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]ushield.UserTrxSubscriptions{}, "id in ?", ids).Error
	return err
}

// UpdateUserTrxSubscriptions 更新userTrxSubscriptions表记录
// Author [yourname](https://github.com/yourname)
func (userTrxSubscriptionsService *UserTrxSubscriptionsService) UpdateUserTrxSubscriptions(ctx context.Context, userTrxSubscriptions ushield.UserTrxSubscriptions) (err error) {
	err = global.GVA_DB.Model(&ushield.UserTrxSubscriptions{}).Where("id = ?", userTrxSubscriptions.Id).Updates(&userTrxSubscriptions).Error
	return err
}

// GetUserTrxSubscriptions 根据id获取userTrxSubscriptions表记录
// Author [yourname](https://github.com/yourname)
func (userTrxSubscriptionsService *UserTrxSubscriptionsService) GetUserTrxSubscriptions(ctx context.Context, id string) (userTrxSubscriptions ushield.UserTrxSubscriptions, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&userTrxSubscriptions).Error
	return
}

// GetUserTrxSubscriptionsInfoList 分页获取userTrxSubscriptions表记录
// Author [yourname](https://github.com/yourname)
func (userTrxSubscriptionsService *UserTrxSubscriptionsService) GetUserTrxSubscriptionsInfoList(ctx context.Context, info ushieldReq.UserTrxSubscriptionsSearch) (list []ushield.UserTrxSubscriptions, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&ushield.UserTrxSubscriptions{})
	var userTrxSubscriptionss []ushield.UserTrxSubscriptions
	// 如果有条件搜索 下方会自动创建搜索语句

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&userTrxSubscriptionss).Error
	return userTrxSubscriptionss, total, err
}
func (userTrxSubscriptionsService *UserTrxSubscriptionsService) GetUserTrxSubscriptionsPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
