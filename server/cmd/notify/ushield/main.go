package main

import (
	"context"
	"fmt"
	"github.com/ushield/aurora-admin/server/core"
	"github.com/ushield/aurora-admin/server/global"
	"github.com/ushield/aurora-admin/server/initialize"
	ushieldReq "github.com/ushield/aurora-admin/server/model/ushield/request"
	"github.com/ushield/aurora-admin/server/service"
	"go.uber.org/zap"
	"log"
)

var (
	userAddressMonitorEventService = service.ServiceGroupApp.UshieldServiceGroup.UserAddressMonitorEventService
)

func main() {

	global.GVA_VP = core.Viper() // 初始化Viper

	buddha := `============================================
                       以太坊紧急通知上线
   `
	fmt.Println(buddha)
	global.GVA_DB = initialize.Gorm() // gorm连接数据库
	initialize.DBList()

	global.GVA_LOG = core.Zap() // 初始化zap日志库
	zap.ReplaceGlobals(global.GVA_LOG)

	//得到正在运行的
	var info ushieldReq.UserAddressMonitorEventSearch

	info.Page = 1
	info.PageSize = 1_000_000

	events, _, err := userAddressMonitorEventService.GetUserAddressMonitorEventInfoList(context.Background(), info, 1)
	if err != nil {
		return
	}

	for _, event := range events {
		log.Printf("%v", event)
		log.Printf("times : %d", event.Times)

		event.Times = event.Times + 1
		if event.Times <= 10 {
			userAddressMonitorEventService.UpdateUserAddressMonitorEvent(context.Background(), event)
		}
	}

}
