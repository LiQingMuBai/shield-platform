package ushield

import (
	"context"
	"github.com/ushield/aurora-admin/server/global"
	"github.com/ushield/aurora-admin/server/model/ushield"
	ushieldReq "github.com/ushield/aurora-admin/server/model/ushield/request"
)

type UserUsdtDepositsService struct{}

// CreateUserUsdtDeposits 创建userUsdtDeposits表记录
// Author [yourname](https://github.com/yourname)
func (userUsdtDepositsService *UserUsdtDepositsService) CreateUserUsdtDeposits(ctx context.Context, userUsdtDeposits *ushield.UserUsdtDeposits) (err error) {
	err = global.GVA_DB.Create(userUsdtDeposits).Error
	return err
}

// DeleteUserUsdtDeposits 删除userUsdtDeposits表记录
// Author [yourname](https://github.com/yourname)
func (userUsdtDepositsService *UserUsdtDepositsService) DeleteUserUsdtDeposits(ctx context.Context, id string) (err error) {
	err = global.GVA_DB.Delete(&ushield.UserUsdtDeposits{}, "id = ?", id).Error
	return err
}

// DeleteUserUsdtDepositsByIds 批量删除userUsdtDeposits表记录
// Author [yourname](https://github.com/yourname)
func (userUsdtDepositsService *UserUsdtDepositsService) DeleteUserUsdtDepositsByIds(ctx context.Context, ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]ushield.UserUsdtDeposits{}, "id in ?", ids).Error
	return err
}

// UpdateUserUsdtDeposits 更新userUsdtDeposits表记录
// Author [yourname](https://github.com/yourname)
func (userUsdtDepositsService *UserUsdtDepositsService) UpdateUserUsdtDeposits(ctx context.Context, userUsdtDeposits ushield.UserUsdtDeposits) (err error) {
	err = global.GVA_DB.Model(&ushield.UserUsdtDeposits{}).Where("id = ?", userUsdtDeposits.Id).Updates(&userUsdtDeposits).Error
	return err
}

// GetUserUsdtDeposits 根据id获取userUsdtDeposits表记录
// Author [yourname](https://github.com/yourname)
func (userUsdtDepositsService *UserUsdtDepositsService) GetUserUsdtDeposits(ctx context.Context, id string) (userUsdtDeposits ushield.UserUsdtDeposits, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&userUsdtDeposits).Error
	return
}

func (userUsdtDepositsService *UserUsdtDepositsService) CountUserUsdtDepositsByTxHash(ctx context.Context, txHash string) (records int) {
	var users []ushield.UserUsdtDeposits
	_ = global.GVA_DB.Where("tx_hash = ?", txHash).Find(&users)
	return len(users)
}

// GetUserUsdtDepositsInfoList 分页获取userUsdtDeposits表记录
// Author [yourname](https://github.com/yourname)
func (userUsdtDepositsService *UserUsdtDepositsService) GetUserUsdtDepositsInfoList(ctx context.Context, info ushieldReq.UserUsdtDepositsSearch) (list []ushield.UserUsdtDeposits, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&ushield.UserUsdtDeposits{})
	var userUsdtDepositss []ushield.UserUsdtDeposits
	// 如果有条件搜索 下方会自动创建搜索语句

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&userUsdtDepositss).Error
	return userUsdtDepositss, total, err
}

func (userUsdtDepositsService *UserUsdtDepositsService) GetUserTrxDepositsByStatus(ctx context.Context, _status int64) (list []ushield.UserUsdtDeposits, err error) {
	// 此方法为获取数据源定义的数据
	// 请自行实现

	// 创建db
	db := global.GVA_DB.Model(&ushield.UserUsdtDeposits{})
	var userUsdtDepositss []ushield.UserUsdtDeposits
	// 如果有条件搜索 下方会自动创建搜索语句

	err = db.Where("status = ?", _status).Find(&userUsdtDepositss).Error
	return userUsdtDepositss, err
}

func (userUsdtDepositsService *UserUsdtDepositsService) GetUserUsdtDepositsPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
