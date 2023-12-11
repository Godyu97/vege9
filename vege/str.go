package vege

import (
	"strings"
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
