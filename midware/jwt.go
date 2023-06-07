package midware

import (
	"errors"
	"github.com/Godyu97/vege9/midware/jwtApi"
	"github.com/gin-gonic/gin"
	"sync"
)

const (
	JwtCookieKey = "Vege9JwtToken"
	JwtCtxErrKey = "Vege9JwtCtxErr"
	JwtCtxMcKey  = "Vege9JwtCtxMc"
)

type JwtReq struct {
	JwtObj    *jwtApi.JwtCfg
	JwtCtxKey CtxKey
}

type CtxKey struct {
	JwtCookieKey string
	JwtCtxErrKey string
	JwtCtxMcKey  string
}

// JWTAuthMiddleware
// 基于JWT认证中间件
func JWTAuthMiddleware(req JwtReq) gin.HandlerFunc {
	if req.JwtObj == nil {
		panic("WsHcGnaO please req.JwtObj is nil")
	}
	return func(c *gin.Context) {
		cookie, err := c.Request.Cookie(req.JwtCtxKey.JwtCookieKey)
		if err != nil {
			// Set jwt Err
			c.Set(req.JwtCtxKey.JwtCtxErrKey, err)
			return
		}
		mc, err := req.JwtObj.ParseToken(cookie.Value)
		if err != nil {
			c.Set(req.JwtCtxKey.JwtCtxErrKey, err)
			return
		}
		//后续的处理函数可以通过c.Get("JwtCtxMcKey")来获取请求的用户信息
		c.Set(req.JwtCtxKey.JwtCtxMcKey, mc)
	}
}

// GetMcFromCtx
// 从ctx取得 mc
func GetMcFromCtx(ctx *gin.Context, req JwtReq) (mc *jwtApi.MyClaims, err error) {
	t, ok := ctx.Get(req.JwtCtxKey.JwtCtxMcKey)
	if !ok {
		return nil, errors.New("hLrtXLod JwtCtxMcKey not exists")
	}
	mc = t.(*jwtApi.MyClaims)
	return mc, nil
}

// GetTokenObjFromCtx
// form ctx 取得 mc 并且将token JsonUnmarshal 到obj中，要求obj为对象地址
func GetTokenObjFromCtx(ctx *gin.Context, req JwtReq, obj any) error {
	mc, err := GetMcFromCtx(ctx, req)
	if err != nil {
		return err
	}
	err = mc.TokenObj(obj)
	if err != nil {
		return err
	}
	return nil
}

// 单例模式
var jwtObj *jwtApi.JwtCfg
var setOnce sync.Once

func InitJwtObj(obj *jwtApi.JwtCfg) {
	if obj == nil {
		panic("GWnRffXb2 InitJwtObj obj nil")
	}
	setOnce.Do(func() {
		jwtObj = obj
	})
}

func GetJwtObj() *jwtApi.JwtCfg {
	if jwtObj != nil {
		return jwtObj
	}
	panic("IsRzQnTh please InitJwtObj first")
}
