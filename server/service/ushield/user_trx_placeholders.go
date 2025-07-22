package ushield

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ushield"
	ushieldReq "github.com/flipped-aurora/gin-vue-admin/server/model/ushield/request"
)

type UserTrxPlaceholdersService struct{}

// CreateUserTrxPlaceholders 创建userTrxPlaceholders表记录
// Author [yourname](https://github.com/yourname)
func (userTrxPlaceholdersService *UserTrxPlaceholdersService) CreateUserTrxPlaceholders(ctx context.Context, userTrxPlaceholders *ushield.UserTrxPlaceholders) (err error) {
	err = global.GVA_DB.Create(userTrxPlaceholders).Error
	return err
}

// DeleteUserTrxPlaceholders 删除userTrxPlaceholders表记录
// Author [yourname](https://github.com/yourname)
func (userTrxPlaceholdersService *UserTrxPlaceholdersService) DeleteUserTrxPlaceholders(ctx context.Context, id string) (err error) {
	err = global.GVA_DB.Delete(&ushield.UserTrxPlaceholders{}, "id = ?", id).Error
	return err
}

// DeleteUserTrxPlaceholdersByIds 批量删除userTrxPlaceholders表记录
// Author [yourname](https://github.com/yourname)
func (userTrxPlaceholdersService *UserTrxPlaceholdersService) DeleteUserTrxPlaceholdersByIds(ctx context.Context, ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]ushield.UserTrxPlaceholders{}, "id in ?", ids).Error
	return err
}

// UpdateUserTrxPlaceholders 更新userTrxPlaceholders表记录
// Author [yourname](https://github.com/yourname)
func (userTrxPlaceholdersService *UserTrxPlaceholdersService) UpdateUserTrxPlaceholders(ctx context.Context, userTrxPlaceholders ushield.UserTrxPlaceholders) (err error) {
	err = global.GVA_DB.Model(&ushield.UserTrxPlaceholders{}).Where("id = ?", userTrxPlaceholders.Id).Updates(&userTrxPlaceholders).Error
	return err
}
func (userTrxPlaceholdersService *UserTrxPlaceholdersService) UpdateUserTrxPlaceholdersByName(ctx context.Context, _placeholder string, _status int64) (err error) {
	err = global.GVA_DB.Model(&ushield.UserTrxPlaceholders{}).Where("placeholder = ?", _placeholder).Update("status", _status).Error
	return err
}

// GetUserTrxPlaceholders 根据id获取userTrxPlaceholders表记录
// Author [yourname](https://github.com/yourname)
func (userTrxPlaceholdersService *UserTrxPlaceholdersService) GetUserTrxPlaceholders(ctx context.Context, id string) (userTrxPlaceholders ushield.UserTrxPlaceholders, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&userTrxPlaceholders).Error
	return
}

// GetUserTrxPlaceholdersInfoList 分页获取userTrxPlaceholders表记录
// Author [yourname](https://github.com/yourname)
func (userTrxPlaceholdersService *UserTrxPlaceholdersService) GetUserTrxPlaceholdersInfoList(ctx context.Context, info ushieldReq.UserTrxPlaceholdersSearch) (list []ushield.UserTrxPlaceholders, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&ushield.UserTrxPlaceholders{})
	var userTrxPlaceholderss []ushield.UserTrxPlaceholders
	// 如果有条件搜索 下方会自动创建搜索语句

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&userTrxPlaceholderss).Error
	return userTrxPlaceholderss, total, err
}
func (userTrxPlaceholdersService *UserTrxPlaceholdersService) GetUserTrxPlaceholdersPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
