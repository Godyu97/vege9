package vege

import (
	"math"
)

// CountTotalPage 计算总页数
func CountTotalPage(total int64, size int) int {
	return int(math.Ceil(float64(total) / float64(size)))
}
