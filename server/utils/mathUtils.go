package utils

import (
	"fmt"
	"github.com/shopspring/decimal"
	"strconv"
	"strings"
	"time"
)

func AddDecimalStringsWithPrecision(a, b string, precision int32) (string, error) {
	num1, err := decimal.NewFromString(a)
	if err != nil {
		return "", err
	}

	num2, err := decimal.NewFromString(b)
	if err != nil {
		return "", err
	}

	result := num1.Add(num2)
	return result.StringFixed(precision), nil
}

// GenerateTronOrderID 生成波场订单号（年月日时分 + 波场地址后4位）
func GenerateOrderID(userId string) string {
	// 1. 校验波场地址格式
	userId = strings.TrimSpace(userId)
	//if len(tronAddress) != 34 || !strings.HasPrefix(tronAddress, "T") {
	//	return "", fmt.Errorf("无效的波场地址（必须34位且以T开头）")
	//}

	// 2. 获取当前时间的 "年月日时分"（格式：200601021504）
	timestamp := time.Now().Format("200601021504")

	// 4. 拼接时间 + 地址后4位
	orderID := timestamp + userId
	return orderID
}
func StringToUint32(input string) uint32 {
	// First convert to uint64 to check for overflow
	val, err := strconv.ParseUint(input, 10, 32)
	if err != nil {
		return 0
	}
	return uint32(val)
}

func MultiplyBy100ToInt64Decimal(floatStr string) (int64, error) {
	d, err := decimal.NewFromString(floatStr)
	if err != nil {
		return 0, err
	}

	result := d.Mul(decimal.NewFromInt(100))
	return result.IntPart(), nil
}
func AddMultipleStringNumbers(numbers ...string) (string, error) {
	var total float64

	for _, numStr := range numbers {
		num, err := strconv.ParseFloat(numStr, 64)
		if err != nil {
			return "", fmt.Errorf("'%s' 不是有效的数字: %v", numStr, err)
		}
		total += num
	}

	return strconv.FormatFloat(total, 'f', -1, 64), nil
}

func StringToFloat64(str string) (float64, error) {
	// 处理可能存在的逗号小数点
	str = strings.Replace(str, ",", ".", -1)

	// 去除前后空格
	str = strings.TrimSpace(str)

	// 转换为 float64
	return strconv.ParseFloat(str, 64)
}

func SubtractStringNumbers(a, b string, n float64) (string, error) {
	// 1. 将字符串转为 float64
	numA, err := strconv.ParseFloat(a, 64)
	if err != nil {
		return "", fmt.Errorf("转换 %s 失败: %v", a, err)
	}

	numB, err := strconv.ParseFloat(b, 64)
	if err != nil {
		return "", fmt.Errorf("转换 %s 失败: %v", b, err)
	}

	// 2. 计算减法
	result := numA - numB*n

	// 3. 将结果转为字符串
	return fmt.Sprintf("%v", result), nil
}
func CompareStringsWithFloat(a, b string, n float64) bool {
	// 将字符串转换为 float64
	floatA, errA := strconv.ParseFloat(a, 64)
	floatB, errB := strconv.ParseFloat(b, 64)

	if errA != nil || errB != nil {
		return false
	}

	// 计算 b * 2
	bTimesTwo := floatB * n

	// 比较 a 和 b * 2
	return floatA > bTimesTwo
}
func StringMultiply(s string, n int64) (string, error) {
	// 将字符串转换为 int64
	num, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return "", fmt.Errorf("无法将字符串转换为int64: %v", err)
	}

	// 执行乘法运算
	result := num * n

	// 将结果转换回字符串
	return strconv.FormatInt(result, 10), nil
}

func AddStringsAsFloats(a, b string) string {
	// 1. 将第一个字符串转换成 float64
	num1, err := strconv.ParseFloat(a, 64)
	if err != nil {
		return "0"
	}

	// 2. 将第二个字符串转换成 float64
	num2, err := strconv.ParseFloat(b, 64)
	if err != nil {
		return "0"
	}

	// 3. 相加并返回结果

	sum := num1 + num2
	amount := fmt.Sprintf("%f", sum)

	return amount[0 : len(amount)-3]
}
