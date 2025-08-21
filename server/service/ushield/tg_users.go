package ushield

import (
	"context"
	"github.com/ushield/aurora-admin/server/global"
	"github.com/ushield/aurora-admin/server/model/ushield"
	ushieldReq "github.com/ushield/aurora-admin/server/model/ushield/request"
)

type TgUsersService struct{}

// UserStat 用于存储每日用户统计结果
type DailyUserStat struct {
	Date  string
	Count int
}

// 或者使用GORM的构建器方式（如果数据库支持DATE函数）
func (tgUsersService *TgUsersService) QueryDailyNewUsersBuilder(ctx context.Context) ([]DailyUserStat, error) {
	var stats []DailyUserStat

	err := global.GVA_DB.Model(&ushield.TgUsers{}).
		Select("DATE(created_at) as date, COUNT(*) as count").
		Group("DATE(created_at)").
		Order("date desc").
		Scan(&stats).Error

	if err != nil {
		return nil, err
	}

	return stats, nil
}

// CreateTgUsers 创建tgUsers表记录
// Author [yourname](https://github.com/yourname)
func (tgUsersService *TgUsersService) CreateTgUsers(ctx context.Context, tgUsers *ushield.TgUsers) (err error) {
	err = global.GVA_DB.Create(tgUsers).Error
	return err
}

// DeleteTgUsers 删除tgUsers表记录
// Author [yourname](https://github.com/yourname)
func (tgUsersService *TgUsersService) DeleteTgUsers(ctx context.Context, id string) (err error) {
	err = global.GVA_DB.Delete(&ushield.TgUsers{}, "id = ?", id).Error
	return err
}

// DeleteTgUsersByIds 批量删除tgUsers表记录
// Author [yourname](https://github.com/yourname)
func (tgUsersService *TgUsersService) DeleteTgUsersByIds(ctx context.Context, ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]ushield.TgUsers{}, "id in ?", ids).Error
	return err
}

// UpdateTgUsers 更新tgUsers表记录
// Author [yourname](https://github.com/yourname)
func (tgUsersService *TgUsersService) UpdateTgUsers(ctx context.Context, tgUsers ushield.TgUsers) (err error) {
	err = global.GVA_DB.Model(&ushield.TgUsers{}).Where("id = ?", tgUsers.Id).Updates(&tgUsers).Error
	return err
}
func (tgUsersService *TgUsersService) UpdateTgUsersTimes(ctx context.Context) (err error) {
	err = global.GVA_DB.Model(&ushield.TgUsers{}).Update("status", 0).Error
	return err
}

// GetTgUsers 根据id获取tgUsers表记录
// Author [yourname](https://github.com/yourname)
func (tgUsersService *TgUsersService) GetTgUsersByAssociates(ctx context.Context, _chatID int64) (tgUsers ushield.TgUsers, err error) {
	err = global.GVA_DB.Where("associates = ?", _chatID).First(&tgUsers).Error
	return
}

func (tgUsersService *TgUsersService) GetTgUsers(ctx context.Context, id string) (tgUsers ushield.TgUsers, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&tgUsers).Error
	return
}

// GetTgUsersInfoList 分页获取tgUsers表记录
// Author [yourname](https://github.com/yourname)
func (tgUsersService *TgUsersService) GetTgUsersInfoList(ctx context.Context, info ushieldReq.TgUsersSearch) (list []ushield.TgUsers, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&ushield.TgUsers{})
	var tgUserss []ushield.TgUsers
	// 如果有条件搜索 下方会自动创建搜索语句

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&tgUserss).Error
	return tgUserss, total, err
}
func (tgUsersService *TgUsersService) GetTgUsersPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
