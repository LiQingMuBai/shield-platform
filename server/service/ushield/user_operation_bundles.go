package ushield

import (
	"context"
	"github.com/ushield/aurora-admin/server/global"
	"github.com/ushield/aurora-admin/server/model/ushield"
	ushieldReq "github.com/ushield/aurora-admin/server/model/ushield/request"
)

type UserOperationBundlesService struct{}

// CreateUserOperationBundles 创建userOperationBundles表记录
// Author [yourname](https://github.com/yourname)
func (userOperationBundlesService *UserOperationBundlesService) CreateUserOperationBundles(ctx context.Context, userOperationBundles *ushield.UserOperationBundles) (err error) {
	err = global.GVA_DB.Create(userOperationBundles).Error
	return err
}

// DeleteUserOperationBundles 删除userOperationBundles表记录
// Author [yourname](https://github.com/yourname)
func (userOperationBundlesService *UserOperationBundlesService) DeleteUserOperationBundles(ctx context.Context, id string) (err error) {
	err = global.GVA_DB.Delete(&ushield.UserOperationBundles{}, "id = ?", id).Error
	return err
}

// DeleteUserOperationBundlesByIds 批量删除userOperationBundles表记录
// Author [yourname](https://github.com/yourname)
func (userOperationBundlesService *UserOperationBundlesService) DeleteUserOperationBundlesByIds(ctx context.Context, ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]ushield.UserOperationBundles{}, "id in ?", ids).Error
	return err
}

// UpdateUserOperationBundles 更新userOperationBundles表记录
// Author [yourname](https://github.com/yourname)
func (userOperationBundlesService *UserOperationBundlesService) UpdateUserOperationBundles(ctx context.Context, userOperationBundles ushield.UserOperationBundles) (err error) {
	err = global.GVA_DB.Model(&ushield.UserOperationBundles{}).Where("id = ?", userOperationBundles.Id).Updates(&userOperationBundles).Error
	return err
}

// GetUserOperationBundles 根据id获取userOperationBundles表记录
// Author [yourname](https://github.com/yourname)
func (userOperationBundlesService *UserOperationBundlesService) GetUserOperationBundles(ctx context.Context, id string) (userOperationBundles ushield.UserOperationBundles, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&userOperationBundles).Error
	return
}

// GetUserOperationBundlesInfoList 分页获取userOperationBundles表记录
// Author [yourname](https://github.com/yourname)
func (userOperationBundlesService *UserOperationBundlesService) GetUserOperationBundlesInfoList(ctx context.Context, info ushieldReq.UserOperationBundlesSearch) (list []ushield.UserOperationBundles, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&ushield.UserOperationBundles{})
	var userOperationBundless []ushield.UserOperationBundles
	// 如果有条件搜索 下方会自动创建搜索语句

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&userOperationBundless).Error
	return userOperationBundless, total, err
}
func (userOperationBundlesService *UserOperationBundlesService) GetUserOperationBundlesPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
