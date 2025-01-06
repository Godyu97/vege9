package vege

import (
	"strings"
	"time"
	"fmt"
)

// HasNum s包含数字字符返回true
func HasNum(s string) bool {
	for _, c := range s {
		if '0' <= c && c <= '9' {
			return true
		}
	}
	return false
}

// DelRightRN 移除字符串 item 中最后 n 个特定字符 r，并返回修改后的字符串。
// 如果 item 中 r 出现的次数小于 n，则不进行修改。
// 参数:
//
//	item: 要处理的字符串
//	r: 要移除的特定字符
//	n: 要移除的 r 的数量
//
// 返回值:
//
//	修改后的字符串
func DelRightRN(item string, r string, n int) string {
	count := strings.Count(item, r)
	if count >= n {
		for i := 0; i < n; i++ {
			lastIndex := strings.LastIndex(item, r)
			item = item[:lastIndex] + item[lastIndex+1:]
		}
	}
	return item
}

//StringsContainsSlice 展开arr判断s是否Contains
func StringsContainsSlice(s string, arr []string) bool {
	for _, item := range arr {
		if strings.Contains(s, item) == true {
			return true
		}
	}
	return false
}

//ReplaceLast 替换字符串最后一个匹配项
func ReplaceLast(s, old, new string) string {
	lastIndex := strings.LastIndex(s, old)
	if lastIndex == -1 {
		return s
	}
	return s[:lastIndex] + new + s[lastIndex+len(old):]
}

// ParseDateStr 解析常见的日期格式字符串到 time.Time
func ParseDateStr(dateStr string) (time.Time, error) {
	// 定义常见日期格式的切片
	formats := []string{
		"2006-01-02",                // yyyy-MM-dd
		"2006/01/02",                // yyyy/MM/dd
		"02-01-2006",                // dd-MM-yyyy
		"02/01/2006",                // dd/MM/yyyy
		"2006-01-02 15:04:05",       // yyyy-MM-dd HH:mm:ss
		"2006/01/02 15:04:05",       // yyyy/MM/dd HH:mm:ss
		"02-01-2006 15:04:05",       // dd-MM-yyyy HH:mm:ss
		"02/01/2006 15:04:05",       // dd/MM/yyyy HH:mm:ss
		"2006-01-02T15:04:05Z07:00", // ISO8601
		"2006-01-02T15:04:05.000Z",  // ISO8601 with milliseconds

		"2006-01-02 15:04",
		"2006/01/02 15:04",
	}

	var parsedTime time.Time
	var err error

	// 尝试使用每个格式解析日期字符串
	for _, format := range formats {
		parsedTime, err = time.Parse(format, dateStr)
		if err == nil {
			return parsedTime, nil
		}
	}

	return time.Time{}, fmt.Errorf("krEKIhxV 无法解析日期字符串: %s", dateStr)
}
