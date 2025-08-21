package main

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"github.com/ushield/aurora-admin/server/core"
	"github.com/ushield/aurora-admin/server/global"
	"github.com/ushield/aurora-admin/server/initialize"
	ushieldReq "github.com/ushield/aurora-admin/server/model/ushield/request"
	"github.com/ushield/aurora-admin/server/service"
	"go.uber.org/zap"
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

var (
	currentKeyIndex                    uint32
	userService                        = service.ServiceGroupApp.SystemServiceGroup.UserService
	sysOrderService                    = service.ServiceGroupApp.SystemServiceGroup.SysOrderService
	userUsdtDepositsService            = service.ServiceGroupApp.UshieldServiceGroup.UserUsdtDepositsService
	userTrxDepositsService             = service.ServiceGroupApp.UshieldServiceGroup.UserTrxDepositsService
	userUsdtPlaceholdersService        = service.ServiceGroupApp.UshieldServiceGroup.UserUsdtPlaceholdersService
	userTrxPlaceholdersService         = service.ServiceGroupApp.UshieldServiceGroup.UserTrxPlaceholdersService
	tgUsersService                     = service.ServiceGroupApp.UshieldServiceGroup.TgUsersService
	dictDetailService                  = service.ServiceGroupApp.SystemServiceGroup.DictionaryDetailService
	merchantAddressMonitorEventService = service.ServiceGroupApp.UshieldServiceGroup.MerchantAddressMonitorEventService
)

type App struct {
	ticker *time.Ticker
	done   chan bool
	logger *log.Logger
}

func main() {

	global.GVA_VP = core.Viper() // 初始化Viper

	buddha := `============================================
                       波场紧急通知上线
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
		logger: log.New(os.Stdout, "TRON-TRACE-ADDRESS-TASK: ", log.LstdFlags),
	}

	// 每隔5min启动定时任务
	app.startScheduler(12 * time.Minute)

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
	a.logger.Println("波场紧急通知-执行定时任务...")
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
			time.Sleep(1 * time.Second)
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

	var info ushieldReq.MerchantAddressMonitorEventSearch

	info.Page = 1
	info.PageSize = 1_000_000

	//得到正在运行的
	monitorEvents, _, err := merchantAddressMonitorEventService.GetMerchantAddressMonitorEventInfoList(context.Background(), info, 1)
	if err != nil {
		return
	}

	for _, event := range monitorEvents {

		if sumbitMap[event.Address] >= 0 {
			event.Times = event.Times + 1
			if event.Times <= 10 {
				err := merchantAddressMonitorEventService.UpdateMerchantAddressMonitorEvent(context.Background(), event)
				if err != nil {

					return
				}
				log.Println("进行回调：", event.Callback)
				callback(event.Callback, event.Address, 2)
			}
		}

		if event.Days >= 30 {
			event.Status = 2
			err := merchantAddressMonitorEventService.UpdateMerchantAddressMonitorEvent(context.Background(), event)
			if err != nil {
				return
			}
		}
	}

	a.logger.Printf("任务完成， 耗时: %v", time.Since(startTime))

}

func callback(url string, _address string, result int64) {

	code := 20000
	message := "20000"

	if result == 1 {
		//正常
		code = 20000
		message = "正常"
	} else if result == 0 {
		//暂停服务
		code = 20001
		message = "服务已经暂停该预警功能"
	} else if result == 2 {
		//将被拉入黑名单
		code = 20002
		message = "即将被拉入黑名单"
	} else if result == 4 {
		//已经是黑名单
		code = 20004
		message = "已经被拉入黑名单"
	}

	data := map[string]interface{}{
		"code":    code,     // 或直接用 chat_id 如 "123456789"=
		"address": _address, // 或直接用 chat_id 如 "123456789"=
		"message": message,
	}
	// 转换为 JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("JSON  parse error...:", err)
		return
	}

	// 发送 POST 请求

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
