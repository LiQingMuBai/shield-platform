package tron

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"testing"
)

func TestEstimateEnrgyandBandwidth(t *testing.T) {

	tronService, _ := NewTronService("https://api.trongrid.io")

	// 示例参数
	fromAddress := "TXyNzUFYLzSqzJCSwkyRRPJHCMXBDMCCDB" // 替换为你的地址
	//toAddress := "TPXH2iHQY6V58uPy8LYAc2gqqyGZtXpKBN"   // 替换为接收地址
	//amount := big.NewInt(1_000_000_000_000)             // 100 USDT (6 decimals)

	// 估算USDT转账费用
	energy, bandwidth, err := tronService.GetEnergyBalance(fromAddress)
	if err != nil {
		log.Fatalf("估算失败: %v", err)
	}
	// 打印结果
	fmt.Printf("\n资源:\n")
	fmt.Printf("├─ 能量消耗: %d Energy\n", energy)
	fmt.Printf("├─ 带宽消耗: %d Bandwidth\n", bandwidth)
	//fmt.Printf("├─ 预估手续费: %.6f TRX\n", fee)
	//fmt.Printf("└─ 备注: 实际费用取决于网络状况\n\n")
}

func TestGetAccountResources(t *testing.T) {

	url := "https://api.trongrid.io/wallet/getaccountresource"

	payload := strings.NewReader("{\"address\":\"TXyNzUFYLzSqzJCSwkyRRPJHCMXBDMCCDB\",\"visible\":true}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	var result struct {
		FreeNetLimit int `json:"freeNetLimit"`
		NetLimit     int `json:"NetLimit"`
		AssetNetUsed []struct {
			Key   string `json:"key"`
			Value int    `json:"value"`
		} `json:"assetNetUsed"`
		AssetNetLimit []struct {
			Key   string `json:"key"`
			Value int    `json:"value"`
		} `json:"assetNetLimit"`
		TotalNetLimit     int64 `json:"TotalNetLimit"`
		TotalNetWeight    int64 `json:"TotalNetWeight"`
		EnergyLimit       int   `json:"EnergyLimit"`
		TotalEnergyLimit  int64 `json:"TotalEnergyLimit"`
		TotalEnergyWeight int64 `json:"TotalEnergyWeight"`
	}

	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
	}

	fmt.Println(result.EnergyLimit)
	fmt.Println(result.FreeNetLimit + result.NetLimit)
}
