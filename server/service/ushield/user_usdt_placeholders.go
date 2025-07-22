package ushield

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ushield"
	ushieldReq "github.com/flipped-aurora/gin-vue-admin/server/model/ushield/request"
)

type UserUsdtPlaceholdersService struct{}

// CreateUserUsdtPlaceholders 创建userUsdtPlaceholders表记录
// Author [yourname](https://github.com/yourname)
func (userUsdtPlaceholdersService *UserUsdtPlaceholdersService) CreateUserUsdtPlaceholders(ctx context.Context, userUsdtPlaceholders *ushield.UserUsdtPlaceholders) (err error) {
	err = global.GVA_DB.Create(userUsdtPlaceholders).Error
	return err
}

// DeleteUserUsdtPlaceholders 删除userUsdtPlaceholders表记录
// Author [yourname](https://github.com/yourname)
func (userUsdtPlaceholdersService *UserUsdtPlaceholdersService) DeleteUserUsdtPlaceholders(ctx context.Context, id string) (err error) {
	err = global.GVA_DB.Delete(&ushield.UserUsdtPlaceholders{}, "id = ?", id).Error
	return err
}

// DeleteUserUsdtPlaceholdersByIds 批量删除userUsdtPlaceholders表记录
// Author [yourname](https://github.com/yourname)
func (userUsdtPlaceholdersService *UserUsdtPlaceholdersService) DeleteUserUsdtPlaceholdersByIds(ctx context.Context, ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]ushield.UserUsdtPlaceholders{}, "id in ?", ids).Error
	return err
}

// UpdateUserUsdtPlaceholders 更新userUsdtPlaceholders表记录
// Author [yourname](https://github.com/yourname)
func (userUsdtPlaceholdersService *UserUsdtPlaceholdersService) UpdateUserUsdtPlaceholders(ctx context.Context, userUsdtPlaceholders ushield.UserUsdtPlaceholders) (err error) {
	err = global.GVA_DB.Model(&ushield.UserUsdtPlaceholders{}).Where("id = ?", userUsdtPlaceholders.Id).Updates(&userUsdtPlaceholders).Error
	return err
}
func (userUsdtPlaceholdersService *UserUsdtPlaceholdersService) UpdateUserUsdtPlaceholdersByName(ctx context.Context, _placeholder string, _status int64) (err error) {
	err = global.GVA_DB.Model(&ushield.UserUsdtPlaceholders{}).Where("placeholder = ?", _placeholder).Update("status", _status).Error
	return err
}

// GetUserUsdtPlaceholders 根据id获取userUsdtPlaceholders表记录
// Author [yourname](https://github.com/yourname)
func (userUsdtPlaceholdersService *UserUsdtPlaceholdersService) GetUserUsdtPlaceholders(ctx context.Context, id string) (userUsdtPlaceholders ushield.UserUsdtPlaceholders, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&userUsdtPlaceholders).Error
	return
}

// GetUserUsdtPlaceholdersInfoList 分页获取userUsdtPlaceholders表记录
// Author [yourname](https://github.com/yourname)
func (userUsdtPlaceholdersService *UserUsdtPlaceholdersService) GetUserUsdtPlaceholdersInfoList(ctx context.Context, info ushieldReq.UserUsdtPlaceholdersSearch) (list []ushield.UserUsdtPlaceholders, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&ushield.UserUsdtPlaceholders{})
	var userUsdtPlaceholderss []ushield.UserUsdtPlaceholders
	// 如果有条件搜索 下方会自动创建搜索语句

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&userUsdtPlaceholderss).Error
	return userUsdtPlaceholderss, total, err
}
func (userUsdtPlaceholdersService *UserUsdtPlaceholdersService) GetUserUsdtPlaceholdersPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
