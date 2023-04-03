package midware

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

// SetJwtObj
// 只允许调用一次
func SetJwtObj(obj *jwtApi.JwtCfg) {
	if jwtObj != nil {
		panic("GWnRffXb1 SetJwtObj has run")
	}
	if obj == nil {
		panic("GWnRffXb2 SetJwtObj obj nil")
	}
	jwtObj = obj
}

// JWTAuthMiddleware
// 基于JWT认证中间件
func JWTAuthMiddleware() gin.HandlerFunc {
	if jwtObj == nil {
		panic("WsHcGnaO please SetJwtObj first")
	}
	return func(c *gin.Context) {
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

// GetMcFromCtx
// 从ctx取得 mc
func GetMcFromCtx(ctx *gin.Context) (mc *jwtApi.MyClaims, err error) {
	t, ok := ctx.Get(JwtCtxMcKey)
	if !ok {
		return nil, errors.New("hLrtXLod JwtCtxMcKey not exists")
	}
	mc = t.(*jwtApi.MyClaims)
	return mc, nil
}

// GetTokenObjFromCtx
// form ctx 取得 mc 并且将token JsonUnmarshal 到obj中，要求obj为对象地址
func GetTokenObjFromCtx(ctx *gin.Context, obj any) error {
	mc, err := GetMcFromCtx(ctx)
	if err != nil {
		return err
	}
	err = mc.TokenObj(obj)
	if err != nil {
		return err
	}
	return nil
}
