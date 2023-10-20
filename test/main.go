package main

import (
	"github.com/Godyu97/vege9/vegelog"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"time"
)

func main() {
	if DEBUG == false {
		log.Println("tags release~,return")
		return
	}
	vegelog.InitLogger("./log.log", zapcore.DebugLevel)
	logger := vegelog.GetZapLogger()
	for {
		time.Sleep(time.Second)
		logger.Debug("hh", zap.Int("a", 125))
	}
}
