package utils

import (
	"math"
	"time"
)

func GetTimeDaysAgo(days int) time.Time {
	return time.Now().AddDate(0, 0, -days)
}
func MinutesUntil(target time.Time) float64 {
	now := time.Now()
	if target.Before(now) {
		return -1 // 表示目标时间已过去
	}
	return target.Sub(now).Minutes()
}
func GetRoundedMinuteDiff(start, end time.Time) int {
	diff := end.Sub(start)
	return int(math.Round(diff.Minutes()))
}
