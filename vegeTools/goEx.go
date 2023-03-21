package vegeTools

// IfEx
// if 实现的三元表达式
func IfEx[T any](boolExpr bool, trueReturn, falseReturn T) T {
	if boolExpr {
		return trueReturn
	} else {
		return falseReturn
	}
}
