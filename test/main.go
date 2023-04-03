package main

import (
	"github.com/Godyu97/vege9/vegeRouter/demoApi"
	"github.com/gin-gonic/gin"
)

func main() {
	mux := gin.Default()
	demoApi.InitHttpDemo(mux)
	mux.Run(":10928")
}
