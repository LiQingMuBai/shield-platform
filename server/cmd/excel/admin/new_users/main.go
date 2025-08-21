package main

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/robfig/cron/v3"
	"github.com/ushield/aurora-admin/server/core"
	"github.com/ushield/aurora-admin/server/global"
	"github.com/ushield/aurora-admin/server/initialize"
	"github.com/ushield/aurora-admin/server/service"
	"github.com/ushield/aurora-admin/server/service/ushield"
	"github.com/xuri/excelize/v2"
	"go.uber.org/zap"
	"log"
	"os"
	"strconv"
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

	c := cron.New()
	// 每天 12 点执行任务
	_, _ = c.AddFunc("0 12 * * *", func() {
		log.Println("开始执行每日任务：发送新增用户")
		userStat, err := tgUsersService.QueryDailyNewUsersBuilder(context.Background())
		if err != nil {
		}
		for _, record := range userStat {
			fmt.Printf("record : %v\n", record)
		}

		var sum int
		for _, dailyStat := range userStat {
			sum += dailyStat.Count
		}

		var lastLine ushield.DailyUserStat

		lastLine.Date = "总计"
		lastLine.Count = sum

		userStat = append(userStat, lastLine)

		fileName, error1 := generateExcelWithChart(userStat)
		if error1 != nil {
			fmt.Printf("generateExcelWithChart error: %v\n", error1)
			return
		}
		sendTG(fileName)
	})
	// 6. 启动定时任务
	c.Start()
	log.Println("定时任务已启动，每天 12 点执行")

	// 7. 保持程序运行
	select {}

}
func sendTG(name string) {

	TG_BOT_API := global.GVA_CONFIG.System.BotToken
	CHAT_ID := global.GVA_CONFIG.System.ChatID
	chatID, _ := strconv.ParseInt(CHAT_ID, 10, 64)
	// 1. 初始化 bot，使用你的 Telegram Bot Token
	bot, err := tgbotapi.NewBotAPI(TG_BOT_API)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true // 开启调试模式

	log.Printf("Authorized on account %s", bot.Self.UserName)

	// 2. 准备要发送的 Excel 文件
	filePath := "./" + name // 替换为你的 Excel 文件路径

	// 3. 读取文件内容
	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		log.Panicf("Error reading file: %v", err)
	}

	// 4. 创建文件上传配置
	// 注意: 群组ID应该是负数，例如 -100123456789

	// 5. 创建文件上传请求
	fileConfig := tgbotapi.NewDocument(chatID, tgbotapi.FileBytes{
		Name:  name, // 接收方看到的文件名
		Bytes: fileBytes,
	})
	//fileConfig.Caption = "这是要发送的 Excel 文件" // 可选的文件描述

	// 6. 发送文件
	if _, err := bot.Send(fileConfig); err != nil {
		log.Panic(err)
	}

	log.Println("Excel file sent successfully!")
}

// generateExcelWithChart 生成包含折线图的Excel文件
func generateExcelWithChart(stats []ushield.DailyUserStat) (string, error) {
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
		return "", err
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
		return "", err
	}

	// 设置默认工作表
	f.SetActiveSheet(index)

	// 保存文件
	filename := fmt.Sprintf("用户统计_%s.xlsx", time.Now().Format("20060102_150405"))
	if err := f.SaveAs(filename); err != nil {
		return "", err
	}

	return filename, nil
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
