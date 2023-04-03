package demoApi

import (
	"github.com/gin-gonic/gin"

	"github.com/Godyu97/vege9/jwtApi"
	"github.com/Godyu97/vege9/midware"
	"github.com/Godyu97/vege9/vegeRouter"
	"net/http"
	"time"
)

func InitHttpDemo(r *gin.Engine) {
	r.GET("/Ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Pong pong pong",
		})
	})
	auth := r.Group("/auth")
	midware.SetJwtObj(
		jwtApi.InitJwt("HQoiqrNg",
			jwtApi.WithExp(time.Hour),
			jwtApi.WithIssuer("hongyu"),
		))
	auth.Use(midware.JWTAuthMiddleware())
	auth.Any("*uri", vegeRouter.RegApiHandler(ApiObj))
}
