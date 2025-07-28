package ushield

import (
	"context"
	"github.com/ushield/aurora-admin/server/global"
	"github.com/ushield/aurora-admin/server/model/ushield"
	ushieldReq "github.com/ushield/aurora-admin/server/model/ushield/request"
)

type UserTrxDepositsService struct{}

// CreateUserTrxDeposits 创建userTrxDeposits表记录
// Author [yourname](https://github.com/yourname)
func (userTrxDepositsService *UserTrxDepositsService) CreateUserTrxDeposits(ctx context.Context, userTrxDeposits *ushield.UserTrxDeposits) (err error) {
	err = global.GVA_DB.Create(userTrxDeposits).Error
	return err
}

// DeleteUserTrxDeposits 删除userTrxDeposits表记录
// Author [yourname](https://github.com/yourname)
func (userTrxDepositsService *UserTrxDepositsService) DeleteUserTrxDeposits(ctx context.Context, id string) (err error) {
	err = global.GVA_DB.Delete(&ushield.UserTrxDeposits{}, "id = ?", id).Error
	return err
}

// DeleteUserTrxDepositsByIds 批量删除userTrxDeposits表记录
// Author [yourname](https://github.com/yourname)
func (userTrxDepositsService *UserTrxDepositsService) DeleteUserTrxDepositsByIds(ctx context.Context, ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]ushield.UserTrxDeposits{}, "id in ?", ids).Error
	return err
}

// UpdateUserTrxDeposits 更新userTrxDeposits表记录
// Author [yourname](https://github.com/yourname)
func (userTrxDepositsService *UserTrxDepositsService) UpdateUserTrxDeposits(ctx context.Context, userTrxDeposits ushield.UserTrxDeposits) (err error) {
	err = global.GVA_DB.Model(&ushield.UserTrxDeposits{}).Where("id = ?", userTrxDeposits.Id).Updates(&userTrxDeposits).Error
	return err
}

// GetUserTrxDeposits 根据id获取userTrxDeposits表记录
// Author [yourname](https://github.com/yourname)
func (userTrxDepositsService *UserTrxDepositsService) GetUserTrxDeposits(ctx context.Context, id string) (userTrxDeposits ushield.UserTrxDeposits, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&userTrxDeposits).Error
	return
}

// GetUserTrxDepositsInfoList 分页获取userTrxDeposits表记录
// Author [yourname](https://github.com/yourname)
func (userTrxDepositsService *UserTrxDepositsService) GetUserTrxDepositsInfoList(ctx context.Context, info ushieldReq.UserTrxDepositsSearch) (list []ushield.UserTrxDeposits, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&ushield.UserTrxDeposits{})
	var userTrxDepositss []ushield.UserTrxDeposits
	// 如果有条件搜索 下方会自动创建搜索语句

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&userTrxDepositss).Error
	return userTrxDepositss, total, err
}
func (userTrxDepositsService *UserTrxDepositsService) GetUserTrxDepositsByStatus(ctx context.Context, _status int64) (list []ushield.UserTrxDeposits, err error) {
	// 此方法为获取数据源定义的数据
	// 请自行实现

	// 创建db
	db := global.GVA_DB.Model(&ushield.UserTrxDeposits{})
	var userTrxDepositss []ushield.UserTrxDeposits
	// 如果有条件搜索 下方会自动创建搜索语句

	err = db.Where("status = ?", _status).Find(&userTrxDepositss).Error
	return userTrxDepositss, err
}
func (userTrxDepositsService *UserTrxDepositsService) GetUserTrxDepositsPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
