package main

import (
	"context"
	"fmt"
	"github.com/ushield/aurora-admin/server/core"
	"github.com/ushield/aurora-admin/server/global"
	"github.com/ushield/aurora-admin/server/initialize"
	"github.com/ushield/aurora-admin/server/service"
	"github.com/ushield/aurora-admin/server/service/ushield"
	"github.com/xuri/excelize/v2"
	"go.uber.org/zap"
	"log"
	"time"
)

var (
	tgUsersService = service.ServiceGroupApp.UshieldServiceGroup.TgUsersService
)

type App struct {
	ticker *time.Ticker
	done   chan bool
	logger *log.Logger
}

func main() {

	global.GVA_VP = core.Viper() // 初始化Viper

	buddha := `============================================
                       统计新增用户数上线
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

	userStat, err := tgUsersService.QueryDailyNewUsersBuilder(context.Background())

	if err != nil {

	}
	for _, record := range userStat {

		fmt.Printf("record : %v\n", record)

	}

	generateExcelWithChart(userStat)
	//c := cron.New()
	//
	//// 4. 每天 12 点执行任务
	//_, err := c.AddFunc("0 12 * * *", func() {
	//	log.Println("开始执行每日任务：更新 times=0 的 status=0")
	//
	//	// 5. 使用 GORM 更新符合条件的记录
	//	error1 := tgUsersService.UpdateTgUsersTimes(context.Background())
	//
	//	if error1 != nil {
	//		log.Printf("更新失败: %v", error1)
	//		return
	//	}
	//
	//	//log.Printf("更新成功，影响行数: %d", result.RowsAffected)
	//})
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

	//if err != nil {
	//	log.Fatalf("定时任务设置失败: %v", err)
	//}
	//
	//// 6. 启动定时任务
	//c.Start()
	//log.Println("定时任务已启动，每天 0 点执行")
	//
	//// 7. 保持程序运行
	//select {}

}

// generateExcelWithChart 生成包含折线图的Excel文件
func generateExcelWithChart(stats []ushield.DailyUserStat) error {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			log.Println("关闭Excel文件时出错:", err)
		}
	}()

	// 设置工作表名称为"用户统计"
	sheetName := "用户统计"
	index, err := f.NewSheet(sheetName)
	if err != nil {
		return err
	}

	// 设置表头
	f.SetCellValue(sheetName, "A1", "日期")
	f.SetCellValue(sheetName, "B1", "新增用户数")

	// 填充数据
	for i, stat := range stats {
		row := i + 2 // 从第2行开始（第1行是表头）
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", row), stat.Date)
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", row), stat.Count)
	}

	// 设置列宽
	f.SetColWidth(sheetName, "A", "A", 12) // 日期列宽
	f.SetColWidth(sheetName, "B", "B", 15) // 数量列宽

	// 创建折线图
	if err := createLineChart(f, sheetName, len(stats)); err != nil {
		return err
	}

	// 设置默认工作表
	f.SetActiveSheet(index)

	// 保存文件
	filename := fmt.Sprintf("用户统计_%s.xlsx", time.Now().Format("20060102_150405"))
	if err := f.SaveAs(filename); err != nil {
		return err
	}

	return nil
}

// createLineChart 在Excel中创建折线图
func createLineChart(f *excelize.File, sheetName string, dataPoints int) error {
	if dataPoints == 0 {
		return nil // 没有数据时不创建图表
	}

	// 定义图表数据范围
	categoriesRange := fmt.Sprintf("%s!$A$2:$A$%d", sheetName, dataPoints+1)
	valuesRange := fmt.Sprintf("%s!$B$2:$B$%d", sheetName, dataPoints+1)

	// 创建折线图
	chart := &excelize.Chart{
		Type: excelize.Line,
		Series: []excelize.ChartSeries{
			{
				Name:       fmt.Sprintf("%s!$B$1", sheetName),
				Categories: categoriesRange,
				Values:     valuesRange,
			},
		},
		Title: []excelize.RichTextRun{
			{
				Text: "每日新增用户统计",
			},
		},
		Legend: excelize.ChartLegend{
			Position: "bottom",
		},
		XAxis: excelize.ChartAxis{
			Title: []excelize.RichTextRun{
				{
					Text: "日期",
				},
			},
		},
		YAxis: excelize.ChartAxis{
			Title: []excelize.RichTextRun{
				{
					Text: "用户数量",
				},
			},
		},
	}

	// 添加图表到工作表
	if err := f.AddChart(sheetName, "D2", chart); err != nil {
		return err
	}

	return nil
}
