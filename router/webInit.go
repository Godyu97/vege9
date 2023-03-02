package router

import (
	"github.com/gin-gonic/gin"

	"authApi"
	"net/http"
)

func InitHttp(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	auth := r.Group("/auth")
	auth.Any("*uri", RegApiHandler(authApi.AuthApiObj))
}
