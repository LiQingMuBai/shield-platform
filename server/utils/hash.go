package utils

import (
	"crypto/md5"
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
)

// BcryptHash 使用 bcrypt 对密码进行加密
func BcryptHash(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes)
}

// BcryptCheck 对比明文密码和数据库的哈希值
func BcryptCheck(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: MD5V
//@description: md5加密
//@param: str []byte
//@return: string

func MD5V(str []byte, b ...byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(b))
}

// GenerateTronOrderID 生成波场订单号（年月日时分 + 波场地址后4位）
func GenerateOrderID(tronAddress string, suffix int) (string, error) {
	// 1. 校验波场地址格式
	tronAddress = strings.TrimSpace(tronAddress)
	//if len(tronAddress) != 34 || !strings.HasPrefix(tronAddress, "T") {
	//	return "", fmt.Errorf("无效的波场地址（必须34位且以T开头）")
	//}

	// 2. 获取当前时间的 "年月日时分"（格式：200601021504）
	timestamp := time.Now().Format("200601021504")

	// 3. 截取波场地址后4位
	addressSuffix := tronAddress[len(tronAddress)-suffix:]

	// 4. 拼接时间 + 地址后4位
	orderID := timestamp + addressSuffix
	return orderID, nil
}
