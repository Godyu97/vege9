package vege

import (
	"strings"
)

//RuneIndex 查找子切片在源切片中的 unicode索引位置
func RuneIndex(src, sub []rune) int {
	result := strings.Index(string(src), string(sub))
	if result > 0 {
		prefix := []byte(string(src))[0:result]
		rs := []rune(string(prefix))
		result = len(rs)
	}
	return result
}
