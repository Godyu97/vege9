package main

import (
	"github.com/Godyu97/vege9/vege"
	"log"
)

func main() {
	if DEBUG == false {
		log.Println("tags release~,return")
		return
	}
	//vegelog.InitLogger("./log.log", zapcore.DebugLevel)
	//logger := vegelog.GetZapLogger()
	//for {
	//	time.Sleep(time.Second)
	//	logger.Fatal("hh", zap.Int("a", 125))
	//}
	macMap, _ := vege.GetMacAddr()
	for k, v := range macMap {
		log.Println(k, v)
	}

}
