package vegeTools

//s包含数字字符返回true
func HasNum(s string) bool {
	for _, c := range s {
		if '0' <= c && c <= '9' {
			return true
		}
	}
	return false
}
