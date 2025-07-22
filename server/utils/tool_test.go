package utils

import (
	"fmt"
	"testing"
	"time"
)

func TestMinutesUntil(t *testing.T) {
	t1 := time.Now()
	t2 := t1.Add(111*time.Hour + 32*time.Minute + 30*time.Second)

	minutes := GetRoundedMinuteDiff(t1, t2)
	fmt.Printf("四舍五入后的分钟差: %d\n", minutes) // 输出: 93

}
