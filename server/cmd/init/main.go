package main

import (
	"context"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/core"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/initialize"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"log"
	"time"
)

var (
	currentKeyIndex             uint32
	userService                 = service.ServiceGroupApp.SystemServiceGroup.UserService
	sysOrderService             = service.ServiceGroupApp.SystemServiceGroup.SysOrderService
	userUsdtDepositsService     = service.ServiceGroupApp.UshieldServiceGroup.UserUsdtDepositsService
	userTrxDepositsService      = service.ServiceGroupApp.UshieldServiceGroup.UserTrxDepositsService
	userUsdtPlaceholdersService = service.ServiceGroupApp.UshieldServiceGroup.UserUsdtPlaceholdersService
	userTrxPlaceholdersService  = service.ServiceGroupApp.UshieldServiceGroup.UserTrxPlaceholdersService
	tgUsersService              = service.ServiceGroupApp.UshieldServiceGroup.TgUsersService
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

	// 4. 每天 0 点执行任务
	_, err := c.AddFunc("0 0 * * *", func() {
		log.Println("开始执行每日任务：更新 times=0 的 status=0")

		// 5. 使用 GORM 更新符合条件的记录
		error1 := tgUsersService.UpdateTgUsersTimes(context.Background())

		if error1 != nil {
			log.Printf("更新失败: %v", error1)
			return
		}

		//log.Printf("更新成功，影响行数: %d", result.RowsAffected)
	})
	//执行发送excel表格
	//_, err = c.AddFunc("0 10 * * *", func() {
	//	log.Println("任务2：开始更新 name=1 的记录")
	//
	//	result := db.Model(&TgUser{}).
	//		Where("name = ?", "1").
	//		Update("status", 0) // 假设你要更新 status，按需调整
	//
	//	if result.Error != nil {
	//		log.Printf("任务2失败: %v", result.Error)
	//		return
	//	}
	//
	//	log.Printf("任务2成功，影响行数: %d", result.RowsAffected)
	//})
	//
	//if err != nil {
	//	log.Fatalf("定时任务设置失败: %v", err)
	//}

	if err != nil {
		log.Fatalf("定时任务设置失败: %v", err)
	}

	// 6. 启动定时任务
	c.Start()
	log.Println("定时任务已启动，每天 0 点执行")

	// 7. 保持程序运行
	select {}

}
