package vege

import (
	"regexp"
)

func ValidatePhone(phone string) bool {
	// 正则表达式：以 1 开头，后跟 10 位数字
	regex := `^1[3-9]\d{9}$`
	re := regexp.MustCompile(regex)
	return re.MatchString(phone)
}

func ValidateEmail(email string) bool {
	// 正则表达式：简单的邮箱格式验证
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(regex)
	return re.MatchString(email)
}
