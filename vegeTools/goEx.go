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

// 比较old new,相等返回类型零值，不相等返回new
func DiffField[T comparable](o, n T) (t T) {
	if o == n {
		return
	} else {
		return n
	}
}
