package main

import (
	"github.com/Godyu97/vege9/vegeTools"
	"log"
)

func main() {
	s := vegeTools.RemoveInvalidParentheses("()()()(（丝丝））()", [2]rune{'（', '）'})
	s = vegeTools.RemoveInvalidParentheses(s, [2]rune{'(', ')'})
	log.Println(s)
}
