package main

import (
	"log"
	"fmt"
)

func main() {
	const s = `
	%s
  	昨夜雨疏风骤，
  	浓睡不消残酒。
	试问卷帘人，
	却道海棠依旧。
	知否？知否？
	应是绿肥红瘦。
    %s
`
	log.Println(fmt.Sprintf(s, "```", "```"))
	log.Println(fmt.Sprintf(`
	%s
  	昨夜雨疏风骤，
  	浓睡不消残酒。
	试问卷帘人，
	却道海棠依旧。
	知否？知否？
	应是绿肥红瘦。
    %s
`, "```", "```"))

}
