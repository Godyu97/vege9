package main

import (
	"github.com/Godyu97/vege9/vegeRouter/demoApi"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	demoApi.InitHttpDemo(engine)
	engine.Run(":9010")
}
