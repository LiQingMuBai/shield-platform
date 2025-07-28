package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/fbsobreira/gotron-sdk/pkg/client"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/robfig/cron/v3"
	"github.com/ushield/aurora-admin/server/core"
	"github.com/ushield/aurora-admin/server/global"
	"github.com/xuri/excelize/v2"
	"google.golang.org/grpc"
	"io"
	"log"
	"math/big"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

type App struct {
	ticker *time.Ticker
	done   chan bool
	logger *log.Logger
}

func main() {

	global.GVA_VP = core.Viper() // 初始化Viper

	buddha := `============================================
                       波场黑名单文件通知上线
   `
	fmt.Println(buddha)
	//global.GVA_DB = initialize.Gorm() // gorm连接数据库
	//initialize.DBList()
	//
	//global.GVA_LOG = core.Zap() // 初始化zap日志库
	//zap.ReplaceGlobals(global.GVA_LOG)
	////
	//if global.GVA_DB != nil {
	//	// 程序结束前关闭数据库链接
	//	db, _ := global.GVA_DB.DB()
	//	defer db.Close()
	//}
	//global.TRONGRID_KEYS = strings.Split(global.GVA_CONFIG.System.TRONGRID_KEYS, ",")
	//
	//log.Println(global.TRONGRID_KEYS)
	// 初始化应用
	app := &App{
		done:   make(chan bool),
		logger: log.New(os.Stdout, "TRON-TRACE-ADDRESS-TASK: ", log.LstdFlags),
	}

	//// 每隔1min启动定时任务
	////app.startScheduler(24 * time.Hour)
	//app.startScheduler(24 * time.Hour)
	//
	//// 等待关闭信号
	//app.waitForShutdown()

	c := cron.New()

	// 4. 每天 0 点执行任务
	_, err := c.AddFunc("59 11 * * *", func() {
		//log.Println("开始执行每日任务：更新 times=0 的 status=0")
		app.executeTask()
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

func (a *App) executeTask() {
	a.logger.Println("波场黑名单文件通知上线-执行定时任务...")
	startTime := time.Now()

	time.Sleep(10 * time.Second)

	url := "https://api.trongrid.io/v1/contracts/TBPxhVAsuzoFnKyXtc1o2UySEydPHgATto/events"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("accept", "application/json")
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	//fmt.Println(string(body))
	var events TronEvents
	if err := json.Unmarshal(body, &events); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	var txIDMap map[string]string /*创建集合 */
	txIDMap = make(map[string]string)

	for _, datum := range events.Data {
		//log.Println(datum.EventName)
		//log.Println(datum.TransactionID)
		txIDMap[datum.TransactionID] = datum.EventName
	}

	sumbitMap := make(map[string]int64)
	commitMap := make(map[string]int64)
	for txID := range txIDMap {
		result := getTransactionData(txID)

		if len(result.RawDataHex) > 600 {
			tAddress := getTronAddress(result)
			_, _amount := getBalance(tAddress)
			//预备拉入黑名單
			sumbitMap[tAddress] = _amount
		} else {
			//已經拉入黑名單
			_address, _amount := getCommitAddressBalance(txID)
			commitMap[_address] = _amount
		}
	}

	for address, target := range sumbitMap {

		log.Println("address ", address)
		log.Println("amount ", target)
	}
	exportExcel(sumbitMap, "24小时内波场预冻结.xlsx")
	exportExcel(commitMap, "24小时内波场已冻结.xlsx")

	time.Sleep(time.Second * 10)

	//
	a.sendTG("24小时内波场预冻结.xlsx")

	a.sendTG("24小时内波场已冻结.xlsx")

	a.logger.Printf("任务完成， 耗时: %v", time.Since(startTime))

}

func (a *App) sendTG(name string) {
	// 1. 初始化 bot，使用你的 Telegram Bot Token
	bot, err := tgbotapi.NewBotAPI(global.GVA_CONFIG.System.GroupBotToken)
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
	chatID := global.GVA_CONFIG.System.GroupChatID // 替换为你的群组ID (注意是负数)

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

// -4934902849
// 8168393119:AAGwYmRLi_Dlwsxpl2TncPa0W951gOTZMsU
func exportExcel(source map[string]int64, fileName string) {
	data := map[string]float64{}
	for k, v := range source {
		data[k] = float64(v) / 1000000
	}

	log.Println(data)

	// 创建一个新的 Excel 文件
	f := excelize.NewFile()

	// 设置表头
	headers := []string{"Address", "Value"}
	for col, header := range headers {
		cell, _ := excelize.CoordinatesToCellName(col+1, 1) // 列从1开始，行从1开始
		f.SetCellValue("Sheet1", cell, header)
	}

	// 填充数据
	row := 1 // 从第二行开始填充数据
	for address, value := range data {
		// 写入地址
		cell, _ := excelize.CoordinatesToCellName(1, row)
		f.SetColWidth("Sheet1", "1", "2", 50)
		f.SetCellValue("Sheet1", cell, address)

		// 写入值
		cell, _ = excelize.CoordinatesToCellName(2, row)
		f.SetCellValue("Sheet1", cell, value)

		row++
	}
	_cell1, _ := excelize.CoordinatesToCellName(1, row)
	f.SetCellValue("Sheet1", _cell1, "总计")
	_cell2, _ := excelize.CoordinatesToCellName(2, row)

	_total := 0.0
	for _, _value := range data {
		_total = _total + _value
	}

	f.SetCellValue("Sheet1", _cell2, _total)
	autoAdjustColumnWidth(f, "Sheet1", "A")
	// 创建折线图
	if err := createLineChart(f, len(data)); err != nil {
		log.Fatalf("创建折线图失败: %v", err)
	}

	// 保存文件
	if err := f.SaveAs(fileName); err != nil {
		log.Fatalf("保存文件失败: %v", err)
	}

	fmt.Println("Excel 文件已生成: output_with_chart.xlsx")
}

// Function to adjust the column width based on the longest cell content in the column
func autoAdjustColumnWidth(f *excelize.File, sheet, col string) {
	// Get the maximum length of the content in the column
	maxLength := 0
	rows, err := f.GetRows(sheet)
	if err != nil {
		log.Fatal(err)
	}

	for _, row := range rows {
		// Only check the cell in the given column
		if len(row) > 0 && len(row[0]) > maxLength {
			maxLength = len(row[0])
		}
	}

	// Set the column width (here, you can adjust the factor to make it more or less wide)
	// `float64(maxLength)` sets the column width based on the longest string's length
	if maxLength > 0 {
		f.SetColWidth(sheet, col, col, float64(maxLength)+2) // "+2" adds a bit of padding
	}
}

// 创建折线图
func createLineChart(f *excelize.File, dataLength int) error {
	// 定义图表数据范围
	categoriesRange := fmt.Sprintf("Sheet1!$A$2:$A$%d", dataLength+1) // Address 数据范围
	valuesRange := fmt.Sprintf("Sheet1!$B$2:$B$%d", dataLength+1)     // Value 数据范围

	// 定义图表系列
	series := []excelize.ChartSeries{
		{
			Name:       "Sheet1!$B$1",   // 系列名称
			Categories: categoriesRange, // X 轴数据范围
			Values:     valuesRange,     // Y 轴数据范围
		},
	}

	// 定义图表
	chart := excelize.Chart{
		Type:   excelize.Line, // 折线图
		Series: series,        // 系列数据
		Title: []excelize.RichTextRun{
			{
				Text: "统计今日冻结金额",
			},
		}, // 图表标题
		Legend: excelize.ChartLegend{Position: "bottom"}, // 图例位置
		XAxis: excelize.ChartAxis{Title: []excelize.RichTextRun{
			{
				Text: "地址",
			},
		}}, // X 轴标题
		YAxis: excelize.ChartAxis{Title: []excelize.RichTextRun{
			{
				Text: "冻结金额",
			},
		}}, // Y 轴标题
	}

	// 添加图表到工作表
	if err := f.AddChart("Sheet1", "D1", &chart); err != nil {
		return err
	}

	return nil
}

// 等待关闭信号并关闭
func (a *App) waitForShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// 阻塞等待关闭信号
	sig := <-quit
	a.logger.Printf("接收到信号: %v, 开始关闭...", sig)

	// 停止定时任务
	a.ticker.Stop()
	a.done <- true

	// 等待所有任务完成
	time.Sleep(1 * time.Second) // 可根据需要调整
}

func getTransactionData(_txid string) TransactionInfo {
	//txid := ""
	url := "https://api.trongrid.io/walletsolidity/gettransactionbyid"

	payload := strings.NewReader("{\"value\":\"" + _txid + "\"}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var result TransactionInfo
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	return result
}

type TronEvents struct {
	Data []struct {
		BlockNumber           int    `json:"block_number"`
		BlockTimestamp        int64  `json:"block_timestamp"`
		CallerContractAddress string `json:"caller_contract_address"`
		ContractAddress       string `json:"contract_address"`
		EventIndex            int    `json:"event_index"`
		EventName             string `json:"event_name"`
		Result                struct {
			Num0          string `json:"0"`
			TransactionID string `json:"transactionId"`
		} `json:"result"`
		ResultType struct {
			TransactionID string `json:"transactionId"`
		} `json:"result_type"`
		Event         string `json:"event"`
		TransactionID string `json:"transaction_id"`
	} `json:"data"`
	Success bool `json:"success"`
	Meta    struct {
		At          int64  `json:"at"`
		Fingerprint string `json:"fingerprint"`
		Links       struct {
			Next string `json:"next"`
		} `json:"links"`
		PageSize int `json:"page_size"`
	} `json:"meta"`
}

func getCommitAddressBalance(txid string) (string, int64) {
	url := "https://api.trongrid.io/v1/transactions/" + txid + "/events"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("accept", "application/json")
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	//fmt.Println(string(body))
	var result TronTxEvent
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	//log.Println(result)

	for index, datum := range result.Data {

		if index == 1 {
			//log.Println(datum.Result.Num0)

			address41 := strings.ReplaceAll(datum.Result.Num0, "0x", "41")
			target, _ := Convert41ToTAddress(address41)
			//log.Println(target)
			_, amount := getBalance(target)

			return target, amount
		}
	}
	return "", 0
}
func getBalance(tAddress string) (error, int64) {
	trc20Contract := "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t" // USDT
	address := tAddress

	conn := client.NewGrpcClient("grpc.trongrid.io:50051")
	err := conn.Start(grpc.WithInsecure())

	balance, err := conn.TRC20ContractBalance(address, trc20Contract)

	if err != nil {
		return err, 0
	}
	//log.Println(err)
	//log.Println("余额：", balance)
	return err, balance.Int64()
}

func getTronAddress(result TransactionInfo) string {
	address41 := result.RawDataHex[558 : 558+42]
	tAddress, _ := Convert41ToTAddress(address41)
	//log.Println("41address ", address41)
	//log.Println("taddress ", tAddress)
	return tAddress
}

// Convert a 41-prefixed address to a T-prefixed address
func Convert41ToTAddress(address41 string) (string, error) {
	// Step 1: Validate the input address
	if len(address41) < 2 || address41[:2] != tronAddressPrefix {
		return "", fmt.Errorf("invalid 41-prefixed address")
	}

	// Step 2: Decode the hex address into bytes
	addrBytes, err := hex.DecodeString(address41)
	if err != nil {
		return "", fmt.Errorf("failed to decode hex address: %v", err)
	}

	// Step 3: Compute checksum (double SHA256 of the address bytes)
	sha256Hash := sha256.Sum256(addrBytes)
	sha256Hash2 := sha256.Sum256(sha256Hash[:])
	checksum := sha256Hash2[:4]

	// Step 4: Combine address bytes and checksum
	finalAddress := append(addrBytes, checksum...)

	// Step 5: Encode the result in Base58
	tAddress := base58Encode(finalAddress)
	return tAddress, nil
}

// Base58 encoding function
func base58Encode(input []byte) string {
	// Convert the byte array to a big integer
	x := new(big.Int).SetBytes(input)
	base := big.NewInt(58)
	zero := big.NewInt(0)
	mod := new(big.Int)

	var result string
	for x.Cmp(zero) > 0 {
		x.DivMod(x, base, mod)
		result = string(base58Alphabet[mod.Int64()]) + result
	}

	// Add leading '1's for leading zeros in the input
	for _, b := range input {
		if b != 0 {
			break
		}
		result = "1" + result
	}

	return result
}

const (
	// Tron address prefix (hex)
	tronAddressPrefix = "41"
	// Base58 alphabet
	base58Alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
)

type TronTxEvent struct {
	Data []struct {
		BlockNumber           int    `json:"block_number"`
		BlockTimestamp        int64  `json:"block_timestamp"`
		CallerContractAddress string `json:"caller_contract_address"`
		ContractAddress       string `json:"contract_address"`
		EventIndex            int    `json:"event_index"`
		EventName             string `json:"event_name"`
		Result                struct {
			Num0          string `json:"0"`
			TransactionID string `json:"transactionId"`
		} `json:"result"`
		ResultType struct {
			TransactionID string `json:"transactionId"`
		} `json:"result_type"`
		Event         string `json:"event"`
		TransactionID string `json:"transaction_id"`
	} `json:"data"`
	Success bool `json:"success"`
	Meta    struct {
		At       int64 `json:"at"`
		PageSize int   `json:"page_size"`
	} `json:"meta"`
}

type TransactionInfo struct {
	Ret []struct {
		ContractRet string `json:"contractRet"`
	} `json:"ret"`
	Signature []string `json:"signature"`
	TxID      string   `json:"txID"`
	RawData   struct {
		Contract []struct {
			Parameter struct {
				Value struct {
					Data            string `json:"data"`
					OwnerAddress    string `json:"owner_address"`
					ContractAddress string `json:"contract_address"`
				} `json:"value"`
				TypeURL string `json:"type_url"`
			} `json:"parameter"`
			Type string `json:"type"`
		} `json:"contract"`
		RefBlockBytes string `json:"ref_block_bytes"`
		RefBlockHash  string `json:"ref_block_hash"`
		Expiration    int64  `json:"expiration"`
		FeeLimit      int    `json:"fee_limit"`
		Timestamp     int64  `json:"timestamp"`
	} `json:"raw_data"`
	RawDataHex string `json:"raw_data_hex"`
}
