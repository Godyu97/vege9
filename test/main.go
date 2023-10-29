package main

import (
	"github.com/Godyu97/vege9/vegeRouter/demoApi"
	"github.com/Godyu97/vege9/vegelog"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap/zapcore"
)

func main() {
	//if DEBUG == false {
	//	log.Println("tags release~,return")
	//	return
	//}
	vegelog.InitLogger("./t.log", zapcore.DebugLevel)
	w := vegelog.GetLogWriter()
	gin.DefaultErrorWriter = w
	gin.DefaultWriter = w
	engine := gin.Default()
	demoApi.InitHttpDemo(engine)
	engine.Run(":8000")
}
