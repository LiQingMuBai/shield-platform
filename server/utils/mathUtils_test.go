package utils

import (
	"fmt"
	"testing"
)

func TestMultiplyBy100ToInt64(t *testing.T) {
	// 示例用法
	floatStr := "16.88"
	result, err := MultiplyBy100ToInt64Decimal(floatStr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Result:", result) // 输出: 12345
}

func TestAddMultipleStringNumbers(t *testing.T) {
	result, err := AddMultipleStringNumbers("10", "0.003")
	if err != nil {
		fmt.Println("错误:", err)
		return
	}
	fmt.Println("总和:", result) // 输出: 总和: 6.6
}
