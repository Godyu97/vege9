package vege

// RemoveInvalidParentheses 去除s中不成对的括号，pair为成对的括号
func RemoveInvalidParentheses(s string, pair [2]rune) string {
	sr := []rune(s)
	result := make([]rune, 0, len(sr))
	//栈操作
	stack := make([]int, 0)
	removeRight := 0
	for i, ru := range sr {
		if ru == pair[0] {
			stack = append(stack, i-removeRight)
		} else if ru == pair[1] {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			} else {
				removeRight++
				continue
			}
		}
		result = append(result, ru)
	}
	//分割字符串
	str := make([]rune, 0, len(result))
	if len(stack) > 0 {
		str = append(str, result[:stack[0]]...)
		for i := 1; i < len(stack); i++ {
			str = append(str, result[stack[i-1]+1:stack[i]]...)
		}
		//结尾内容
		last := stack[len(stack)-1]
		if last < len(result)-1 {
			str = append(str, result[last+1:]...)
		}
	} else {
		str = result
	}
	return string(str)
}
