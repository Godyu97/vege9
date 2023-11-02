package main

import (
	"fmt"
	"strings"
)

func main() {
	//if DEBUG == false {
	//	log.Println("tags release~,return")
	//	return
	//}
	tokenStr := `B eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJUb2tlbiI6eyJpZCI6MTZ9LCJpc3MiOiJhY2Nlc3NfdG9rZW5fand0IiwiZXhwIjoxNzAwMTg0NDIyLCJqdGkiOiI0NWEwZjFhZC1hYzNjLTRkNDQtODQ3Yy03YjUyZmFkMzQzMzUifQ.MTs_ffc5xGC1z5l9KeISzU5kCOvQoCNTYRJKpXZ04yI`
	i := strings.IndexFunc(tokenStr, func(r rune) bool {
		return r == 32
	})
	tokenStr = tokenStr[i+1:]
	fmt.Println(tokenStr)
}
