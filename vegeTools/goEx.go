package vegeTools

// IfEx
//
//	@Description: if 实现的三元表达式
//	@param boolExpression: 布尔表达式，最终返回一个布尔值
//	@param trueReturnValue: 当 boolExpression 返回值为 true 的时候返回的值
//	@param falseReturnValue: 当 boolExpression 返回值为 false 的时候返回的值
//	@return bool: 三元表达式的结果，为 trueReturnValue 或者 falseReturnValue 中的一个
func IfEx[T any](boolExpression bool, trueReturnValue, falseReturnValue T) T {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}
