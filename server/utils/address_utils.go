package utils

import (
	"encoding/hex"
	"net/url"
	"regexp"
	"strings"
)

// IsValidEthereumAddress 检查是否为有效的以太坊地址
func IsValidEthereumAddress(address string) bool {
	// 移除可能的前后空格
	address = strings.TrimSpace(address)

	// 检查是否以0x开头
	if !strings.HasPrefix(address, "0x") {
		return false
	}

	// 移除0x前缀
	address = address[2:]

	// 检查长度是否为40个字符（20字节）
	if len(address) != 40 {
		return false
	}

	// 检查是否全部是十六进制字符
	_, err := hex.DecodeString(address)
	return err == nil
}

// IsValidTronAddress 检查是否为有效的波场地址
func IsValidTronAddress(address string) bool {
	// 移除可能的前后空格
	address = strings.TrimSpace(address)

	// 波场地址通常以T开头，长度为34个字符
	if len(address) != 34 {
		return false
	}

	// 检查是否以T开头
	if !strings.HasPrefix(address, "T") {
		return false
	}

	// 检查是否只包含Base58字符（1-9, A-Z, a-z，不包括0, O, I, l）
	base58Pattern := "^[1-9A-HJ-NP-Za-km-z]+$"
	matched, err := regexp.MatchString(base58Pattern, address)
	if err != nil || !matched {
		return false
	}

	return true
}

// IsValidCryptoAddress 检查是否为有效的加密货币地址（以太坊或波场）
func IsValidCryptoAddress(address string) (bool, string) {
	address = strings.TrimSpace(address)

	if IsValidEthereumAddress(address) {
		return true, "Ethereum"
	}

	if IsValidTronAddress(address) {
		return true, "Tron"
	}

	return false, "Unknown"
}

// IsValidURL 检查是否为有效的URL
func IsValidURL(str string) bool {
	u, err := url.Parse(str)
	if err != nil {
		return false
	}

	if u.Scheme == "" || u.Host == "" {
		return false
	}

	// 检查常见协议
	validSchemes := []string{"http", "https", "ftp", "ftps", "sftp"}
	valid := false
	for _, scheme := range validSchemes {
		if u.Scheme == scheme {
			valid = true
			break
		}
	}

	if !valid {
		return false
	}

	// 简单验证主机名格式
	hostPattern := `^([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])(\.([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]{0,61}[a-zA-Z0-9]))*$`
	matched, _ := regexp.MatchString(hostPattern, u.Hostname())
	if !matched {
		return false
	}

	return true
}
