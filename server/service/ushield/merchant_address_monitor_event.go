package ushield

import (
	"context"
	"github.com/ushield/aurora-admin/server/global"
	"github.com/ushield/aurora-admin/server/model/ushield"
	ushieldReq "github.com/ushield/aurora-admin/server/model/ushield/request"
)

type MerchantAddressMonitorEventService struct{}

// CreateMerchantAddressMonitorEvent 创建merchantAddressMonitorEvent表记录
// Author [yourname](https://github.com/yourname)
func (merchantAddressMonitorEventService *MerchantAddressMonitorEventService) CreateMerchantAddressMonitorEvent(ctx context.Context, merchantAddressMonitorEvent *ushield.MerchantAddressMonitorEvent) (err error) {
	err = global.GVA_DB.Create(merchantAddressMonitorEvent).Error
	return err
}

// DeleteMerchantAddressMonitorEvent 删除merchantAddressMonitorEvent表记录
// Author [yourname](https://github.com/yourname)
func (merchantAddressMonitorEventService *MerchantAddressMonitorEventService) DeleteMerchantAddressMonitorEvent(ctx context.Context, id string) (err error) {
	err = global.GVA_DB.Delete(&ushield.MerchantAddressMonitorEvent{}, "id = ?", id).Error
	return err
}

// DeleteMerchantAddressMonitorEventByIds 批量删除merchantAddressMonitorEvent表记录
// Author [yourname](https://github.com/yourname)
func (merchantAddressMonitorEventService *MerchantAddressMonitorEventService) DeleteMerchantAddressMonitorEventByIds(ctx context.Context, ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]ushield.MerchantAddressMonitorEvent{}, "id in ?", ids).Error
	return err
}

// UpdateMerchantAddressMonitorEvent 更新merchantAddressMonitorEvent表记录
// Author [yourname](https://github.com/yourname)
func (merchantAddressMonitorEventService *MerchantAddressMonitorEventService) UpdateMerchantAddressMonitorEvent(ctx context.Context, merchantAddressMonitorEvent ushield.MerchantAddressMonitorEvent) (err error) {
	err = global.GVA_DB.Model(&ushield.MerchantAddressMonitorEvent{}).Where("id = ?", merchantAddressMonitorEvent.Id).Updates(&merchantAddressMonitorEvent).Error
	return err
}

// GetMerchantAddressMonitorEvent 根据id获取merchantAddressMonitorEvent表记录
// Author [yourname](https://github.com/yourname)
func (merchantAddressMonitorEventService *MerchantAddressMonitorEventService) GetMerchantAddressMonitorEvent(ctx context.Context, id string) (merchantAddressMonitorEvent ushield.MerchantAddressMonitorEvent, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&merchantAddressMonitorEvent).Error
	return
}

func (merchantAddressMonitorEventService *MerchantAddressMonitorEventService) GetMerchantAddressMonitorEventByAddressAndUser(ctx context.Context, _address string, _userID uint) (merchantAddressMonitorEvent ushield.MerchantAddressMonitorEvent, err error) {
	err = global.GVA_DB.Where("address = ?", _address).Where("_userID = ?", _userID).First(&merchantAddressMonitorEvent).Error
	return
}

// GetMerchantAddressMonitorEventInfoList 分页获取merchantAddressMonitorEvent表记录
// Author [yourname](https://github.com/yourname)
func (merchantAddressMonitorEventService *MerchantAddressMonitorEventService) GetMerchantAddressMonitorEventInfoList(ctx context.Context, info ushieldReq.MerchantAddressMonitorEventSearch) (list []ushield.MerchantAddressMonitorEvent, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&ushield.MerchantAddressMonitorEvent{})
	var merchantAddressMonitorEvents []ushield.MerchantAddressMonitorEvent
	// 如果有条件搜索 下方会自动创建搜索语句

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&merchantAddressMonitorEvents).Error
	return merchantAddressMonitorEvents, total, err
}
func (merchantAddressMonitorEventService *MerchantAddressMonitorEventService) GetMerchantAddressMonitorEventPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
