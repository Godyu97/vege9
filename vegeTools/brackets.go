package vegeTools

import (
	"strings"
	"unicode/utf8"
)

// RemoveInvalidParentheses 去除s中不成对的括号，pair为成对的括号
func RemoveInvalidParentheses(s string, pair [2]rune) string {
	result := new(strings.Builder)
	//栈操作
	stack := make([]int, 0)
	for i, ch := range s {
		if ch == pair[0] {
			stack = append(stack, i)
		} else if ch == pair[1] {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			} else {
				continue
			}
		}
		result.WriteRune(ch)
	}
	//分割字符串
	str := result.String()
	for _, idx := range stack {
		if idx < len(str) {
			//utf8.RuneLen 1
			str = str[:idx] + str[idx+utf8.RuneLen(pair[0]):]
		}
	}
	return str
}
