package ushield

import (
	"context"
	"errors"
	"github.com/ushield/aurora-admin/server/global"
	"github.com/ushield/aurora-admin/server/model/ushield"
	ushieldReq "github.com/ushield/aurora-admin/server/model/ushield/request"
	"gorm.io/gorm"
)

type UserEnergyOrdersService struct{}

// CreateUserEnergyOrders 创建userEnergyOrders表记录
// Author [yourname](https://github.com/yourname)
func (userEnergyOrdersService *UserEnergyOrdersService) CreateUserEnergyOrders(ctx context.Context, userEnergyOrders *ushield.UserEnergyOrders) (err error) {
	err = global.GVA_DB.Create(userEnergyOrders).Error
	return err
}

// DeleteUserEnergyOrders 删除userEnergyOrders表记录
// Author [yourname](https://github.com/yourname)
func (userEnergyOrdersService *UserEnergyOrdersService) DeleteUserEnergyOrders(ctx context.Context, id string) (err error) {
	err = global.GVA_DB.Delete(&ushield.UserEnergyOrders{}, "id = ?", id).Error
	return err
}

// DeleteUserEnergyOrdersByIds 批量删除userEnergyOrders表记录
// Author [yourname](https://github.com/yourname)
func (userEnergyOrdersService *UserEnergyOrdersService) DeleteUserEnergyOrdersByIds(ctx context.Context, ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]ushield.UserEnergyOrders{}, "id in ?", ids).Error
	return err
}

// UpdateUserEnergyOrders 更新userEnergyOrders表记录
// Author [yourname](https://github.com/yourname)
func (userEnergyOrdersService *UserEnergyOrdersService) UpdateUserEnergyOrders(ctx context.Context, userEnergyOrders ushield.UserEnergyOrders) (err error) {
	err = global.GVA_DB.Model(&ushield.UserEnergyOrders{}).Where("id = ?", userEnergyOrders.ID).Updates(&userEnergyOrders).Error
	return err
}

// GetUserEnergyOrders 根据id获取userEnergyOrders表记录
// Author [yourname](https://github.com/yourname)
func (userEnergyOrdersService *UserEnergyOrdersService) GetUserEnergyOrders(ctx context.Context, id string) (userEnergyOrders ushield.UserEnergyOrders, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&userEnergyOrders).Error
	return
}

// GetUserEnergyOrdersInfoList 分页获取userEnergyOrders表记录
// Author [yourname](https://github.com/yourname)
func (userEnergyOrdersService *UserEnergyOrdersService) GetUserEnergyOrdersInfoList(ctx context.Context, info ushieldReq.UserEnergyOrdersSearch) (list []ushield.UserEnergyOrders, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&ushield.UserEnergyOrders{})
	var userEnergyOrderss []ushield.UserEnergyOrders
	// 如果有条件搜索 下方会自动创建搜索语句

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&userEnergyOrderss).Error
	return userEnergyOrderss, total, err
}

func (userEnergyOrdersService *UserEnergyOrdersService) GetUserEnergyOrderInfoByTxID(txID string) (order ushield.UserEnergyOrders, err error) {
	err = global.GVA_DB.Where("tx_id = ?", txID).First(&order).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 处理没有找到记录的情况
		return order, nil
	} else if err != nil {
		return order, err
	}
	return
}

func (userEnergyOrdersService *UserEnergyOrdersService) GetUserEnergyOrdersPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
