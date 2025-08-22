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
	tgUsersService          = service.ServiceGroupApp.UshieldServiceGroup.TgUsersService
	userTrxDepositsService  = service.ServiceGroupApp.UshieldServiceGroup.UserTrxDepositsService
	userUsdtDepositsService = service.ServiceGroupApp.UshieldServiceGroup.UserUsdtDepositsService
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
	//
	c := cron.New()
	//// 每天 12 点执行任务
	_, _ = c.AddFunc("59 11 * * *", func() {
		log.Println("开始执行每日任务：发送新增用户")
		userStat, err := tgUsersService.QueryDailyNewUsersBuilder(context.Background())
		if err != nil {
		}
		for _, record := range userStat {
			fmt.Printf("record : %v\n", record)
		}

		results1, err := userTrxDepositsService.GetDailyTRXDeposits()
		results2, err := userUsdtDepositsService.GetDailyUSDTDeposits()

		error1 := generateTxtFile(results1, results2, userStat, "每日运营报表.txt")
		if error1 != nil {
			fmt.Printf("generateExcelWithChart error: %v\n", error1)
			return
		}
		sendTG("每日运营报表.txt")
	})
	//6. 启动定时任务
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
func extractDateSimple(isoString string) string {
	// 简单截取前10个字符（年-月-日部分）
	if len(isoString) >= 10 {
		return isoString[:10]
	}
	return isoString
}

// 生成TXT文件
func generateTxtFile(result1, result2 []ushield.DailyDeposit, stats []ushield.DailyUserStat, filename string) error {
	// 创建文件
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// 写入文件头部
	header := "日期\t\t\t每日人数\n"
	file.WriteString(header)

	// 写入分隔线
	separator := "=============================================\n"
	file.WriteString(separator)

	// 写入每日统计数据
	for _, stat := range stats {
		line := fmt.Sprintf("%s\t\t%d人\n", extractDateSimple(stat.Date), stat.Count)
		file.WriteString(line)
	}

	// 计算并写入总计
	_, totalUsers := calculateTotals(stats)
	totalLine := fmt.Sprintf("\n总计:\t\t%d人\n", totalUsers)
	file.WriteString(separator)
	file.WriteString(totalLine)

	file.WriteString("\n")

	file.WriteString("=================TRX充值统计====================\n")
	file.WriteString("\n")
	// 写入文件头部
	header2 := "日期\t\ttrx充值\t\t每日人数\n"
	file.WriteString(header2)

	// 写入分隔线
	separator2 := "=============================================\n"
	file.WriteString(separator2)

	// 写入每日统计数据
	for _, stat := range result1 {
		line := fmt.Sprintf("%s\t%f\t\t%d人\n", extractDateSimple(stat.Date.String()), stat.Total, stat.Count)
		file.WriteString(line)
	}

	// 计算并写入总计
	totalAmount2, totalUsers2 := calculateTotals2(result1)
	totalLine2 := fmt.Sprintf("\n总计:\t%s\t\t%d人\n", totalAmount2, totalUsers2)
	file.WriteString(separator2)
	file.WriteString(totalLine2)
	file.WriteString("\n")
	file.WriteString("\n")
	file.WriteString("=================USDT充值统计====================\n")
	file.WriteString("\n")
	// 写入文件头部

	// 写入文件头部
	header3 := "日期\t\tusdt充值\t\t每日人数\n"
	file.WriteString(header3)

	// 写入分隔线
	separator3 := "=============================================\n"
	file.WriteString(separator3)

	// 写入每日统计数据
	for _, stat := range result2 {
		line := fmt.Sprintf("%s\t%f\t\t%d人\n", extractDateSimple(stat.Date.String()), stat.Total, stat.Count)
		file.WriteString(line)
	}

	// 计算并写入总计
	totalAmount3, totalUsers3 := calculateTotals2(result2)
	totalLine3 := fmt.Sprintf("\n总计:\t%s\t\t%d人\n", totalAmount3, totalUsers3)
	file.WriteString(separator3)
	file.WriteString(totalLine3)

	return nil
}

// 计算总计金额和总用户数
func calculateTotals2(stats []ushield.DailyDeposit) (string, int) {
	var totalAmount float64
	var totalUsers int

	for _, stat := range stats {
		// 将字符串金额转换为浮点数进行计算
		//amount, err := parseAmount2(stat.Total)
		//if err != nil {
		//	fmt.Printf("转换金额失败: %v, 跳过此项\n", err)
		//	continue
		//}
		totalAmount += stat.Total
		totalUsers += stat.Count
	}

	// 将总计金额格式化为字符串返回
	return fmt.Sprintf("%.2f", totalAmount), totalUsers
}

// 解析金额字符串为浮点数
func parseAmount2(amountStr string) (float64, error) {
	// 这里可以根据您的实际数据格式进行调整
	// 例如，如果金额字符串包含货币符号或逗号，需要先清理
	// cleaned := cleanAmountString(amountStr)

	// 直接转换为浮点数
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		return 0, fmt.Errorf("无法解析金额 '%s': %v", amountStr, err)
	}

	return amount, nil
}

// 计算总计金额和总用户数
func calculateTotals(stats []ushield.DailyUserStat) (string, int) {
	//var totalAmount float64
	var totalUsers int

	for _, stat := range stats {
		// 将字符串金额转换为浮点数进行计算
		//amount, err := parseAmount(stat.Amount)
		//if err != nil {
		//	fmt.Printf("转换金额失败: %v, 跳过此项\n", err)
		//	continue
		//}
		//totalAmount += amount
		totalUsers += stat.Count
	}

	// 将总计金额格式化为字符串返回
	//return fmt.Sprintf("%.2f", totalAmount), totalUsers
	return "", totalUsers
}

// 解析金额字符串为浮点数
func parseAmount(amountStr string) (float64, error) {
	// 这里可以根据您的实际数据格式进行调整
	// 例如，如果金额字符串包含货币符号或逗号，需要先清理
	// cleaned := cleanAmountString(amountStr)

	// 直接转换为浮点数
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		return 0, fmt.Errorf("无法解析金额 '%s': %v", amountStr, err)
	}

	return amount, nil
}

// 清理金额字符串中的非数字字符
func cleanAmountString(amountStr string) string {
	// 移除所有非数字字符（除了小数点和负号）
	// 例如：￥1,250.50 -> 1250.50
	cleaned := make([]rune, 0, len(amountStr))

	for _, r := range amountStr {
		if (r >= '0' && r <= '9') || r == '.' || r == '-' {
			cleaned = append(cleaned, r)
		}
	}

	return string(cleaned)
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
