
package ushield

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ushield"
    ushieldReq "github.com/flipped-aurora/gin-vue-admin/server/model/ushield/request"
)

type UserAddressMonitorEventService struct {}
// CreateUserAddressMonitorEvent 创建userAddressMonitorEvent表记录
// Author [yourname](https://github.com/yourname)
func (userAddressMonitorEventService *UserAddressMonitorEventService) CreateUserAddressMonitorEvent(ctx context.Context, userAddressMonitorEvent *ushield.UserAddressMonitorEvent) (err error) {
	err = global.GVA_DB.Create(userAddressMonitorEvent).Error
	return err
}

// DeleteUserAddressMonitorEvent 删除userAddressMonitorEvent表记录
// Author [yourname](https://github.com/yourname)
func (userAddressMonitorEventService *UserAddressMonitorEventService)DeleteUserAddressMonitorEvent(ctx context.Context, id string) (err error) {
	err = global.GVA_DB.Delete(&ushield.UserAddressMonitorEvent{},"id = ?",id).Error
	return err
}

// DeleteUserAddressMonitorEventByIds 批量删除userAddressMonitorEvent表记录
// Author [yourname](https://github.com/yourname)
func (userAddressMonitorEventService *UserAddressMonitorEventService)DeleteUserAddressMonitorEventByIds(ctx context.Context, ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]ushield.UserAddressMonitorEvent{},"id in ?",ids).Error
	return err
}

// UpdateUserAddressMonitorEvent 更新userAddressMonitorEvent表记录
// Author [yourname](https://github.com/yourname)
func (userAddressMonitorEventService *UserAddressMonitorEventService)UpdateUserAddressMonitorEvent(ctx context.Context, userAddressMonitorEvent ushield.UserAddressMonitorEvent) (err error) {
	err = global.GVA_DB.Model(&ushield.UserAddressMonitorEvent{}).Where("id = ?",userAddressMonitorEvent.Id).Updates(&userAddressMonitorEvent).Error
	return err
}

// GetUserAddressMonitorEvent 根据id获取userAddressMonitorEvent表记录
// Author [yourname](https://github.com/yourname)
func (userAddressMonitorEventService *UserAddressMonitorEventService)GetUserAddressMonitorEvent(ctx context.Context, id string) (userAddressMonitorEvent ushield.UserAddressMonitorEvent, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&userAddressMonitorEvent).Error
	return
}
// GetUserAddressMonitorEventInfoList 分页获取userAddressMonitorEvent表记录
// Author [yourname](https://github.com/yourname)
func (userAddressMonitorEventService *UserAddressMonitorEventService)GetUserAddressMonitorEventInfoList(ctx context.Context, info ushieldReq.UserAddressMonitorEventSearch) (list []ushield.UserAddressMonitorEvent, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&ushield.UserAddressMonitorEvent{})
    var userAddressMonitorEvents []ushield.UserAddressMonitorEvent
    // 如果有条件搜索 下方会自动创建搜索语句
    
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&userAddressMonitorEvents).Error
	return  userAddressMonitorEvents, total, err
}
func (userAddressMonitorEventService *UserAddressMonitorEventService)GetUserAddressMonitorEventPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
