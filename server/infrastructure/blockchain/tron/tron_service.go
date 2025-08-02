package tron

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strings"
)

type TronService struct {
	client *TronClient
	url    string
}

func NewTronService(rpcURL string) (*TronService, error) {
	client := NewTronClient(rpcURL)
	return &TronService{client: client, url: rpcURL}, nil
}

// CheckEnergyBalance 检查账户能量是否足够执行操作
func (s TronService) GetEnergyBalance(fromAddr string) (int64, int64, error) {
	//fromHex, err := Base58ToHex(fromAddr)

	//fmt.Println(fromHex)
	account, err := s.client.GetAccountResources(fromAddr)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to get account: %w", err)
	}

	return account.Energy, account.Bandwidth, nil
}

func (s TronService) GetNativeBalance(ctx context.Context, address string) (*big.Int, error) {
	return s.client.GetNativeBalance(ctx, address)
}

func (s TronService) GetTokenBalance(ctx context.Context, address, tokenAddress string) (*big.Int, error) {
	return s.client.GetTokenBalance(ctx, address, tokenAddress)
}

func (s TronService) IsAcccountActive(ctx context.Context, fromAddress string) (bool, error) {
	_, err := s.client.FetchAccountData(ctx, fromAddress)
	if err != nil {
		return false, err
	}
	return true, nil
}

// 估算USDT转账资源消耗
func (s TronService) EstimateUSDTTransfer(fromAddr, toAddr string, amount *big.Int) (int64, int64, float64, error) {

	//tronGridAPI := "https://young-clean-orb.tron-mainnet.quiknode.pro/9283c9ddb51102d9d22cf1ac5a6e6fc898eeaf77"

	// 1. 构建交易原始数据
	rawData, err := buildUSDTTransferData(fromAddr, toAddr, amount)
	if err != nil {
		return 0, 0, 0, err
	}

	// 2. 估算能量消耗
	energy, err := estimateEnergy(s.url, rawData)
	if err != nil {
		return 0, 0, 0, err
	}

	// 3. 固定带宽消耗 (TRC20标准转账)
	bandwidth := int64(350)

	// 4. 计算手续费 (1 Energy = 1 SUN, 1 Bandwidth = 1 SUN)
	fee := float64(energy+bandwidth) / 1_000_000 // 1 TRX = 1,000,000 SUN

	return energy, bandwidth, fee, nil
}

// 构建USDT转账的原始交易数据
func buildUSDTTransferData(fromAddr, toAddr string, amount *big.Int) (string, error) {
	// 将地址转换为Hex格式 (去掉开头的T)
	fromHex, err := Base58ToHex(fromAddr)
	if err != nil {
		return "", fmt.Errorf("地址转换失败: %w", err)
	}
	toHex, err := Base58ToHex(toAddr)
	if err != nil {
		return "", fmt.Errorf("地址转换失败: %w", err)
	}

	// 构造transfer函数调用数据
	// function: transfer(address,uint256)
	// selector: a9059cbb
	// params: toAddress (32字节) + amount (32字节)
	paramData := fmt.Sprintf("%064s%064x", toHex, amount)

	// 完整的调用数据
	callData := "a9059cbb" + paramData
	usdtContract := "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t"
	// 构造交易请求体
	request := map[string]interface{}{
		"owner_address":     fromHex,
		"contract_address":  HexToBytes(usdtContract),
		"function_selector": "transfer(address,uint256)",
		"parameter":         callData,
		"visible":           true,
	}

	jsonData, err := json.Marshal(request)
	if err != nil {
		return "", fmt.Errorf("JSON编码失败: %w", err)
	}

	return string(jsonData), nil
}

// 估算能量消耗
func estimateEnergy(_tronGridAPI, rawData string) (int64, error) {
	url := _tronGridAPI + "/wallet/estimateenergy"

	resp, err := http.Post(url, "application/json", strings.NewReader(rawData))
	if err != nil {
		return 0, fmt.Errorf("API请求失败: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("读取响应失败: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("API返回错误: %s", string(body))
	}

	var result struct {
		EnergyRequired int64 `json:"energy_required"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return 0, fmt.Errorf("JSON解析失败: %w", err)
	}

	return result.EnergyRequired, nil
}

// 工具函数: Base58地址转Hex (去掉开头的T)
func Base58ToHex(address string) (string, error) {
	if len(address) < 1 {
		return "", fmt.Errorf("无效地址")
	}
	return address[1:], nil // 简单实现，实际需要完整的Base58解码
}

// 工具函数: Hex字符串转字节数组
func HexToBytes(hexStr string) string {
	// 简单实现，实际应该处理各种格式
	return hex.EncodeToString([]byte(hexStr))
}
