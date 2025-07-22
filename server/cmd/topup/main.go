package main

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/btcsuite/btcd/btcutil/base58"
	"github.com/flipped-aurora/gin-vue-admin/server/core"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/initialize"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ushield"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"io"
	"io/ioutil"
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
                       充值上线
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
		logger: log.New(os.Stdout, "ORDER-EXCHANGE-ENERGY-SCHEDULER: ", log.LstdFlags),
	}

	// 每隔1min启动定时任务
	app.startScheduler(1 * time.Minute)

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
	users, total, err := userService.GetUserInfoListAndAddressNotNull()
	if err != nil {
		global.GVA_LOG.Error(fmt.Sprintf("GetUserInfoListAndAddressNotNull失败: %v\n", err))
	}
	if total > 0 {
		for _, user := range users {

			user.DepositAddress = "TP6vxC82dc4YqBzGjEUV7XPfBFJ2m74Yjk"
			log.Println("========================================================")
			log.Println("存款地址 : ", user.DepositAddress)
			log.Println("========================================================")
			//存款地址
			transactions, err := getTRXTransactionsByAddress(user.DepositAddress, global.GVA_CONFIG.System.TRON_FULL_NODE, "50")
			if err != nil {
				global.GVA_LOG.Error(fmt.Sprintf("Error fetching bussiness's transactions: %v\n", err))
				continue
			}

			//先获取trx等待的充值 status=0
			trxlist, err := userTrxDepositsService.GetUserTrxDepositsByStatus(context.Background(), 0)

			if err != nil {
			}
			var trxDeposits []ushield.UserTrxDeposits
			for _, trxModel := range trxlist {
				log.Printf("trxModel:%v\n", trxModel)
				minutes := utils.GetRoundedMinuteDiff(trxModel.CreatedAt, time.Now())
				log.Printf("minutes ： %d \n", minutes)
				if minutes > 15 {
					trxModel.Status = 2
					userTrxDepositsService.UpdateUserTrxDeposits(context.Background(), trxModel)
				} else {
					trxDeposits = append(trxDeposits, trxModel)
				}
			}

			for _, trxModel := range trxDeposits {
				log.Printf("filter trxModel:%v\n", trxModel)
			}

			usdtlist, err := userUsdtDepositsService.GetUserTrxDepositsByStatus(context.Background(), 0)

			if err != nil {
			}
			var usdtDeposits []ushield.UserUsdtDeposits
			for _, usdtModel := range usdtlist {
				log.Printf("usdtModel:%v\n", usdtModel)
				minutes := utils.GetRoundedMinuteDiff(usdtModel.CreatedAt, time.Now())
				log.Printf("minutes ： %d \n", minutes)
				if minutes > 15 {
					usdtModel.Status = 2
					userUsdtDepositsService.UpdateUserUsdtDeposits(context.Background(), usdtModel)
				} else {
					usdtDeposits = append(usdtDeposits, usdtModel)
				}

			}

			//第一步 更新用户的amount
			//第二步 更新用户的tronAmount
			//第三步 根据placehold 去充值placehold表
			log.Println("========================TRX================================")
			for _, trx_transaction := range transactions {
				//fmt.Printf("trx_transaction:%+v\n", trx_transaction)

				//fmt.Println("amount ", trx_transaction.Amount)
				fmt.Println("amountStr ", trx_transaction.AmountStr)

				for _, trxModel := range trxDeposits {

					totalAmount, _ := utils.AddMultipleStringNumbers(trxModel.Amount, trxModel.Placeholder)
					if totalAmount == trx_transaction.AmountStr {
						//命中请给对方添加金额

						//修改状态
						trxModel.Status = 1
						userTrxDepositsService.UpdateUserTrxDeposits(context.Background(), trxModel)

						userTrxPlaceholdersService.UpdateUserTrxPlaceholdersByName(context.Background(), trxModel.Placeholder, 0)

						//修改用户

						tgUser, _ := tgUsersService.GetTgUsersByAssociates(context.Background(), trxModel.UserId)

						tgUser.TronAmount, _ = utils.AddMultipleStringNumbers(tgUser.TronAmount, totalAmount)

						tgUsersService.UpdateTgUsers(context.Background(), tgUser)
					}

				}

			}
			log.Println("=======================USDT=================================")
			_time := utils.GetTimeDaysAgo(1)
			usdt_transactions, err := getIncomingTransactions(user.DepositAddress, global.GVA_CONFIG.System.TRON_FULL_NODE, 50, _time)
			if err != nil {
				global.GVA_LOG.Error(fmt.Sprintf("Error fetching bussiness's transactions: %v\n", err))
				continue
			}
			for _, usdt_transaction := range usdt_transactions {
				//fmt.Printf("usdt_transaction:%+v\n", usdt_transaction)

				//fmt.Println("amount ", usdt_transaction.Amount)
				fmt.Println("amountStr ", usdt_transaction.AmountStr)

				for _, usdtModel := range usdtDeposits {
					totalAmount, _ := utils.AddMultipleStringNumbers(usdtModel.Amount, usdtModel.Placeholder)
					if totalAmount == usdt_transaction.AmountStr {
						//命中请给对方添加金额

						//修改状态
						usdtModel.Status = 1
						userUsdtDepositsService.UpdateUserUsdtDeposits(context.Background(), usdtModel)

						userUsdtPlaceholdersService.UpdateUserUsdtPlaceholdersByName(context.Background(), usdtModel.Placeholder, 0)

						tgUser, _ := tgUsersService.GetTgUsersByAssociates(context.Background(), usdtModel.UserId)

						tgUser.Amount, _ = utils.AddMultipleStringNumbers(tgUser.Amount, totalAmount)

						tgUsersService.UpdateTgUsers(context.Background(), tgUser)

					}

				}

			}

			//for _, transaction := range transactions {
			//	if transaction.Amount <= global.GVA_CONFIG.System.LIMIT_TRANSFER_AMOUNT {
			//		global.GVA_LOG.Info(fmt.Sprintf("订单金额太小，交易: %s，金额: %f\n", transaction.TxID, transaction.Amount))
			//		continue
			//	}
			//
			//	order, err := sysOrderService.GetSysOrderByTxID(transaction.TxID)
			//	if err != nil {
			//		global.GVA_LOG.Error(fmt.Sprintf("获取数据订单失败: %v\n", err))
			//		continue
			//	}
			//
			//	if order.ID > 0 {
			//		global.GVA_LOG.Info(fmt.Sprintf("订单已经发送无需重复: %s\n", order.TxID))
			//		continue
			//	} else {
			//		var sysOrder system.SysOrder
			//		orderNo, _ := pkg.GenerateOrderID(transaction.From, 4)
			//		fmt.Printf("  OrderNo: %s\n", orderNo)
			//		sysOrder.OrderNo = orderNo
			//		sysOrder.TxID = transaction.TxID
			//		sysOrder.FromAddress = transaction.From
			//		sysOrder.ToAddress = transaction.To
			//		sysOrder.Amount = transaction.Amount
			//
			//		//添加一条记录
			//		err := sysOrderService.CreateSysOrder(&sysOrder)
			//
			//		if err != nil {
			//			global.GVA_LOG.Error(fmt.Sprintf("添加一条记录订单失败: %v\n", err))
			//			continue
			//		}
			//
			//		count := int(transaction.Amount / global.GVA_CONFIG.System.LIMIT_TRANSFER_AMOUNT)
			//
			//		if count*int(global.GVA_CONFIG.System.LIMIT_TRANSFER_AMOUNT) > int(trxFeeAccountBalance) {
			//			global.GVA_LOG.Error(fmt.Sprintf("需要(%d)笔数，金额不够需要充值\n", count))
			//			go notifyInsufficientGas(global.GVA_CONFIG.System.ChatID, global.GVA_CONFIG.System.BotToken, accountResp.Data.RechargeAddr, trxFeeAccountBalance)
			//
			//			go func() {
			//				_, err := tronClient.TransferNative(context.Background(), global.GVA_CONFIG.System.MasterPK, trxFeeAccountAddress, sendAmount)
			//				if err != nil {
			//
			//				}
			//			}()
			//
			//			continue
			//		}
			//		global.GVA_LOG.Info(fmt.Sprintf("发送（%d）笔能量给（%s），订单号 %s\n", count, sysOrder.FromAddress, orderNo))
			//		trxfeeClient.Order(sysOrder.OrderNo, sysOrder.FromAddress, 65_000*count)
			//	}
			//}
		}
	}

	a.logger.Printf("任务完成， 耗时: %v", time.Since(startTime))

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
	TxID      string  `json:"txID"`
	From      string  `json:"from"`
	To        string  `json:"to"`
	Amount    float64 `json:"amount"`
	AmountStr string  `json:"amount_str"`
}

func notifyInsufficientGas(_chatID string, _botToken string, _address string, _amount float64) {
	//var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	//	tgbotapi.NewInlineKeyboardRow(
	//		tgbotapi.NewInlineKeyboardButtonURL("交易详情", "https://tronscan.org/#/address/"+_address),
	//	),
	//)

	message := map[string]interface{}{
		"chat_id": _chatID, // 或直接用 chat_id 如 "123456789"=
		"text": "⚠【主地址Trx余额不足警告提醒】\n\n" +
			"📢地址：" + _address + "\n\n" +
			"📢平台余额：      " + fmt.Sprintf("%f", _amount) + "\n\n",
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
					//fmt.Printf("\n交易 #%d\n", trxCount)
					//fmt.Printf("交易ID: %s\n", tx.TxID)
					//fmt.Printf("发送方: %s\n", contract.Parameter.Value.OwnerAddress)
					//fmt.Printf("接收方: %s\n", contract.Parameter.Value.ToAddress)
					amount := float64(contract.Parameter.Value.Amount) / 1000000
					//fmt.Printf("金额: %.6f TRX\n", amount)

					ownerAddress, _ := TronHexToBase58(contract.Parameter.Value.OwnerAddress)
					toAddress, _ := TronHexToBase58(contract.Parameter.Value.ToAddress)

					//fmt.Printf("发送方: %s\n", ownerAddress)
					//fmt.Printf("接收方: %s\n", toAddress)
					var resource TransactionTRXResp
					resource.TxID = tx.TxID
					resource.From = ownerAddress
					resource.To = toAddress
					resource.Amount = amount
					resource.AmountStr = strconv.FormatFloat(amount, 'f', 3, 64)
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

// GetIncomingTransactions 获取地址的转入交易
func getIncomingTransactions(address string, apiURL string, limit int, since time.Time) ([]TxTransaction, error) {
	url := fmt.Sprintf("%s/v1/accounts/%s/transactions/trc20?only_to=true&limit=%d&min_timestamp=%d&contract_address=TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t",
		apiURL, address, limit, since.Unix()*1000)

	//log.Println(url)

	//req, err := http.NewRequest("GET", url, nil)
	//if err != nil {
	//	return nil, fmt.Errorf("创建请求失败: %v", err)
	//}
	//
	//req.Header.Add("Accept", "application/json")
	////req.Header.Add("TRON-PRO-API-KEY", apiKey)
	//
	//client := &http.Client{Timeout: 30 * time.Second}
	//resp, err := client.Do(req)
	//

	keyIndex := atomic.AddUint32(&currentKeyIndex, 1) % uint32(len(global.TRONGRID_KEYS))
	currentKey := global.TRONGRID_KEYS[keyIndex]

	//log.Println(currentKey)

	//time.Sleep(100 * time.Millisecond)

	req, err := http.NewRequest("GET", url, nil)
	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("TRON-PRO-API-KEY", currentKey) // 添加API密钥
	client := &http.Client{Timeout: 30 * time.Second}
	//resp, err := client.Get(url)
	resp, err := client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("API请求失败: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("API返回错误: %s, 响应: %s", resp.Status, string(body))
	}

	var result struct {
		Data []struct {
			TransactionID string `json:"transaction_id"`
			BlockNumber   int64  `json:"block_number"`
			From          string `json:"from"`
			To            string `json:"to"`
			Value         string `json:"value"`
			Timestamp     int64  `json:"block_timestamp"`
			TokenInfo     struct {
				Address string `json:"address"`
			} `json:"token_info"`
		} `json:"data"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("解析响应失败: %v", err)
	}

	var transactions []TxTransaction
	for _, tx := range result.Data {
		// 将字符串金额转换为float64 (USDT通常是6位小数)
		value, err := strconv.ParseFloat(tx.Value, 64)
		if err != nil {
			continue // 跳过金额解析失败的交易
		}
		amount := value / 1e6 // 假设是USDT(6位小数)

		transactions = append(transactions, TxTransaction{
			TxID:        tx.TransactionID,
			BlockNumber: tx.BlockNumber,
			FromAddress: tx.From,
			ToAddress:   tx.To,
			Amount:      decimal.NewFromFloat(amount),
			AmountStr:   strconv.FormatFloat(amount, 'f', 3, 64),
			Timestamp:   time.Unix(tx.Timestamp/1000, 0),
			Contract:    tx.TokenInfo.Address,
		})
	}

	return transactions, nil
}

// Transaction 表示Tron交易数据结构
type TxTransaction struct {
	TxID        string          `json:"txID"`
	BlockNumber int64           `json:"blockNumber"`
	FromAddress string          `json:"from_address"`
	ToAddress   string          `json:"to_address"`
	Amount      decimal.Decimal `json:"amount,omitempty"` // USDT金额
	AmountStr   string          `json:"amount_str"`       // USDT金额
	Timestamp   time.Time       `json:"timestamp"`
	Contract    string          `json:"contract_address,omitempty"` // 合约地址(对于TRC20交易)
}
type Trc20Response struct {
	Data []struct {
		TokenInfo struct {
			Address string `json:"address"`
			Name    string `json:"name"`
			Symbol  string `json:"symbol"`
		} `json:"token_info"`
	} `json:"data"`
	Meta struct {
		Total int `json:"total"`
	} `json:"meta"`
}
