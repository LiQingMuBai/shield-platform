package main

import (
	"context"
	"fmt"
	"github.com/robfig/cron/v3"
	"github.com/ushield/aurora-admin/server/core"
	"github.com/ushield/aurora-admin/server/global"
	"github.com/ushield/aurora-admin/server/initialize"
	"github.com/ushield/aurora-admin/server/service"
	"go.uber.org/zap"
	"log"
	"time"
)

var (
	tgUsersService = service.ServiceGroupApp.UshieldServiceGroup.TgUsersService

	userAddressMonitorEventService = service.ServiceGroupApp.UshieldServiceGroup.UserAddressMonitorEventService
)

type App struct {
	ticker *time.Ticker
	done   chan bool
	logger *log.Logger
}

func main() {

	global.GVA_VP = core.Viper() // 初始化Viper

	buddha := `============================================
                       初始化上线
   `
	fmt.Println(buddha)
	global.GVA_DB = initialize.Gorm() // gorm连接数据库
	initialize.DBList()

	global.GVA_LOG = core.Zap() // 初始化zap日志库
	zap.ReplaceGlobals(global.GVA_LOG)

	if global.GVA_DB != nil {
		// 程序结束前关闭数据库链接
		db, _ := global.GVA_DB.DB()
		defer db.Close()
	}
	c := cron.New()

	// 1. 每天 0 点执行任务
	_, err := c.AddFunc("0 0 * * *", func() {
		log.Println("开始执行每日任务：重置任务 time=0, days+1")

		// 2. 使用 GORM 更新符合条件的记录
		//重置用户访问检测地址次数=0
		error1 := tgUsersService.UpdateTgUsersTimes(context.Background())

		if error1 != nil {
			log.Printf("更新失败: %v", error1)
			return
		}

		//把用户地址预警次数+1
		records, error2 := userAddressMonitorEventService.GetUserAddressMonitorEventPublic(context.Background())

		if error2 != nil {
			log.Printf("获取失败：: %v", error2)
		}

		for _, record := range records {

			log.Printf("record: %+v", record)
			record.Days = record.Days + 1
			if record.Days == 31 {
				record.Status = 2
				error4 := userAddressMonitorEventService.UpdateUserAddressMonitorEventStatus(context.Background(), record)
				if error4 != nil {
					continue
				}
			}

			log.Printf("增加天数 : %d", record.Days)
			error3 := userAddressMonitorEventService.UpdateUserAddressMonitorEvent(context.Background(), record)
			if error3 != nil {
				continue
			}

		}
		//log.Printf("更新成功，影响行数: %d", result.RowsAffected)
	})

	if err != nil {
		log.Fatalf("定时任务设置失败: %v", err)
	}

	// 6. 启动定时任务
	c.Start()
	log.Println("定时任务已启动，每天 0 点执行")

	// 7. 保持程序运行
	select {}

}
