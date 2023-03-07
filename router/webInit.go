package router

import (
	"github.com/gin-gonic/gin"

	"github.com/Godyu97/vege9/jwtApi"
	middleware "github.com/Godyu97/vege9/middleWare"
	"github.com/Godyu97/vege9/router/demoApi"
	"net/http"
	"time"
)

func InitHttp(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	auth := r.Group("/auth")
	middleware.SetJwtObj(
		jwtApi.InitJwt("godyu",
			jwtApi.WithExp(time.Hour),
			jwtApi.WithIssuer("hongyu"),
		))
	auth.Use(middleware.JWTAuthMiddleware())
	auth.Any("*uri", RegApiHandler(demoApi.ApiObj))
}
