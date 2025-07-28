
package ushield

import (
	"context"
	"github.com/ushield/aurora-admin/server/global"
	"github.com/ushield/aurora-admin/server/model/ushield"
    ushieldReq "github.com/ushield/aurora-admin/server/model/ushield/request"
)

type SysAnnouncementsInfoService struct {}
// CreateSysAnnouncementsInfo 创建sysAnnouncementsInfo表记录
// Author [yourname](https://github.com/yourname)
func (sysAnnouncementsInfoService *SysAnnouncementsInfoService) CreateSysAnnouncementsInfo(ctx context.Context, sysAnnouncementsInfo *ushield.SysAnnouncementsInfo) (err error) {
	err = global.GVA_DB.Create(sysAnnouncementsInfo).Error
	return err
}

// DeleteSysAnnouncementsInfo 删除sysAnnouncementsInfo表记录
// Author [yourname](https://github.com/yourname)
func (sysAnnouncementsInfoService *SysAnnouncementsInfoService)DeleteSysAnnouncementsInfo(ctx context.Context, id string) (err error) {
	err = global.GVA_DB.Delete(&ushield.SysAnnouncementsInfo{},"id = ?",id).Error
	return err
}

// DeleteSysAnnouncementsInfoByIds 批量删除sysAnnouncementsInfo表记录
// Author [yourname](https://github.com/yourname)
func (sysAnnouncementsInfoService *SysAnnouncementsInfoService)DeleteSysAnnouncementsInfoByIds(ctx context.Context, ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]ushield.SysAnnouncementsInfo{},"id in ?",ids).Error
	return err
}

// UpdateSysAnnouncementsInfo 更新sysAnnouncementsInfo表记录
// Author [yourname](https://github.com/yourname)
func (sysAnnouncementsInfoService *SysAnnouncementsInfoService)UpdateSysAnnouncementsInfo(ctx context.Context, sysAnnouncementsInfo ushield.SysAnnouncementsInfo) (err error) {
	err = global.GVA_DB.Model(&ushield.SysAnnouncementsInfo{}).Where("id = ?",sysAnnouncementsInfo.Id).Updates(&sysAnnouncementsInfo).Error
	return err
}

// GetSysAnnouncementsInfo 根据id获取sysAnnouncementsInfo表记录
// Author [yourname](https://github.com/yourname)
func (sysAnnouncementsInfoService *SysAnnouncementsInfoService)GetSysAnnouncementsInfo(ctx context.Context, id string) (sysAnnouncementsInfo ushield.SysAnnouncementsInfo, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&sysAnnouncementsInfo).Error
	return
}
// GetSysAnnouncementsInfoInfoList 分页获取sysAnnouncementsInfo表记录
// Author [yourname](https://github.com/yourname)
func (sysAnnouncementsInfoService *SysAnnouncementsInfoService)GetSysAnnouncementsInfoInfoList(ctx context.Context, info ushieldReq.SysAnnouncementsInfoSearch) (list []ushield.SysAnnouncementsInfo, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&ushield.SysAnnouncementsInfo{})
    var sysAnnouncementsInfos []ushield.SysAnnouncementsInfo
    // 如果有条件搜索 下方会自动创建搜索语句
    
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&sysAnnouncementsInfos).Error
	return  sysAnnouncementsInfos, total, err
}
func (sysAnnouncementsInfoService *SysAnnouncementsInfoService)GetSysAnnouncementsInfoPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
