package util

import "time"

// 当天开始时间
func GetStart(date time.Time) time.Time {

	return time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.Local)
}

// 当天结束时间 精确到毫秒
func GetEnd(date time.Time) time.Time {

	return time.Date(date.Year(), date.Month(), date.Day(), 23, 59, 59, 999000000, time.Local)
}

// 加减天数
func AddDays(date time.Time, num int) time.Time {
	return date.AddDate(0, 0, num)
}

// 格式化日期 年月日
func FormatDay(date time.Time) string {
	return date.Format("2006-01-02")
}

// 格式化日期 年月日 时分秒
func FormatSec(date time.Time) string {
	return date.Format("2006-01-02 15:04:05")
}

// 格式化日期 年月日 时分秒 毫秒
func FormatMil(date time.Time) string {
	return date.Format("2006-01-02 15:04:05.000")
}

// 字符串转时间
func DateToTime(date string) time.Time {
	times, _ := time.Parse("2006-01-02 15:04:05", date)
	return times
}
