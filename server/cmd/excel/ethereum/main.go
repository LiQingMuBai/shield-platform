package main

import (
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
	"net/http"
	"os"
	"strconv"
	"strings"
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
                       以太坊黑名单文件通知上线
   `
	fmt.Println(buddha)
	// 初始化应用
	app := &App{
		done:   make(chan bool),
		logger: log.New(os.Stdout, "ETHEREUM-TRACE-ADDRESS-TASK: ", log.LstdFlags),
	}

	c := cron.New()

	// 4. 每天 0 点执行任务
	_, err := c.AddFunc("0 8 * * *", func() {
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
	a.logger.Println("以太坊黑名单文件通知上线-执行定时任务...")
	startTime := time.Now()

	time.Sleep(10 * time.Second)
	parameter := GetTransactionsByAddress_JSONData{
		ID:      67,
		Jsonrpc: "2.0",
		Method:  "qn_getTransactionsByAddress",
		Params: []Params{{
			Address: "0xC6CDE7C39eB2f0F0095F41570af89eFC2C1Ea828",
			Page:    1,
			PerPage: 20}},
	}
	reqParam, err := json.Marshal(parameter)

	if err != nil {
		log.Fatal(err)
		return
	}
	reqBody := strings.NewReader(string(reqParam))
	url := "https://docs-demo.quiknode.pro/"
	req, _ := http.NewRequest("POST", url, reqBody)
	req.Header.Add("accept", "application/json")

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	//fmt.Println(string(body))
	var txs EthereumContractTX
	if err := json.Unmarshal(body, &txs); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

	//log.Println(txs)
	// time.Sleep(1 * time.Second)
	sumbitMap := make(map[string]int64)
	commitMap := make(map[string]int64)
	for _, tx := range txs.Result.PaginatedItems {
		log.Println(tx.TransactionHash)
		//获取交易tx hash
		_txHash := tx.TransactionHash
		time.Sleep(1 * time.Second)
		_address := getPeddingBlackedAddress(_txHash)

		if len(_address) > 0 {
			//说明是pendding的地址，直接获取余额
			log.Println("待定黑名单地址：", _address)
			// time.Sleep(1 * time.Second)
			balance, err := getUSDTBalance(_address)
			if err != nil {
			}
			if balance > 0 {
				sumbitMap[_address] = balance
				log.Println(balance)
			}
		}
		if len(_address) == 0 {
			//说明是已经拉入黑名单
			//fmt.Println(_address)
			_address := getBlackAddress(_txHash)
			time.Sleep(1 * time.Second)
			if len(_address) > 0 {
				log.Println("黑名单地址：", _address)
				balance, err := getUSDTBalance(_address)
				if err != nil {
				}
				if balance > 0 {
					commitMap[_address] = balance
					log.Println(balance)
				}
			}
		}
	}

	//export excel

	log.Println("==========================已提交的====================================")
	for tx, _balance := range sumbitMap {
		log.Println(tx, _balance)

	}
	exportExcel(sumbitMap, "24小时内以太坊预冻结.xlsx")

	// time.Sleep(1 * time.Second)
	//filePath1 := "C:\\Users\\Administrator\\Documents\\shiled-platform\\server\\api\\v1\\system\\24小时内以太坊预冻结.xlsx"
	//filePath1 := "/soft/shiled-platform/server/今日预冻结.xlsx"
	//sendTelegram(filePath1)
	log.Println("==========================已确认的====================================")
	for tx, _balance := range commitMap {
		log.Println(tx, _balance)

	}
	exportExcel(commitMap, "24小时内以太坊已冻结.xlsx")
	// time.Sleep(1 * time.Second)
	//filePath2 := "C:\\Users\\Administrator\\Documents\\shiled-platform\\server\\api\\v1\\system\\24小时内以太坊已冻结.xlsx"
	//filePath2 := "/soft/shiled-platform/server/今日已冻结.xlsx"
	//sendTelegram(filePath2)
	time.Sleep(time.Second * 10)

	//
	a.sendTG("24小时内以太坊预冻结.xlsx")

	a.sendTG("24小时内以太坊已冻结.xlsx")

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

type GetTransactionsByAddress_JSONData struct {
	ID      int      `json:"id"`
	Jsonrpc string   `json:"jsonrpc"`
	Method  string   `json:"method"`
	Params  []Params `json:"params"`
}
type Params struct {
	Address string `json:"address"`
	Page    int    `json:"page"`
	PerPage int    `json:"perPage"`
}

type EthereumContractTX struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  struct {
		Address        string      `json:"address"`
		EnsName        interface{} `json:"ensName"`
		PaginatedItems []struct {
			BlockTimestamp   time.Time   `json:"blockTimestamp"`
			TransactionHash  string      `json:"transactionHash"`
			BlockNumber      string      `json:"blockNumber"`
			TransactionIndex int         `json:"transactionIndex"`
			FromAddress      string      `json:"fromAddress"`
			ToAddress        string      `json:"toAddress"`
			ContractAddress  interface{} `json:"contractAddress"`
			Value            string      `json:"value"`
			Status           string      `json:"status"`
		} `json:"paginatedItems"`
		TotalPages int `json:"totalPages"`
		TotalItems int `json:"totalItems"`
		PageNumber int `json:"pageNumber"`
	} `json:"result"`
}
type Confirm_JSONData struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  struct {
		BlockHash         string      `json:"blockHash"`
		BlockNumber       string      `json:"blockNumber"`
		ContractAddress   interface{} `json:"contractAddress"`
		CumulativeGasUsed string      `json:"cumulativeGasUsed"`
		EffectiveGasPrice string      `json:"effectiveGasPrice"`
		From              string      `json:"from"`
		GasUsed           string      `json:"gasUsed"`
		Logs              []struct {
			Address          string   `json:"address"`
			Topics           []string `json:"topics"`
			Data             string   `json:"data"`
			BlockNumber      string   `json:"blockNumber"`
			TransactionHash  string   `json:"transactionHash"`
			TransactionIndex string   `json:"transactionIndex"`
			BlockHash        string   `json:"blockHash"`
			LogIndex         string   `json:"logIndex"`
			Removed          bool     `json:"removed"`
		} `json:"logs"`
		LogsBloom        string `json:"logsBloom"`
		Status           string `json:"status"`
		To               string `json:"to"`
		TransactionHash  string `json:"transactionHash"`
		TransactionIndex string `json:"transactionIndex"`
		Type             string `json:"type"`
	} `json:"result"`
}
type EthereumERC20 struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  string `json:"result"`
}

type GetTransactionByHash_JSONData struct {
	Method  string   `json:"method"`
	Params  []string `json:"params"`
	ID      int      `json:"id"`
	Jsonrpc string   `json:"jsonrpc"`
}

func getPeddingBlackedAddress(_txHash string) string {
	parameter := GetTransactionByHash_JSONData{
		ID:      1,
		Jsonrpc: "2.0",
		Method:  "eth_getTransactionByHash",
		Params:  []string{_txHash},
	}
	reqParam, err := json.Marshal(parameter)

	if err != nil {

		return ""
	}
	reqBody := strings.NewReader(string(reqParam))
	url := "https://docs-demo.quiknode.pro/"
	req, _ := http.NewRequest("POST", url, reqBody)
	req.Header.Add("accept", "application/json")

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	//fmt.Println(string(body))
	var tx EthereumTX
	if err := json.Unmarshal(body, &tx); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

	if len(tx.Result.Input) > 300 {

		fmt.Println("地址：", "0x"+tx.Result.Input[298:298+40])

		return "0x" + tx.Result.Input[298:298+40]
	} else {
		return ""
	}
}

type EthereumTX struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  struct {
		BlockHash            string        `json:"blockHash"`
		BlockNumber          string        `json:"blockNumber"`
		From                 string        `json:"from"`
		Gas                  string        `json:"gas"`
		GasPrice             string        `json:"gasPrice"`
		MaxFeePerGas         string        `json:"maxFeePerGas"`
		MaxPriorityFeePerGas string        `json:"maxPriorityFeePerGas"`
		Hash                 string        `json:"hash"`
		Input                string        `json:"input"`
		Nonce                string        `json:"nonce"`
		To                   string        `json:"to"`
		TransactionIndex     string        `json:"transactionIndex"`
		Value                string        `json:"value"`
		Type                 string        `json:"type"`
		AccessList           []interface{} `json:"accessList"`
		ChainID              string        `json:"chainId"`
		V                    string        `json:"v"`
		R                    string        `json:"r"`
		S                    string        `json:"s"`
		YParity              string        `json:"yParity"`
	} `json:"result"`
}

func getBlackAddress(_txHash string) string {
	parameter := GetTransactionByHash_JSONData{
		ID:      1,
		Jsonrpc: "2.0",
		Method:  "eth_getTransactionReceipt",
		Params:  []string{_txHash},
	}
	reqParam, err := json.Marshal(parameter)

	if err != nil {

		return ""
	}
	reqBody := strings.NewReader(string(reqParam))
	url := "https://docs-demo.quiknode.pro/"
	req, _ := http.NewRequest("POST", url, reqBody)
	req.Header.Add("accept", "application/json")

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	//fmt.Println(string(body))
	var txlogs Confirm_JSONData
	if err := json.Unmarshal(body, &txlogs); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

	//log.Println("地址: ", "0x"+txlogs.Result.Logs[1].Data[26:26+40])

	if len(txlogs.Result.Logs) > 1 {

		if "0x42e160154868087d6bfdc0ca23d96a1c1cfa32f1b72ba9ba27b69b98a0d819dc" == txlogs.Result.Logs[1].Topics[0] {
			log.Println("Topics : ", txlogs.Result.Logs[1].Topics[0])
			return "0x" + txlogs.Result.Logs[1].Data[26:26+40]
		}
		return ""
	}
	return ""
}

func getUSDTBalance(_address string) (int64, error) {
	url := "https://api.etherscan.io/api?module=account&action=tokenbalance&contractaddress=0xdAC17F958D2ee523a2206206994597C13D831ec7&address=" + _address + "&tag=latest&apikey=X95EDAITM2ASW5QXWDQJMRHP2VDUZ7H85W"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("accept", "application/json")

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	//fmt.Println(string(body))
	var result EthereumERC20
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

	log.Println("余额: ", result.Result)

	i, err := strconv.ParseInt(result.Result, 10, 64)

	return i, err
}

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

func getBalance(tAddress string) (error, int64) {
	trc20Contract := "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t" // USDT
	address := tAddress

	conn := client.NewGrpcClient("grpc.trongrid.io:50051")
	err := conn.Start(grpc.WithInsecure())

	balance, err := conn.TRC20ContractBalance(address, trc20Contract)

	if err != nil {
		return err, 0
	}
	log.Println(err)
	log.Println("余额：", balance)
	return err, balance.Int64()
}

func ExportExcel2(source map[string]int64, fileName string) {
	data := map[string]float64{}
	for k, v := range source {
		data[k] = float64(v) / 1000000
	}

	log.Println(data)

	// 创建一个新的 Excel 文件
	f := excelize.NewFile()

	// 设置表头
	headers := []string{"Address", "Value"}

	//f.SetColWidth("Sheet1", "A1", "B2", 200)
	for col, header := range headers {
		cell, _ := excelize.CoordinatesToCellName(col+1, 1) // 列从1开始，行从1开始
		f.SetCellValue("Sheet1", cell, header)

		//f.SetColWidth("Sheet1", "1", "1", 550)
	}

	// 填充数据
	row := 2 // 从第二行开始填充数据
	for address, value := range data {
		// 写入地址
		cell, _ := excelize.CoordinatesToCellName(1, row)

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
