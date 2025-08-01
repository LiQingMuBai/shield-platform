
package ushield

import (
	"context"
	"github.com/ushield/aurora-admin/server/global"
	"github.com/ushield/aurora-admin/server/model/ushield"
    ushieldReq "github.com/ushield/aurora-admin/server/model/ushield/request"
)

type UserBundleEnergyOrdersService struct {}
// CreateUserBundleEnergyOrders 创建userBundleEnergyOrders表记录
// Author [yourname](https://github.com/yourname)
func (userBundleEnergyOrdersService *UserBundleEnergyOrdersService) CreateUserBundleEnergyOrders(ctx context.Context, userBundleEnergyOrders *ushield.UserBundleEnergyOrders) (err error) {
	err = global.GVA_DB.Create(userBundleEnergyOrders).Error
	return err
}

// DeleteUserBundleEnergyOrders 删除userBundleEnergyOrders表记录
// Author [yourname](https://github.com/yourname)
func (userBundleEnergyOrdersService *UserBundleEnergyOrdersService)DeleteUserBundleEnergyOrders(ctx context.Context, id string) (err error) {
	err = global.GVA_DB.Delete(&ushield.UserBundleEnergyOrders{},"id = ?",id).Error
	return err
}

// DeleteUserBundleEnergyOrdersByIds 批量删除userBundleEnergyOrders表记录
// Author [yourname](https://github.com/yourname)
func (userBundleEnergyOrdersService *UserBundleEnergyOrdersService)DeleteUserBundleEnergyOrdersByIds(ctx context.Context, ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]ushield.UserBundleEnergyOrders{},"id in ?",ids).Error
	return err
}

// UpdateUserBundleEnergyOrders 更新userBundleEnergyOrders表记录
// Author [yourname](https://github.com/yourname)
func (userBundleEnergyOrdersService *UserBundleEnergyOrdersService)UpdateUserBundleEnergyOrders(ctx context.Context, userBundleEnergyOrders ushield.UserBundleEnergyOrders) (err error) {
	err = global.GVA_DB.Model(&ushield.UserBundleEnergyOrders{}).Where("id = ?",userBundleEnergyOrders.Id).Updates(&userBundleEnergyOrders).Error
	return err
}

// GetUserBundleEnergyOrders 根据id获取userBundleEnergyOrders表记录
// Author [yourname](https://github.com/yourname)
func (userBundleEnergyOrdersService *UserBundleEnergyOrdersService)GetUserBundleEnergyOrders(ctx context.Context, id string) (userBundleEnergyOrders ushield.UserBundleEnergyOrders, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&userBundleEnergyOrders).Error
	return
}
// GetUserBundleEnergyOrdersInfoList 分页获取userBundleEnergyOrders表记录
// Author [yourname](https://github.com/yourname)
func (userBundleEnergyOrdersService *UserBundleEnergyOrdersService)GetUserBundleEnergyOrdersInfoList(ctx context.Context, info ushieldReq.UserBundleEnergyOrdersSearch) (list []ushield.UserBundleEnergyOrders, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&ushield.UserBundleEnergyOrders{})
    var userBundleEnergyOrderss []ushield.UserBundleEnergyOrders
    // 如果有条件搜索 下方会自动创建搜索语句
    
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&userBundleEnergyOrderss).Error
	return  userBundleEnergyOrderss, total, err
}
func (userBundleEnergyOrdersService *UserBundleEnergyOrdersService)GetUserBundleEnergyOrdersPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
