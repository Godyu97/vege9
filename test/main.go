package main

import (
	"github.com/Godyu97/vege9/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.InitHttp(r)
	r.Run(":8080")
}
