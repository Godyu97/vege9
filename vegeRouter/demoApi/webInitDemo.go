package demoApi

import (
	"github.com/gin-gonic/gin"

	"github.com/Godyu97/vege9/midware"
	"github.com/Godyu97/vege9/midware/jwtApi"
	"github.com/Godyu97/vege9/vegeRouter"
	"net/http"
	"time"
)

var ApiObj = Api{
	UriToFnNameM: map[string]string{
		"login": "LOGIN",
	},
}

var JwtDefaultReq midware.JwtReq

func InitHttpDemo(r *gin.Engine) {
	r.GET("/Ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Pong pong pong",
		})
	})
	//group router
	auth := r.Group("/auth")
	//init jwt
	midware.InitJwtObj(
		jwtApi.InitJwt("HQoiqrNg",
			jwtApi.WithExp(time.Hour),
			jwtApi.WithIssuer("hongyu"),
		))
	JwtDefaultReq = midware.JwtReq{
		JwtObj: midware.GetJwtObj(),
		JwtCtxKey: midware.CtxKey{
			JwtCookieKey: midware.JwtCookieKey,
			JwtCtxErrKey: midware.JwtCtxErrKey,
			JwtCtxMcKey:  midware.JwtCtxMcKey,
		},
	}
	auth.Use(midware.JWTAuthMiddleware(JwtDefaultReq))
	//reg Api
	auth.POST(vegeRouter.PathUri, vegeRouter.RegApiHandler(ApiObj))
}
