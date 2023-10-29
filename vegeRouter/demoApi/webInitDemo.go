package demoApi

import (
	"github.com/gin-gonic/gin"

	"github.com/Godyu97/vege9/midware"
	"github.com/Godyu97/vege9/midware/jwtApi"
	"github.com/Godyu97/vege9/vege"
	"github.com/Godyu97/vege9/vegeRouter"
	"net/http"
	"time"
)

var AuthObj = Auth{
	UriToFnNameM: map[string]string{
		"vege_api": "VegeApi",
	},
}

var CustomerObj = Customer{
	UriToFnNameM: map[string]string{
		"vege_api": "VegeApi",
	},
}

var JwtDefaultReq midware.JwtReq

func InitHttpDemo(r *gin.Engine) {
	r.GET("/Ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Pong pong pong:" + vege.RemoteIp(c.Request),
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
		JwtCtx: midware.CtxKey{
			JwtHeaderKey: "Authorization",
			JwtCtxErrKey: "Vege9AcJwtErr",
			JwtCtxMcKey:  "Vege9AcJwtMc",
		},
	}
	//auth.Use(midware.JWTAuthMiddleware(JwtDefaultReq))
	//reg Auth
	auth.GET(vegeRouter.PathUri, vegeRouter.RegApiHandler(AuthObj))
	auth.POST(vegeRouter.PathUri, vegeRouter.RegApiHandler(AuthObj))

	customer := r.Group("/customer")
	customer.GET(vegeRouter.PathUri, vegeRouter.RegApiHandler(CustomerObj))
	customer.POST(vegeRouter.PathUri, vegeRouter.RegApiHandler(CustomerObj))
}
