package main

import (
	"github.com/Godyu97/vege9/vege"
	"log"
	"time"
)

func main() {
	for {
		log.Println(vege.NewGoogleUUID())
		time.Sleep(time.Second)
	}
}
