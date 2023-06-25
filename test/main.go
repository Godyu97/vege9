package main

import (
	"github.com/Godyu97/vege9/vegeTools"
	"log"
	"time"
)

func main() {
	for {
		log.Println(vegeTools.RandSelfDefMask(8, "0123456789abcdefghijklmnopqrstuvwxyz"))
		time.Sleep(time.Second)
	}
}
