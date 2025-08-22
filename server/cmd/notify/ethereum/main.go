package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/ushield/aurora-admin/server/core"
	"github.com/ushield/aurora-admin/server/global"
	"github.com/ushield/aurora-admin/server/initialize"
	ushieldReq "github.com/ushield/aurora-admin/server/model/ushield/request"
	"github.com/ushield/aurora-admin/server/service"
	"go.uber.org/zap"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"
)

var (
	currentKeyIndex                uint32
	userService                    = service.ServiceGroupApp.SystemServiceGroup.UserService
	sysOrderService                = service.ServiceGroupApp.SystemServiceGroup.SysOrderService
	userUsdtDepositsService        = service.ServiceGroupApp.UshieldServiceGroup.UserUsdtDepositsService
	userTrxDepositsService         = service.ServiceGroupApp.UshieldServiceGroup.UserTrxDepositsService
	userUsdtPlaceholdersService    = service.ServiceGroupApp.UshieldServiceGroup.UserUsdtPlaceholdersService
	dictDetailService              = service.ServiceGroupApp.SystemServiceGroup.DictionaryDetailService
	tgUsersService                 = service.ServiceGroupApp.UshieldServiceGroup.TgUsersService
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
                       以太坊紧急通知上线
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
	global.TRONGRID_KEYS = strings.Split(global.GVA_CONFIG.System.TRONGRID_KEYS, ",")

	log.Println(global.TRONGRID_KEYS)
	// 初始化应用
	app := &App{
		done:   make(chan bool),
		logger: log.New(os.Stdout, "ETHEREUM-TRACE-ADDRESS-TASK: ", log.LstdFlags),
	}

	// 每隔5min启动定时任务
	app.startScheduler(5 * time.Minute)

	// 等待关闭信号
	app.waitForShutdown()

	app.logger.Println("应用程序已关闭")

}

// 定时任务调度器
func (a *App) startScheduler(interval time.Duration) {
	a.ticker = time.NewTicker(interval)

	// 立即执行第一次任务
	go a.executeTask()

	go func() {
		for {
			select {
			case <-a.ticker.C:
				a.executeTask()
			case <-a.done:
				a.logger.Println("定时任务调度器已停止")
				return
			}
		}
	}()
}
func (a *App) executeTask() {
	a.logger.Println("以太坊紧急通知-执行定时任务...")
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
		// time.Sleep(1 * time.Second)
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

	var info ushieldReq.UserAddressMonitorEventSearch

	info.Page = 1
	info.PageSize = 1_000_000

	//得到正在运行的
	events, _, err := userAddressMonitorEventService.GetUserAddressMonitorEventInfoList(context.Background(), info, 1)
	if err != nil {
		return
	}
	botToken := global.GVA_CONFIG.System.BotToken
	for _, event := range events {

		_, ok := sumbitMap[event.Address]
		if ok {
			event.Times = event.Times + 1
			if event.Times <= 10 {
				err := userAddressMonitorEventService.UpdateUserAddressMonitorEvent(context.Background(), event)
				if err != nil {

					return
				}
				notifyRisk(strconv.FormatInt(event.ChatId, 10), botToken, event.Address, strconv.FormatInt(event.Times, 10))
			}
		}

		//serverTrxPrice, _ := dictDetailService.GetDictionaryInfoByLabel("server_trx_price")
		//serverUSDTPrice, _ := dictDetailService.GetDictionaryInfoByLabel("server_usdt_price")

		tgUser, _ := tgUsersService.GetTgUsersByAssociates(context.Background(), event.ChatId)

		//if utils.CompareStringsWithFloat(serverTrxPrice.Value, tgUser.TronAmount, 1) && utils.CompareStringsWithFloat(serverUSDTPrice.Value, tgUser.Amount, 1) {

		if event.InsufficientTimes == 0 && event.Days >= 29 {
			event.InsufficientTimes = event.InsufficientTimes + 1
			err := userAddressMonitorEventService.UpdateUserAddressMonitorEvent(context.Background(), event)
			if err != nil {
				return
			}
			notifyRiskInsufficientBalance(strconv.FormatInt(event.ChatId, 10), botToken, event.Address, strconv.FormatInt(30-event.Days, 10), tgUser.TronAmount, tgUser.Amount)
			//	}
		}

		//如果到了第30天就需要status=2 结束了

		if event.Days >= 29 {
			event.Status = 2
			err := userAddressMonitorEventService.UpdateUserAddressMonitorEvent(context.Background(), event)
			if err != nil {
				return
			}
		}

	}

	a.logger.Printf("任务完成， 耗时: %v", time.Since(startTime))

}

func notifyRiskInsufficientBalance(_chatID string, _botToken string, _address string, _days string, _tronAmount, _amount string) {
	currentTime := time.Now()

	// 格式化时间字符串，例如：YYYY-MM-DD HH:MM:SS
	formattedTime := currentTime.Format("2006-01-02 15:04:05")
	message := map[string]interface{}{
		"chat_id": _chatID, // 或直接用 chat_id 如 "123456789"=
		"text":
		//	"📢 冻结预警服务即将到期检测余额不足推送\n\n" +
		"📢 冻结预警服务即将到期！\n\n" +
			"地址：" + _address + " \n\n" +
			"剩余天数：" + _days + " 天\n\n" +
			"到期时间：" + formattedTime + "\n\n" +
			"🛑 到期后将自动停止监测，不再推送风险提醒\n\n" +
			"💰 当前余额：\n\n-TRX：" + _tronAmount + "\n-USDT：" + _amount + "\n\n" +
			"请尽快充值以继续保障资产安全",
	}
	// 转换为 JSON
	jsonData, err := json.Marshal(message)
	if err != nil {
		fmt.Println("JSON  parse error...:", err)
		return
	}

	// 发送 POST 请求到 Telegram Bot API
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", _botToken)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("发送消息失败:", err)
		return
	}
	defer resp.Body.Close()

	// 打印响应结果
	//fmt.Println("消息发送状态:", resp.Status)
}

func notifyRisk(_chatID string, _botToken string, _address string, _times string) {
	currentTime := time.Now()

	// 格式化时间字符串，例如：YYYY-MM-DD HH:MM:SS
	formattedTime := currentTime.Format("2006-01-02 15:04:05")
	message := map[string]interface{}{
		"chat_id": _chatID,
		"text": "🚨【USDT冻结预警】第" + _times + "/10次（持续预警中）\n\n" +
			"⚠️ 您的地址已被风控系统标记为即将冻结！\n\n" +
			"地址：" + _address + "\n\n" +
			"风险类型：异常资金流动 + 与受制裁实体交互\n\n" +
			"⚠️ 状态：冻结即将触发\n\n" +
			"⏰ 当前时间：" + formattedTime + "\n\n" +
			"‼️请立即转出资产！避免资产损失！\n\n",
	}
	// 转换为 JSON
	jsonData, err := json.Marshal(message)
	if err != nil {
		fmt.Println("JSON  parse error...:", err)
		return
	}

	// 发送 POST 请求到 Telegram Bot API
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", _botToken)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("发送消息失败:", err)
		return
	}
	defer resp.Body.Close()

	// 打印响应结果
	//fmt.Println("消息发送状态:", resp.Status)
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
	time.Sleep(1 * time.Second)
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

		log.Println("地址： ", "0x"+txlogs.Result.Logs[1].Data[26:26+40])
		if "0x42e160154868087d6bfdc0ca23d96a1c1cfa32f1b72ba9ba27b69b98a0d819dc" == txlogs.Result.Logs[1].Topics[0] {
			log.Println("Topics : ", txlogs.Result.Logs[1].Topics[0])
			return "0x" + txlogs.Result.Logs[1].Data[26:26+40]
		}
		return ""
	}
	return ""
}

func getUSDTBalance(_address string) (int64, error) {
	time.Sleep(1 * time.Second)
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
