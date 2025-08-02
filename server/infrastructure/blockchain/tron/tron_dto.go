package tron

import "encoding/json"

type AccountResources struct {
	Energy         int64   `json:"energy"`
	EnergyLimit    int64   `json:"energy_limit"`
	Bandwidth      int64   `json:"bandwidth"`
	BandwidthLimit int64   `json:"bandwidth_limit"`
	FrozenBalance  int64   `json:"frozen_balance"` // 冻结的TRX数量
	EnergyCost     float64 `json:"energy_cost"`    // 1 Energy = X TRX
	BandwidthCost  float64 `json:"bandwidth_cost"` // 1 Bandwidth = X TRX
}

type AccountResponse struct {
	Data    []AccountData   `json:"data"`
	Success bool            `json:"success"`
	Meta    json.RawMessage `json:"meta"`
}

type AccountData struct {
	Address string              `json:"address"`
	Balance int64               `json:"balance"`
	TRC20   []map[string]string `json:"trc20"`
}

type TransactionCreateRequest struct {
	OwnerAddress string `json:"owner_address"`
	ToAddress    string `json:"to_address"`
	Amount       int    `json:"amount"`
	Visible      bool   `json:"visible"`
}

type Transaction struct {
	Visible    bool            `json:"visible"`
	TxID       string          `json:"txID"`
	RawData    json.RawMessage `json:"raw_data"`
	RawDataHex string          `json:"raw_data_hex"`
	Signature  []string        `json:"signature"`
}

type ErrorResponse struct {
	Error string `json:"Error"`
}

type BroadcastTransactionResponse struct {
	Code    string `json:"code"`
	TxID    string `json:"txid"`
	Message string `json:"message"`
}

type TriggerSmartContractRequest struct {
	OwnerAddress     string `json:"owner_address"`
	ContractAddress  string `json:"contract_address"`
	FunctionSelector string `json:"function_selector"`
	Parameter        string `json:"parameter"`
	CallValue        int    `json:"call_value"`
	FeeLimit         int    `json:"fee_limit"`
	Visible          bool   `json:"visible"`
}

type ResponseResult struct {
	Result bool `json:"result"`
}

type TriggerSmartContractResponse struct {
	Result      ResponseResult `json:"result"`
	Transaction Transaction    `json:"transaction"`
}
