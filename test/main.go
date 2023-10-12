package main

import (
	"github.com/Godyu97/vege9/vegeTools"
	"log"
)

func main() {
	result := []any{1, 2, 3}
	log.Println(result[:0]...)
	print(vegeTools.RemoveInvalidParentheses(`（asd（（（啊啊啊`, [2]rune{'（', '）'}))
}
