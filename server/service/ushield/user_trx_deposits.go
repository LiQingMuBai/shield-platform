package ushield

import (
	"context"
	"fmt"
	"github.com/ushield/aurora-admin/server/global"
	"github.com/ushield/aurora-admin/server/model/ushield"
	ushieldReq "github.com/ushield/aurora-admin/server/model/ushield/request"
	"log"
	"time"
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
func (userTrxDepositsService *UserTrxDepositsService) CountUserTRXDepositsByTxHash(ctx context.Context, txHash string) (records int) {
	var users []ushield.UserTrxDeposits
	_ = global.GVA_DB.Where("tx_hash = ?", txHash).Find(&users)
	return len(users)
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

type DailyDeposit struct {
	Date  time.Time `json:"date"`
	Total float64   `json:"total"`
	Count int       `json:"count"`
}

// 按日期统计金额总和
func (userTrxDepositsService *UserTrxDepositsService) GetDailyTRXDeposits() (list []DailyDeposit, err error) {

	var results []DailyDeposit

	// 使用SQLite的date函数提取日期部分，并转换字符串金额为浮点数进行求和
	result := global.GVA_DB.Model(&ushield.UserTrxDeposits{}).
		Select("date(created_at) as date, "+
			"sum(cast(amount as real)) as total, "+
			"count(*) as count").Where("status = ?", 1).
		Group("date(created_at)").
		Order("date desc").
		Scan(&results)

	if result.Error != nil {
		log.Fatal("failed to query daily deposits: ", result.Error)
	}

	fmt.Println("\n每日存款统计:")
	fmt.Println("日期\t\t总计\t\t笔数")
	for _, r := range results {
		fmt.Printf("%s\t%.2f\t\t%d\n", r.Date.Format("2006-01-02"), r.Total, r.Count)
	}

	return results, err
}
