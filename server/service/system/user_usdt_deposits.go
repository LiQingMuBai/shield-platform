
package system

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
    systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
)

type UserUsdtDepositsService struct {}
// CreateUserUsdtDeposits 创建userUsdtDeposits表记录
// Author [yourname](https://github.com/yourname)
func (userUsdtDepositsService *UserUsdtDepositsService) CreateUserUsdtDeposits(ctx context.Context, userUsdtDeposits *system.UserUsdtDeposits) (err error) {
	err = global.GVA_DB.Create(userUsdtDeposits).Error
	return err
}

// DeleteUserUsdtDeposits 删除userUsdtDeposits表记录
// Author [yourname](https://github.com/yourname)
func (userUsdtDepositsService *UserUsdtDepositsService)DeleteUserUsdtDeposits(ctx context.Context, id string) (err error) {
	err = global.GVA_DB.Delete(&system.UserUsdtDeposits{},"id = ?",id).Error
	return err
}

// DeleteUserUsdtDepositsByIds 批量删除userUsdtDeposits表记录
// Author [yourname](https://github.com/yourname)
func (userUsdtDepositsService *UserUsdtDepositsService)DeleteUserUsdtDepositsByIds(ctx context.Context, ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]system.UserUsdtDeposits{},"id in ?",ids).Error
	return err
}

// UpdateUserUsdtDeposits 更新userUsdtDeposits表记录
// Author [yourname](https://github.com/yourname)
func (userUsdtDepositsService *UserUsdtDepositsService)UpdateUserUsdtDeposits(ctx context.Context, userUsdtDeposits system.UserUsdtDeposits) (err error) {
	err = global.GVA_DB.Model(&system.UserUsdtDeposits{}).Where("id = ?",userUsdtDeposits.Id).Updates(&userUsdtDeposits).Error
	return err
}

// GetUserUsdtDeposits 根据id获取userUsdtDeposits表记录
// Author [yourname](https://github.com/yourname)
func (userUsdtDepositsService *UserUsdtDepositsService)GetUserUsdtDeposits(ctx context.Context, id string) (userUsdtDeposits system.UserUsdtDeposits, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&userUsdtDeposits).Error
	return
}
// GetUserUsdtDepositsInfoList 分页获取userUsdtDeposits表记录
// Author [yourname](https://github.com/yourname)
func (userUsdtDepositsService *UserUsdtDepositsService)GetUserUsdtDepositsInfoList(ctx context.Context, info systemReq.UserUsdtDepositsSearch) (list []system.UserUsdtDeposits, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&system.UserUsdtDeposits{})
    var userUsdtDepositss []system.UserUsdtDeposits
    // 如果有条件搜索 下方会自动创建搜索语句
    
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&userUsdtDepositss).Error
	return  userUsdtDepositss, total, err
}
func (userUsdtDepositsService *UserUsdtDepositsService)GetUserUsdtDepositsPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
