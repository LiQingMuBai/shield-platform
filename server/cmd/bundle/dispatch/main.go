package main

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/btcsuite/btcd/btcutil/base58"
	"github.com/ushield/aurora-admin/server/core"
	"github.com/ushield/aurora-admin/server/global"
	"github.com/ushield/aurora-admin/server/infrastructure/blockchain/tron"
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
	"sync/atomic"
	"syscall"
	"time"
)

var (
	currentKeyIndex uint32
	//userService                              = service.ServiceGroupApp.SystemServiceGroup.UserService
	//dictDetailService                        = service.ServiceGroupApp.SystemServiceGroup.DictionaryDetailService
	sysOrderService                          = service.ServiceGroupApp.UshieldServiceGroup.UserEnergyOrdersService
	userOperationPackageSubscriptionsService = service.ServiceGroupApp.UshieldServiceGroup.UserPackageSubscriptionsService
)

type App struct {
	ticker *time.Ticker
	done   chan bool
	logger *log.Logger
}

func main() {
	global.GVA_VP = core.Viper() // 初始化Viper
	global.GVA_LOG = core.Zap()  // 初始化zap日志库
	zap.ReplaceGlobals(global.GVA_LOG)
	buddha := `
                  _ooOoo_
                o8888888o
                  88" . "88
              	    (| -_- |)
                  O\  =  /O
               ____/'---'\____
             .'  \\|     |//  '.
            /  \\|||  :  |||//  \
           /  _||||| -:- |||||_  \
           |   | \\\  -  /'| |   |
           | \_|  '\'---'//  |_/ |
           \  .-\__ '-. -' __/-.  /
         ___'. .'  /--.--\  '. .'___
      ."" '<  '.___\_<|>_/___.' _> \"".
     | | :  '- \'. ;'. _/; .'/ /  .' ; |
     \  \ '-.   \_\_'. _.'_/_/  -' _.' /
   ====='-.____'.___ \_____/___.-'____.-'=====
                   '=---='
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
            佛祖保佑        永无BUG
   `
	global.GVA_LOG.Info(buddha)
	global.GVA_DB = initialize.Gorm() // gorm连接数据库
	initialize.DBList()

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
		logger: log.New(os.Stdout, "ORDER-EXCHANGE-ENERGY-SCHEDULER: ", log.LstdFlags),
	}

	// 每隔1min启动定时任务
	app.startScheduler(15 * time.Second)

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
	a.logger.Println("开始能量兑换系统-执行定时任务...")
	startTime := time.Now()

	tronClient, err := tron.NewTronService(global.GVA_CONFIG.System.TRON_FULL_NODE)

	if err != nil {
		global.GVA_LOG.Error(fmt.Sprintf("获取波场节点失败%v\n", err))
		return
	}

	var info ushieldReq.UserPackageSubscriptionsSearch
	info.Page = 1
	info.PageSize = 10_000_000
	subscribeItems, _, err := userOperationPackageSubscriptionsService.GetAllPendingUserPackageSubscriptions(context.Background(), info)

	if err != nil {
		global.GVA_LOG.Error(fmt.Sprintf("获取用户笔数订单失败%v\n", err))
		return
	}

	for _, item := range subscribeItems {
		fmt.Printf("item %v\n", item)

		//0默认初始化状态  1 自动派送 2 手动 3 结束
		if item.Status == 1 {
			energy, bandwidth, err := tronClient.GetEnergyBalance(item.Address)
			if err != nil {
				global.GVA_LOG.Error(fmt.Sprintf("获取用户能量失败%v\n", err))
				return
			}
			fmt.Printf("\n资源:\n")
			fmt.Printf("├─ 能量余额: %d Energy\n", energy)
			fmt.Printf("├─ 带宽余额: %d Bandwidth\n", bandwidth)

			if energy < 65000 {
				global.GVA_LOG.Info(fmt.Sprintf("发送（%d）笔能量给（%s），笔数套餐订单号 %d\n", 1, item.Address, item.Id))

				//调用trxfee接口

				//
				///
				///
				///
				///
				///
				///
				///

				//扣减次数
				item.Times = item.Times - 1

				if item.Times == 0 {
					item.Status = 3
				}
				err := userOperationPackageSubscriptionsService.UpdateUserPackageSubscriptions(context.Background(), item)

				if err != nil {
					return
				}
				//通知用户
				_botToken := global.GVA_CONFIG.System.BotToken
				notifyDispatchEnergy(strconv.FormatInt(item.ChatId, 10), _botToken, item.Address, strconv.FormatInt(item.Times, 10))
			}

		}
	}
	//sendAmount := utils.ConvertFloatToBigInt(global.GVA_CONFIG.System.DEPOSIT_TRX_AMOUNT, 6)

	//apiSecret := global.GVA_CONFIG.System.TRXFEE_APISECRET
	//apiKey := global.GVA_CONFIG.System.TRXFEE_APIKEY
	//baseUrl := global.GVA_CONFIG.System.TRXFEE_BASE_URL
	//trxfeeClient := pkg.NewTrxfeeClient(baseUrl, apiKey, apiSecret)

	//global.GVA_LOG.Info(fmt.Sprintf("发送（%d）笔能量给（%s），订单号 %s\n", count, sysOrder.FromAddress, orderNo))
	//	trxfeeClient.Order(sysOrder.OrderNo, sysOrder.FromAddress, 65_000*count)

	a.logger.Printf("任务完成， 耗时: %v", time.Since(startTime))

}
func notifyDispatchEnergy(_chatID string, _botToken string, _address string, _times string) {
	//var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	//	tgbotapi.NewInlineKeyboardRow(
	//		tgbotapi.NewInlineKeyboardButtonURL("交易详情", "https://tronscan.org/#/address/"+_address),
	//	),
	//)

	message := map[string]interface{}{
		"chat_id": _chatID, // 或直接用 chat_id 如 "123456789"=
		"text": "📢【✅ U盾成功发送一笔能量】\n\n" +
			"接收地址：" + _address + "\n\n" +
			"剩余笔数：" + _times + "\n\n",
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

type TransactionTRXResp struct {
	TxID   string  `json:"txID"`
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float64 `json:"amount"`
}

// 获取指定地址的交易列表
func getTRXTransactionsByAddress(address string, apiURL string, pageSize string) ([]TransactionTRXResp, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	url := fmt.Sprintf("%s/v1/accounts/%s/transactions?only_to=true&limit="+pageSize, apiURL, address)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	//// 简单轮询
	keyIndex := atomic.AddUint32(&currentKeyIndex, 1) % uint32(len(global.TRONGRID_KEYS))
	currentKey := global.TRONGRID_KEYS[keyIndex]

	req.Header.Set("TRON-PRO-API-KEY", currentKey)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch transactions: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("API endpoint not found (404)")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	// 解析JSON响应
	var txResponse TransactionTRXResponse
	err = json.Unmarshal(body, &txResponse)
	if err != nil {
		//fmt.Printf("解析JSON失败: %v\n", err)
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	// 打印TRX交易列表
	fmt.Printf("地址 %s 的TRX交易列表:\n", address)
	trxCount := 0
	var result []TransactionTRXResp // nil slice
	for _, tx := range txResponse.Data {
		// 检查是否为TRX交易
		if len(tx.RawData.Contract) > 0 {
			contract := tx.RawData.Contract[0]

			// 只处理TransferContract类型的交易(TRX转账)
			if contract.Type == "TransferContract" {
				// 确保不是TRC20交易(TRC20交易会有contract_address字段)
				if contract.Parameter.Value.ContractAddress == "" {
					trxCount++
					fmt.Printf("\n交易 #%d\n", trxCount)
					fmt.Printf("交易ID: %s\n", tx.TxID)
					//fmt.Printf("发送方: %s\n", contract.Parameter.Value.OwnerAddress)
					//fmt.Printf("接收方: %s\n", contract.Parameter.Value.ToAddress)
					amount := float64(contract.Parameter.Value.Amount) / 1000000
					fmt.Printf("金额: %.6f TRX\n", amount)

					ownerAddress, _ := TronHexToBase58(contract.Parameter.Value.OwnerAddress)
					toAddress, _ := TronHexToBase58(contract.Parameter.Value.ToAddress)

					fmt.Printf("发送方: %s\n", ownerAddress)
					fmt.Printf("接收方: %s\n", toAddress)
					var resource TransactionTRXResp
					resource.TxID = tx.TxID
					resource.From = ownerAddress
					resource.To = toAddress
					resource.Amount = amount
					result = append(result, resource)

				}
			}
		}
	}

	if trxCount == 0 {
		fmt.Println("未找到TRX交易记录")
	}
	return result, nil
}

type TransactionTRXResponse struct {
	Data []TransactionTRX `json:"data"`
}
type TransactionTRX struct {
	TxID    string `json:"txID"`
	RawData struct {
		Contract []struct {
			Type      string `json:"type"`
			Parameter struct {
				Value struct {
					Amount       int64  `json:"amount"`
					OwnerAddress string `json:"owner_address"`
					ToAddress    string `json:"to_address"`
					// 对于TRC20交易会有不同的字段
					ContractAddress string `json:"contract_address"` // TRC20交易会有此字段
					Data            string `json:"data"`             // TRC20交易会有此字段
				} `json:"value"`
			} `json:"parameter"`
		} `json:"contract"`
	} `json:"raw_data"`
}

func TronHexToBase58(hexAddr string) (string, error) {
	addrBytes, err := hex.DecodeString(hexAddr)
	if err != nil {
		return "", fmt.Errorf("failed to decode hex address: %w", err)
	}
	first := sha256.Sum256(addrBytes)
	second := sha256.Sum256(first[:])
	checksum := second[:4]
	fullBytes := append(addrBytes, checksum...)
	return base58.Encode(fullBytes), nil
}
