package slowmist

import (
	"strings"
)

func ExtractSlowMistRiskQuery(_message string, _cookie string) string {
	_text := ""
	if strings.HasPrefix(_message, "0x") && len(_message) == 42 {
		_symbol := "USDT-ERC20"
		_addressInfo := GetAddressInfo(_symbol, _message, _cookie)
		_text = GetText(_addressInfo)

		addressProfile := GetAddressProfile(_symbol, _message, _cookie)
		_text7 := "余额：" + addressProfile.BalanceUsd + "\n"
		_text8 := "累计收入：" + addressProfile.TotalReceivedUsd + "\n"
		_text9 := "累计支出：" + addressProfile.TotalSpentUsd + "\n"
		_text10 := "首次活跃时间：" + addressProfile.FirstTxTime + "\n"
		_text11 := "最后活跃时间：" + addressProfile.LastTxTime + "\n"
		_text12 := "交易次数：" + addressProfile.TxCount + "笔" + "\n"
		_text99 := "主要交易对手分析：" + "\n"
		_text100 := ""
		lableAddresList := GetNotSafeAddress("ETH", _message, _cookie)
		if len(lableAddresList.GraphDic.NodeList) > 0 {
			for _, data := range lableAddresList.GraphDic.NodeList {
				if strings.Contains(data.Label, "huione") {
					_text100 = _text100 + data.Title[0:5] + "..." + data.Title[29:34] + " 汇旺" + "\n"
				}
				if strings.Contains(data.Label, "Theft") {
					_text100 = _text100 + data.Title[0:5] + "..." + data.Title[29:34] + " 盗窃" + "\n"
				}
				if strings.Contains(data.Label, "Drainer") {
					_text100 = _text100 + data.Title[0:5] + "..." + data.Title[29:34] + " 诈骗" + "\n"
				}
				if strings.Contains(data.Label, "Banned") {
					_text100 = _text100 + data.Title[0:5] + "..." + data.Title[29:34] + " 制裁" + "\n"
				}
			}
		}

		//_text = _text + _text7 + _text8 + _text9 + _text10 + _text11 + _text12 + _text99 + _text100 + _text16
		_text = _text + _text7 + _text8 + _text9 + _text10 + _text11 + _text12 + _text99 + _text100

		return _text

	}
	if strings.HasPrefix(_message, "T") && len(_message) == 34 {
		_symbol := "USDT-TRC20"
		_addressInfo := GetAddressInfo(_symbol, _message, _cookie)
		_text = GetText(_addressInfo)

		addressProfile := GetAddressProfile(_symbol, _message, _cookie)
		_text7 := "余额：" + addressProfile.BalanceUsd + "\n"
		_text8 := "累计收入：" + addressProfile.TotalReceivedUsd + "\n"
		_text9 := "累计支出：" + addressProfile.TotalSpentUsd + "\n"
		_text10 := "首次活跃时间：" + addressProfile.FirstTxTime + "\n"
		_text11 := "最后活跃时间：" + addressProfile.LastTxTime + "\n"
		_text12 := "交易次数：" + addressProfile.TxCount + "笔" + "\n"
		_text99 := "危险交易对手分析：" + "\n"
		lableAddresList := GetNotSafeAddress(_symbol, _message, _cookie)

		_text100 := ""
		if len(lableAddresList.GraphDic.NodeList) > 0 {
			for _, data := range lableAddresList.GraphDic.NodeList {
				if strings.Contains(data.Label, "huione") {
					_text100 = _text100 + data.Title[0:5] + "..." + data.Title[29:34] + " 汇旺" + "\n"
				}
				if strings.Contains(data.Label, "Theft") {
					_text100 = _text100 + data.Title[0:5] + "..." + data.Title[29:34] + " 盗窃" + "\n"
				}
				if strings.Contains(data.Label, "Drainer") {
					_text100 = _text100 + data.Title[0:5] + "..." + data.Title[29:34] + " 诈骗" + "\n"
				}
				if strings.Contains(data.Label, "Banned") {
					_text100 = _text100 + data.Title[0:5] + "..." + data.Title[29:34] + " 制裁" + "\n"
				}
			}
		}
		_text = _text + _text7 + _text8 + _text9 + _text10 + _text11 + _text12 + _text99 + _text100
		return _text
	}

	return _text
}
