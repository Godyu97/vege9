package vegeTools

import (
	"time"
)

// IfEx
// if 实现的三元表达式
func IfEx[T any](boolExpr bool, trueReturn, falseReturn T) T {
	if boolExpr {
		return trueReturn
	} else {
		return falseReturn
	}
}

// 计算fn耗时
func CostTimeDur(fn func()) time.Duration {
	t := time.Now()
	fn()
	return time.Since(t)
}
