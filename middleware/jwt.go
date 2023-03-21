package middleware

import (
	"errors"
	"github.com/Godyu97/vege9/jwtApi"
	"github.com/gin-gonic/gin"
)

const (
	JwtCookieKey = "VegeJwtToken"
	JwtCtxErrKey = "VegeJwtCtxErr"
	JwtCtxMcKey  = "VegeJwtCtxMc"
)

var jwtObj *jwtApi.JwtCfg

// 只允许调用一次
func SetJwtObj(obj *jwtApi.JwtCfg) {
	if obj != nil {
		jwtObj = obj
	}
}

// 基于JWT认证中间件
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if jwtObj == nil {
			panic("WsHcGnaO please init jwtObj")
		}
		cookie, err := c.Request.Cookie(JwtCookieKey)
		if err != nil {
			// Set jwt Err
			c.Set(JwtCtxErrKey, err)
			return
		}
		mc, err := jwtObj.ParseToken(cookie.Value)
		if err != nil {
			c.Set(JwtCtxErrKey, err)
			return
		}
		//后续的处理函数可以通过c.Get("JwtCtxMcKey")来获取请求的用户信息
		c.Set(JwtCtxMcKey, mc)
	}
}

func GetMcFromCtx(ctx *gin.Context) (mc *jwtApi.MyClaims, err error) {
	t, ok := ctx.Get(JwtCtxMcKey)
	if !ok {
		return nil, errors.New("hLrtXLod JwtCtxMcKey not exists")
	}
	mc = t.(*jwtApi.MyClaims)
	return mc, nil
}
