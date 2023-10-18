package main

import (
	"github.com/Godyu97/vege9/vege"
	"log"
	"time"
)

func main() {
	for {
		log.Println(vege.RandStringMask(8))
		time.Sleep(time.Second)
	}
}
